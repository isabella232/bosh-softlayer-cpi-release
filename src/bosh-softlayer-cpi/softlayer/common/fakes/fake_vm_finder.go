// This file was generated by counterfeiter
package fakes

import (
	"sync"

	"bosh-softlayer-cpi/softlayer/common"
)

type FakeVMFinder struct {
	FindStub        func(int) (common.VM, bool, error)
	findMutex       sync.RWMutex
	findArgsForCall []struct {
		arg1 int
	}
	findReturns struct {
		result1 common.VM
		result2 bool
		result3 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeVMFinder) Find(arg1 int) (common.VM, bool, error) {
	fake.findMutex.Lock()
	fake.findArgsForCall = append(fake.findArgsForCall, struct {
		arg1 int
	}{arg1})
	fake.recordInvocation("Find", []interface{}{arg1})
	fake.findMutex.Unlock()
	if fake.FindStub != nil {
		return fake.FindStub(arg1)
	} else {
		return fake.findReturns.result1, fake.findReturns.result2, fake.findReturns.result3
	}
}

func (fake *FakeVMFinder) FindCallCount() int {
	fake.findMutex.RLock()
	defer fake.findMutex.RUnlock()
	return len(fake.findArgsForCall)
}

func (fake *FakeVMFinder) FindArgsForCall(i int) int {
	fake.findMutex.RLock()
	defer fake.findMutex.RUnlock()
	return fake.findArgsForCall[i].arg1
}

func (fake *FakeVMFinder) FindReturns(result1 common.VM, result2 bool, result3 error) {
	fake.FindStub = nil
	fake.findReturns = struct {
		result1 common.VM
		result2 bool
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeVMFinder) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.findMutex.RLock()
	defer fake.findMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeVMFinder) recordInvocation(key string, args []interface{}) {
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

var _ common.VMFinder = new(FakeVMFinder)