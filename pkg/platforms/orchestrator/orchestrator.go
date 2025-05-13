package orchestrator

import (
	"context"
	"fmt"

	corev1 "k8s.io/api/core/v1"

	"github.com/k8snetworkplumbingwg/sriov-network-operator/pkg/consts"
	"github.com/k8snetworkplumbingwg/sriov-network-operator/pkg/platforms/orchestrator/kubernetes"
	"github.com/k8snetworkplumbingwg/sriov-network-operator/pkg/platforms/orchestrator/openshift"
	"github.com/k8snetworkplumbingwg/sriov-network-operator/pkg/vars"
)

//go:generate ../../../bin/mockgen -destination mock/mock_orchestration.go -source orchestration.go
type OrchestrationInterface interface {
	// ClusterType represent to a cluster type
	ClusterType() consts.ClusterType
	// Flavor represents the internal flavor of the cluster
	Flavor() consts.ClusterFlavor
	// BeforeDrainNode run a platform-specific logic before draining the node
	BeforeDrainNode(context.Context, *corev1.Node) (bool, error)
	// AfterCompleteDrainNode run a platform-specific logic after node draining was completed
	AfterCompleteDrainNode(context.Context, *corev1.Node) (bool, error)
}

func New() (OrchestrationInterface, error) {
	if vars.ClusterType == consts.ClusterTypeOpenshift {
		return openshift.New()
	} else if vars.ClusterType == consts.ClusterTypeKubernetes {
		return kubernetes.New()
	}

	return nil, fmt.Errorf("unknown orchestration type: %s", vars.ClusterType)
}
