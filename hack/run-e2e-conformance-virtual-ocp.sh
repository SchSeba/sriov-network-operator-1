#!/usr/bin/env bash
set -xeo pipefail

cluster_name=${CLUSTER_NAME:-ocp-virt}
domain_name=lab

api_ip=${API_IP:-192.168.122.253}
virtual_router_id=${VIRTUAL_ROUTER_ID:-253}
registry="default-route-openshift-image-registry.apps.${cluster_name}.${domain_name}"
HOME="/root"

here="$(dirname "$(readlink --canonicalize "${BASH_SOURCE[0]}")")"
root="$(readlink --canonicalize "$here/..")"

check_requirements() {
  for cmd in kcli virsh podman make go jq base64; do
    if ! command -v "$cmd" &> /dev/null; then
      echo "$cmd is not available"
      exit 1
    fi
  done
  return 0
}

echo "## checking requirements"
check_requirements
echo "## delete existing cluster name $cluster_name"
kcli delete cluster $cluster_name -y
kcli delete network $cluster_name -y

function cleanup {
  kcli delete cluster $cluster_name -y
  kcli delete network $cluster_name -y
  podman logout $registry
}

if [ -z $SKIP_DELETE ]; then
  trap cleanup EXIT
fi

kcli create network -c 192.168.${virtual_router_id}.0/24 --nodhcp -i $cluster_name

cat <<EOF > ./${cluster_name}-plan.yaml
tag: 4.14.0-rc.1
ctlplane_memory: 24576
worker_memory: 8192
pool: default
disk_size: 50
network: default
api_ip: $api_ip
virtual_router_id: $virtual_router_id
domain: $domain_name
ctlplanes: 1
workers: 2
machine: q35
network_type: OVNKubernetes
pull_secret: /root/openshift_pull.json
vmrules:
  - $cluster_name-worker-.*:
      nets:
        - name: default
          numa: 0
        - name: $cluster_name
          type: igb
          vfio: true
          noconf: true
          numa: 0
        - name: $cluster_name
          type: igb
          vfio: true
          noconf: true
          numa: 1
      numcpus: 6
      numa:
        - id: 0
          vcpus: 0,2,4
          memory: 4096
        - id: 1
          vcpus: 1,3,5
          memory: 4096

EOF

kcli create cluster openshift --paramfile ./${cluster_name}-plan.yaml $cluster_name
kubectl create clusterrolebinding authenticated-registry-viewer --clusterrole registry-viewer --group system:unauthenticated

export KUBECONFIG=$HOME/.kcli/clusters/$cluster_name/auth/kubeconfig
export PATH=$PWD:$PATH


ATTEMPTS=0
MAX_ATTEMPTS=72
ready=false
sleep_time=10

until $ready || [ $ATTEMPTS -eq $MAX_ATTEMPTS ]
do
    echo "waiting for cluster to be ready"
    if [ `kubectl get node | grep Ready | wc -l` == 3 ]; then
        echo "cluster is ready"
        ready=true
    else
        echo "cluster is not ready yet"
        sleep $sleep_time
    fi
    ATTEMPTS=$((ATTEMPTS+1))
done

if ! $ready; then
    echo "Timed out waiting for cluster to be ready"
    kubectl get nodes
    exit 1
fi

echo "## label cluster workers as sriov capable"
kubectl label node $cluster_name-worker-0.$domain_name feature.node.kubernetes.io/network-sriov.capable=true --overwrite
kubectl label node $cluster_name-worker-1.$domain_name feature.node.kubernetes.io/network-sriov.capable=true --overwrite

controller_ip=`kubectl get node -o wide | grep ctlp | awk '{print $6}'`

if [ `cat /etc/hosts | grep ${api_ip} | grep "default-route-openshift-image-registry.apps.${cluster_name}.${domain_name}" | wc -l` == 0 ]; then
  echo "adding registry to hosts"
  sed -i "s/${api_ip}/${api_ip} default-route-openshift-image-registry.apps.${cluster_name}.${domain_name}/g" /etc/hosts
fi


cat <<EOF | oc apply -f -
apiVersion: v1
kind: PersistentVolume
metadata:
  name: registry-pv
spec:
  capacity:
    storage: 60Gi
  volumeMode: Filesystem
  accessModes:
  - ReadWriteMany
  persistentVolumeReclaimPolicy: Delete
  storageClassName: registry-local-storage
  local:
    path: /mnt/
  nodeAffinity:
    required:
      nodeSelectorTerms:
      - matchExpressions:
        - key: kubernetes.io/hostname
          operator: In
          values:
          - ${cluster_name}-ctlplane-0.${domain_name}
---
apiVersion: v1
kind: PersistentVolumeClaim
metadata:
  name: registry-pv-claim
  namespace: openshift-image-registry
spec:
  accessModes:
    - ReadWriteMany
  volumeMode: Filesystem
  resources:
    requests:
      storage: 60Gi
  storageClassName: registry-local-storage
EOF

kubectl patch configs.imageregistry.operator.openshift.io/cluster --patch '{"spec":{"defaultRoute":true,"storage":{"emptyDir": null,"pvc":{"claim":"registry-pv-claim"}},"topologySpreadConstraints":[],"rolloutStrategy":"Recreate","tolerations":[{"effect":"NoSchedule","key":"node-role.kubernetes.io/master","operator":"Exists"},{"effect":"NoSchedule","key":"node-role.kubernetes.io/control-plane","operator":"Exists"}]}}' --type=merge

export ENABLE_ADMISSION_CONTROLLER=true
export SKIP_VAR_SET=""
export NAMESPACE="openshift-sriov-network-operator"
export OPERATOR_NAMESPACE="openshift-sriov-network-operator"
export OPERATOR_EXEC=kubectl
export CLUSTER_TYPE=openshift
export DEV_MODE=TRUE
export VIRTUAL_ENV=TRUE

export SRIOV_NETWORK_OPERATOR_IMAGE="$registry/$NAMESPACE/sriov-network-operator:latest"
export SRIOV_NETWORK_CONFIG_DAEMON_IMAGE="$registry/$NAMESPACE/sriov-network-config-daemon:latest"
export SRIOV_NETWORK_WEBHOOK_IMAGE="$registry/$NAMESPACE/sriov-network-operator-webhook:latest"

envsubst< deploy/namespace.yaml | kubectl apply -f -

echo "## build operator image"
podman build -t "${SRIOV_NETWORK_OPERATOR_IMAGE}" -f "${root}/Dockerfile" "${root}"

echo "## build daemon image"
podman build -t "${SRIOV_NETWORK_CONFIG_DAEMON_IMAGE}" -f "${root}/Dockerfile.sriov-network-config-daemon" "${root}"

echo "## build webhook image"
podman build -t "${SRIOV_NETWORK_WEBHOOK_IMAGE}" -f "${root}/Dockerfile.webhook" "${root}"

dockercgf=`kubectl -n ${NAMESPACE} get sa builder -oyaml | grep imagePullSecrets -A 1 | grep -o "builder-.*"`
auth=`kubectl -n ${NAMESPACE} get secret ${dockercgf} -ojson | jq '.data.".dockercfg"'`
auth="${auth:1:-1}"
auth=`echo ${auth} | base64 -d`
echo ${auth} > registry-login.conf

pass=$( jq .\"$registry\".password registry-login.conf )
podman login -u serviceaccount -p ${pass:1:-1} $registry --tls-verify=false

podman push --tls-verify=false "${SRIOV_NETWORK_OPERATOR_IMAGE}"
podman push --tls-verify=false "${SRIOV_NETWORK_CONFIG_DAEMON_IMAGE}"
podman push --tls-verify=false "${SRIOV_NETWORK_WEBHOOK_IMAGE}"

podman logout $registry

echo "## apply CRDs"
kubectl apply -k $root/config/crd

export SRIOV_NETWORK_OPERATOR_IMAGE="image-registry.openshift-image-registry.svc:5000/$NAMESPACE/sriov-network-operator:latest"
export SRIOV_NETWORK_CONFIG_DAEMON_IMAGE="image-registry.openshift-image-registry.svc:5000/$NAMESPACE/sriov-network-config-daemon:latest"
export SRIOV_NETWORK_WEBHOOK_IMAGE="image-registry.openshift-image-registry.svc:5000/$NAMESPACE/sriov-network-operator-webhook:latest"

echo "## deploying SRIOV Network Operator"
hack/deploy-setup.sh $NAMESPACE

echo "## wait for sriov operator to be ready"

echo "## wait for sriov operator to be ready"
hack/deploy-wait.sh

if [ -z $SKIP_TEST ]; then
 echo "## run sriov e2e conformance tests"
 SUITE=./test/conformance hack/run-e2e-conformance.sh
fi
