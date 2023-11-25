package consts

import "time"

const (
	ResyncPeriod                       = 5 * time.Minute
	DefaultConfigName                  = "default"
	ConfigDaemonPath                   = "./bindata/manifests/daemon"
	InjectorWebHookPath                = "./bindata/manifests/webhook"
	OperatorWebHookPath                = "./bindata/manifests/operator-webhook"
	SystemdServiceOcpPath              = "./bindata/manifests/sriov-config-service/openshift"
	SystemdServiceOcpMachineConfigName = "sriov-config-service"
	ServiceCAConfigMapAnnotation       = "service.beta.openshift.io/inject-cabundle"
	InjectorWebHookName                = "network-resources-injector-config"
	OperatorWebHookName                = "sriov-operator-webhook-config"
	DeprecatedOperatorWebHookName      = "operator-webhook-config"
	PluginPath                         = "./bindata/manifests/plugins"
	DaemonPath                         = "./bindata/manifests/daemon"
	DefaultPolicyName                  = "default"
	ConfigMapName                      = "device-plugin-config"
	DaemonSet                          = "DaemonSet"
	ServiceAccount                     = "ServiceAccount"
	DPConfigFileName                   = "config.json"
	OVSHWOLMachineConfigNameSuffix     = "ovs-hw-offload"

	NodeDrainAnnotation             = "sriovnetwork.openshift.io/state"
	NodeStateDrainAnnotationCurrent = "sriovnetwork.openshift.io/current-state"
	DrainIdle                       = "Idle"
	DrainRequired                   = "Drain_Required"
	DrainMcpPaused                  = "Draining_MCP_Paused"
	Draining                        = "Draining"
	DrainComplete                   = "DrainComplete"

	LinkTypeEthernet   = "ether"
	LinkTypeInfiniband = "infiniband"

	LinkTypeIB  = "IB"
	LinkTypeETH = "ETH"

	DeviceTypeVfioPci   = "vfio-pci"
	DeviceTypeNetDevice = "netdevice"
	VdpaTypeVirtio      = "virtio"
	VdpaTypeVhost       = "vhost"
)
