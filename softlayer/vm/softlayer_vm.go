package vm

import (
	bosherr "bosh/errors"
	boshlog "bosh/logger"

	sl "github.com/maximilien/softlayer-go/softlayer"

	bslcdisk "github.com/maximilien/bosh-softlayer-cpi/softlayer/disk"
)

const softLayerVMtag = "SoftLayerVM"

type SoftLayerVM struct {
	id int

	softLayerClient sl.Client
	agentEnvService AgentEnvService

	logger boshlog.Logger
}

func NewSoftLayerVM(
	id int,
	softLayerClient sl.Client,
	agentEnvService AgentEnvService,
	logger boshlog.Logger,
) SoftLayerVM {
	return SoftLayerVM{
		id: id,

		softLayerClient: softLayerClient,
		agentEnvService: agentEnvService,

		logger: logger,
	}
}

func (vm SoftLayerVM) ID() int { return vm.id }

func (vm SoftLayerVM) Delete() error {
	virtualGuestService, err := vm.softLayerClient.GetSoftLayer_Virtual_Guest_Service()
	if err != nil {
		return bosherr.WrapError(err, "Creating SoftLayer VirtualGuestService from client")
	}

	deleted, err := virtualGuestService.DeleteObject(vm.ID())
	if err != nil {
		return bosherr.WrapError(err, "Deleting SoftLayer VirtualGuest from client")
	}

	if !deleted {
		return bosherr.WrapError(nil, "Did not delete SoftLayer VirtualGuest from client")
	}

	return nil
}

func (vm SoftLayerVM) AttachDisk(disk bslcdisk.Disk) error {
	vm.logger.Info(softLayerVMtag, "Not yet implemented!")

	return nil
}

func (vm SoftLayerVM) DetachDisk(disk bslcdisk.Disk) error {
	vm.logger.Info(softLayerVMtag, "Not yet implemented!")
	
	return nil
}