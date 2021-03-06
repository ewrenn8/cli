// Code generated by counterfeiter. DO NOT EDIT.
package v3fakes

import (
	"sync"

	"code.cloudfoundry.org/cli/actor/v2v3action"
	"code.cloudfoundry.org/cli/command/v3"
)

type FakeShareServiceActor struct {
	ShareServiceInstanceToSpaceNameByNameAndSpaceAndOrganizationNameStub        func(sharedToSpaceName string, serviceInstanceName string, sourceSpaceGUID string, sharedToOrgName string) (v2v3action.Warnings, error)
	shareServiceInstanceToSpaceNameByNameAndSpaceAndOrganizationNameMutex       sync.RWMutex
	shareServiceInstanceToSpaceNameByNameAndSpaceAndOrganizationNameArgsForCall []struct {
		sharedToSpaceName   string
		serviceInstanceName string
		sourceSpaceGUID     string
		sharedToOrgName     string
	}
	shareServiceInstanceToSpaceNameByNameAndSpaceAndOrganizationNameReturns struct {
		result1 v2v3action.Warnings
		result2 error
	}
	shareServiceInstanceToSpaceNameByNameAndSpaceAndOrganizationNameReturnsOnCall map[int]struct {
		result1 v2v3action.Warnings
		result2 error
	}
	ShareServiceInstanceToSpaceNameByNameAndSpaceAndOrganizationStub        func(sharedToSpaceName string, serviceInstanceName string, sourceSpaceGUID string, sharedToOrgGUID string) (v2v3action.Warnings, error)
	shareServiceInstanceToSpaceNameByNameAndSpaceAndOrganizationMutex       sync.RWMutex
	shareServiceInstanceToSpaceNameByNameAndSpaceAndOrganizationArgsForCall []struct {
		sharedToSpaceName   string
		serviceInstanceName string
		sourceSpaceGUID     string
		sharedToOrgGUID     string
	}
	shareServiceInstanceToSpaceNameByNameAndSpaceAndOrganizationReturns struct {
		result1 v2v3action.Warnings
		result2 error
	}
	shareServiceInstanceToSpaceNameByNameAndSpaceAndOrganizationReturnsOnCall map[int]struct {
		result1 v2v3action.Warnings
		result2 error
	}
	CloudControllerV3APIVersionStub        func() string
	cloudControllerV3APIVersionMutex       sync.RWMutex
	cloudControllerV3APIVersionArgsForCall []struct{}
	cloudControllerV3APIVersionReturns     struct {
		result1 string
	}
	cloudControllerV3APIVersionReturnsOnCall map[int]struct {
		result1 string
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeShareServiceActor) ShareServiceInstanceToSpaceNameByNameAndSpaceAndOrganizationName(sharedToSpaceName string, serviceInstanceName string, sourceSpaceGUID string, sharedToOrgName string) (v2v3action.Warnings, error) {
	fake.shareServiceInstanceToSpaceNameByNameAndSpaceAndOrganizationNameMutex.Lock()
	ret, specificReturn := fake.shareServiceInstanceToSpaceNameByNameAndSpaceAndOrganizationNameReturnsOnCall[len(fake.shareServiceInstanceToSpaceNameByNameAndSpaceAndOrganizationNameArgsForCall)]
	fake.shareServiceInstanceToSpaceNameByNameAndSpaceAndOrganizationNameArgsForCall = append(fake.shareServiceInstanceToSpaceNameByNameAndSpaceAndOrganizationNameArgsForCall, struct {
		sharedToSpaceName   string
		serviceInstanceName string
		sourceSpaceGUID     string
		sharedToOrgName     string
	}{sharedToSpaceName, serviceInstanceName, sourceSpaceGUID, sharedToOrgName})
	fake.recordInvocation("ShareServiceInstanceToSpaceNameByNameAndSpaceAndOrganizationName", []interface{}{sharedToSpaceName, serviceInstanceName, sourceSpaceGUID, sharedToOrgName})
	fake.shareServiceInstanceToSpaceNameByNameAndSpaceAndOrganizationNameMutex.Unlock()
	if fake.ShareServiceInstanceToSpaceNameByNameAndSpaceAndOrganizationNameStub != nil {
		return fake.ShareServiceInstanceToSpaceNameByNameAndSpaceAndOrganizationNameStub(sharedToSpaceName, serviceInstanceName, sourceSpaceGUID, sharedToOrgName)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.shareServiceInstanceToSpaceNameByNameAndSpaceAndOrganizationNameReturns.result1, fake.shareServiceInstanceToSpaceNameByNameAndSpaceAndOrganizationNameReturns.result2
}

func (fake *FakeShareServiceActor) ShareServiceInstanceToSpaceNameByNameAndSpaceAndOrganizationNameCallCount() int {
	fake.shareServiceInstanceToSpaceNameByNameAndSpaceAndOrganizationNameMutex.RLock()
	defer fake.shareServiceInstanceToSpaceNameByNameAndSpaceAndOrganizationNameMutex.RUnlock()
	return len(fake.shareServiceInstanceToSpaceNameByNameAndSpaceAndOrganizationNameArgsForCall)
}

func (fake *FakeShareServiceActor) ShareServiceInstanceToSpaceNameByNameAndSpaceAndOrganizationNameArgsForCall(i int) (string, string, string, string) {
	fake.shareServiceInstanceToSpaceNameByNameAndSpaceAndOrganizationNameMutex.RLock()
	defer fake.shareServiceInstanceToSpaceNameByNameAndSpaceAndOrganizationNameMutex.RUnlock()
	return fake.shareServiceInstanceToSpaceNameByNameAndSpaceAndOrganizationNameArgsForCall[i].sharedToSpaceName, fake.shareServiceInstanceToSpaceNameByNameAndSpaceAndOrganizationNameArgsForCall[i].serviceInstanceName, fake.shareServiceInstanceToSpaceNameByNameAndSpaceAndOrganizationNameArgsForCall[i].sourceSpaceGUID, fake.shareServiceInstanceToSpaceNameByNameAndSpaceAndOrganizationNameArgsForCall[i].sharedToOrgName
}

func (fake *FakeShareServiceActor) ShareServiceInstanceToSpaceNameByNameAndSpaceAndOrganizationNameReturns(result1 v2v3action.Warnings, result2 error) {
	fake.ShareServiceInstanceToSpaceNameByNameAndSpaceAndOrganizationNameStub = nil
	fake.shareServiceInstanceToSpaceNameByNameAndSpaceAndOrganizationNameReturns = struct {
		result1 v2v3action.Warnings
		result2 error
	}{result1, result2}
}

func (fake *FakeShareServiceActor) ShareServiceInstanceToSpaceNameByNameAndSpaceAndOrganizationNameReturnsOnCall(i int, result1 v2v3action.Warnings, result2 error) {
	fake.ShareServiceInstanceToSpaceNameByNameAndSpaceAndOrganizationNameStub = nil
	if fake.shareServiceInstanceToSpaceNameByNameAndSpaceAndOrganizationNameReturnsOnCall == nil {
		fake.shareServiceInstanceToSpaceNameByNameAndSpaceAndOrganizationNameReturnsOnCall = make(map[int]struct {
			result1 v2v3action.Warnings
			result2 error
		})
	}
	fake.shareServiceInstanceToSpaceNameByNameAndSpaceAndOrganizationNameReturnsOnCall[i] = struct {
		result1 v2v3action.Warnings
		result2 error
	}{result1, result2}
}

func (fake *FakeShareServiceActor) ShareServiceInstanceToSpaceNameByNameAndSpaceAndOrganization(sharedToSpaceName string, serviceInstanceName string, sourceSpaceGUID string, sharedToOrgGUID string) (v2v3action.Warnings, error) {
	fake.shareServiceInstanceToSpaceNameByNameAndSpaceAndOrganizationMutex.Lock()
	ret, specificReturn := fake.shareServiceInstanceToSpaceNameByNameAndSpaceAndOrganizationReturnsOnCall[len(fake.shareServiceInstanceToSpaceNameByNameAndSpaceAndOrganizationArgsForCall)]
	fake.shareServiceInstanceToSpaceNameByNameAndSpaceAndOrganizationArgsForCall = append(fake.shareServiceInstanceToSpaceNameByNameAndSpaceAndOrganizationArgsForCall, struct {
		sharedToSpaceName   string
		serviceInstanceName string
		sourceSpaceGUID     string
		sharedToOrgGUID     string
	}{sharedToSpaceName, serviceInstanceName, sourceSpaceGUID, sharedToOrgGUID})
	fake.recordInvocation("ShareServiceInstanceToSpaceNameByNameAndSpaceAndOrganization", []interface{}{sharedToSpaceName, serviceInstanceName, sourceSpaceGUID, sharedToOrgGUID})
	fake.shareServiceInstanceToSpaceNameByNameAndSpaceAndOrganizationMutex.Unlock()
	if fake.ShareServiceInstanceToSpaceNameByNameAndSpaceAndOrganizationStub != nil {
		return fake.ShareServiceInstanceToSpaceNameByNameAndSpaceAndOrganizationStub(sharedToSpaceName, serviceInstanceName, sourceSpaceGUID, sharedToOrgGUID)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.shareServiceInstanceToSpaceNameByNameAndSpaceAndOrganizationReturns.result1, fake.shareServiceInstanceToSpaceNameByNameAndSpaceAndOrganizationReturns.result2
}

func (fake *FakeShareServiceActor) ShareServiceInstanceToSpaceNameByNameAndSpaceAndOrganizationCallCount() int {
	fake.shareServiceInstanceToSpaceNameByNameAndSpaceAndOrganizationMutex.RLock()
	defer fake.shareServiceInstanceToSpaceNameByNameAndSpaceAndOrganizationMutex.RUnlock()
	return len(fake.shareServiceInstanceToSpaceNameByNameAndSpaceAndOrganizationArgsForCall)
}

func (fake *FakeShareServiceActor) ShareServiceInstanceToSpaceNameByNameAndSpaceAndOrganizationArgsForCall(i int) (string, string, string, string) {
	fake.shareServiceInstanceToSpaceNameByNameAndSpaceAndOrganizationMutex.RLock()
	defer fake.shareServiceInstanceToSpaceNameByNameAndSpaceAndOrganizationMutex.RUnlock()
	return fake.shareServiceInstanceToSpaceNameByNameAndSpaceAndOrganizationArgsForCall[i].sharedToSpaceName, fake.shareServiceInstanceToSpaceNameByNameAndSpaceAndOrganizationArgsForCall[i].serviceInstanceName, fake.shareServiceInstanceToSpaceNameByNameAndSpaceAndOrganizationArgsForCall[i].sourceSpaceGUID, fake.shareServiceInstanceToSpaceNameByNameAndSpaceAndOrganizationArgsForCall[i].sharedToOrgGUID
}

func (fake *FakeShareServiceActor) ShareServiceInstanceToSpaceNameByNameAndSpaceAndOrganizationReturns(result1 v2v3action.Warnings, result2 error) {
	fake.ShareServiceInstanceToSpaceNameByNameAndSpaceAndOrganizationStub = nil
	fake.shareServiceInstanceToSpaceNameByNameAndSpaceAndOrganizationReturns = struct {
		result1 v2v3action.Warnings
		result2 error
	}{result1, result2}
}

func (fake *FakeShareServiceActor) ShareServiceInstanceToSpaceNameByNameAndSpaceAndOrganizationReturnsOnCall(i int, result1 v2v3action.Warnings, result2 error) {
	fake.ShareServiceInstanceToSpaceNameByNameAndSpaceAndOrganizationStub = nil
	if fake.shareServiceInstanceToSpaceNameByNameAndSpaceAndOrganizationReturnsOnCall == nil {
		fake.shareServiceInstanceToSpaceNameByNameAndSpaceAndOrganizationReturnsOnCall = make(map[int]struct {
			result1 v2v3action.Warnings
			result2 error
		})
	}
	fake.shareServiceInstanceToSpaceNameByNameAndSpaceAndOrganizationReturnsOnCall[i] = struct {
		result1 v2v3action.Warnings
		result2 error
	}{result1, result2}
}

func (fake *FakeShareServiceActor) CloudControllerV3APIVersion() string {
	fake.cloudControllerV3APIVersionMutex.Lock()
	ret, specificReturn := fake.cloudControllerV3APIVersionReturnsOnCall[len(fake.cloudControllerV3APIVersionArgsForCall)]
	fake.cloudControllerV3APIVersionArgsForCall = append(fake.cloudControllerV3APIVersionArgsForCall, struct{}{})
	fake.recordInvocation("CloudControllerV3APIVersion", []interface{}{})
	fake.cloudControllerV3APIVersionMutex.Unlock()
	if fake.CloudControllerV3APIVersionStub != nil {
		return fake.CloudControllerV3APIVersionStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.cloudControllerV3APIVersionReturns.result1
}

func (fake *FakeShareServiceActor) CloudControllerV3APIVersionCallCount() int {
	fake.cloudControllerV3APIVersionMutex.RLock()
	defer fake.cloudControllerV3APIVersionMutex.RUnlock()
	return len(fake.cloudControllerV3APIVersionArgsForCall)
}

func (fake *FakeShareServiceActor) CloudControllerV3APIVersionReturns(result1 string) {
	fake.CloudControllerV3APIVersionStub = nil
	fake.cloudControllerV3APIVersionReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeShareServiceActor) CloudControllerV3APIVersionReturnsOnCall(i int, result1 string) {
	fake.CloudControllerV3APIVersionStub = nil
	if fake.cloudControllerV3APIVersionReturnsOnCall == nil {
		fake.cloudControllerV3APIVersionReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.cloudControllerV3APIVersionReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *FakeShareServiceActor) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.shareServiceInstanceToSpaceNameByNameAndSpaceAndOrganizationNameMutex.RLock()
	defer fake.shareServiceInstanceToSpaceNameByNameAndSpaceAndOrganizationNameMutex.RUnlock()
	fake.shareServiceInstanceToSpaceNameByNameAndSpaceAndOrganizationMutex.RLock()
	defer fake.shareServiceInstanceToSpaceNameByNameAndSpaceAndOrganizationMutex.RUnlock()
	fake.cloudControllerV3APIVersionMutex.RLock()
	defer fake.cloudControllerV3APIVersionMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeShareServiceActor) recordInvocation(key string, args []interface{}) {
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

var _ v3.ShareServiceActor = new(FakeShareServiceActor)
