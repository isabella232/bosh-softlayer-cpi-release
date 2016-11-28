// This file was generated by counterfeiter
package fakes

import (
	"sync"

	"github.com/cloudfoundry/bosh-softlayer-cpi/softlayer/common"
	"github.com/cloudfoundry/bosh-softlayer-cpi/softlayer/stemcell"
)

type FakeVMCreator struct {
	CreateStub        func(string, stemcell.Stemcell, common.VMCloudProperties, common.Networks, common.Environment) (common.VM, error)
	createMutex       sync.RWMutex
	createArgsForCall []struct {
		arg1 string
		arg2 stemcell.Stemcell
		arg3 common.VMCloudProperties
		arg4 common.Networks
		arg5 common.Environment
	}
	createReturns struct {
		result1 common.VM
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeVMCreator) Create(arg1 string, arg2 stemcell.Stemcell, arg3 common.VMCloudProperties, arg4 common.Networks, arg5 common.Environment) (common.VM, error) {
	fake.createMutex.Lock()
	fake.createArgsForCall = append(fake.createArgsForCall, struct {
		arg1 string
		arg2 stemcell.Stemcell
		arg3 common.VMCloudProperties
		arg4 common.Networks
		arg5 common.Environment
	}{arg1, arg2, arg3, arg4, arg5})
	fake.recordInvocation("Create", []interface{}{arg1, arg2, arg3, arg4, arg5})
	fake.createMutex.Unlock()
	if fake.CreateStub != nil {
		return fake.CreateStub(arg1, arg2, arg3, arg4, arg5)
	} else {
		return fake.createReturns.result1, fake.createReturns.result2
	}
}

func (fake *FakeVMCreator) CreateCallCount() int {
	fake.createMutex.RLock()
	defer fake.createMutex.RUnlock()
	return len(fake.createArgsForCall)
}

func (fake *FakeVMCreator) CreateArgsForCall(i int) (string, stemcell.Stemcell, common.VMCloudProperties, common.Networks, common.Environment) {
	fake.createMutex.RLock()
	defer fake.createMutex.RUnlock()
	return fake.createArgsForCall[i].arg1, fake.createArgsForCall[i].arg2, fake.createArgsForCall[i].arg3, fake.createArgsForCall[i].arg4, fake.createArgsForCall[i].arg5
}

func (fake *FakeVMCreator) CreateReturns(result1 common.VM, result2 error) {
	fake.CreateStub = nil
	fake.createReturns = struct {
		result1 common.VM
		result2 error
	}{result1, result2}
}

func (fake *FakeVMCreator) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.createMutex.RLock()
	defer fake.createMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeVMCreator) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ common.VMCreator = new(FakeVMCreator)
