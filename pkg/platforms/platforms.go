package platforms

import (
	"sigs.k8s.io/controller-runtime/pkg/log"

	"github.com/k8snetworkplumbingwg/sriov-network-operator/pkg/host"
	"github.com/k8snetworkplumbingwg/sriov-network-operator/pkg/platforms/hypervisor/openstack"
	"github.com/k8snetworkplumbingwg/sriov-network-operator/pkg/platforms/orchestrator"
)

type PlatformHelper struct {
	Orchestrator orchestrator.OrchestrationInterface
	Hypervisor   openstack.OpenstackInterface
}

func NewDefaultPlatformHelper() (*PlatformHelper, error) {
	orch, err := orchestrator.New()
	if err != nil {
		return nil, err
	}

	hostManager, err := host.NewDefaultHostManager()
	if err != nil {
		log.Log.Error(err, "failed to create host manager")
		return nil, err
	}
	openstackContext := openstack.New(hostManager)

	return &PlatformHelper{
		orch,
		openstackContext,
	}, nil
}
