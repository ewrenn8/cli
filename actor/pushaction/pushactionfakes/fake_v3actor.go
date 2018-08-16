// Code generated by counterfeiter. DO NOT EDIT.
package pushactionfakes

import (
	"io"
	"sync"

	"code.cloudfoundry.org/cli/actor/pushaction"
	"code.cloudfoundry.org/cli/actor/sharedaction"
	"code.cloudfoundry.org/cli/actor/v3action"
)

type FakeV3Actor struct {
	CloudControllerAPIVersionStub        func() string
	cloudControllerAPIVersionMutex       sync.RWMutex
	cloudControllerAPIVersionArgsForCall []struct{}
	cloudControllerAPIVersionReturns     struct {
		result1 string
	}
	cloudControllerAPIVersionReturnsOnCall map[int]struct {
		result1 string
	}
	CreateApplicationDeploymentStub        func(appGUID string) (v3action.Warnings, error)
	createApplicationDeploymentMutex       sync.RWMutex
	createApplicationDeploymentArgsForCall []struct {
		appGUID string
	}
	createApplicationDeploymentReturns struct {
		result1 v3action.Warnings
		result2 error
	}
	createApplicationDeploymentReturnsOnCall map[int]struct {
		result1 v3action.Warnings
		result2 error
	}
	CreateApplicationInSpaceStub        func(app v3action.Application, spaceGUID string) (v3action.Application, v3action.Warnings, error)
	createApplicationInSpaceMutex       sync.RWMutex
	createApplicationInSpaceArgsForCall []struct {
		app       v3action.Application
		spaceGUID string
	}
	createApplicationInSpaceReturns struct {
		result1 v3action.Application
		result2 v3action.Warnings
		result3 error
	}
	createApplicationInSpaceReturnsOnCall map[int]struct {
		result1 v3action.Application
		result2 v3action.Warnings
		result3 error
	}
	CreateBitsPackageByApplicationStub        func(appGUID string) (v3action.Package, v3action.Warnings, error)
	createBitsPackageByApplicationMutex       sync.RWMutex
	createBitsPackageByApplicationArgsForCall []struct {
		appGUID string
	}
	createBitsPackageByApplicationReturns struct {
		result1 v3action.Package
		result2 v3action.Warnings
		result3 error
	}
	createBitsPackageByApplicationReturnsOnCall map[int]struct {
		result1 v3action.Package
		result2 v3action.Warnings
		result3 error
	}
	GetApplicationByNameAndSpaceStub        func(appName string, spaceGUID string) (v3action.Application, v3action.Warnings, error)
	getApplicationByNameAndSpaceMutex       sync.RWMutex
	getApplicationByNameAndSpaceArgsForCall []struct {
		appName   string
		spaceGUID string
	}
	getApplicationByNameAndSpaceReturns struct {
		result1 v3action.Application
		result2 v3action.Warnings
		result3 error
	}
	getApplicationByNameAndSpaceReturnsOnCall map[int]struct {
		result1 v3action.Application
		result2 v3action.Warnings
		result3 error
	}
	PollPackageStub        func(pkg v3action.Package) (v3action.Package, v3action.Warnings, error)
	pollPackageMutex       sync.RWMutex
	pollPackageArgsForCall []struct {
		pkg v3action.Package
	}
	pollPackageReturns struct {
		result1 v3action.Package
		result2 v3action.Warnings
		result3 error
	}
	pollPackageReturnsOnCall map[int]struct {
		result1 v3action.Package
		result2 v3action.Warnings
		result3 error
	}
	SetApplicationDropletStub        func(appGUID string, dropletGUID string) (v3action.Warnings, error)
	setApplicationDropletMutex       sync.RWMutex
	setApplicationDropletArgsForCall []struct {
		appGUID     string
		dropletGUID string
	}
	setApplicationDropletReturns struct {
		result1 v3action.Warnings
		result2 error
	}
	setApplicationDropletReturnsOnCall map[int]struct {
		result1 v3action.Warnings
		result2 error
	}
	StageApplicationPackageStub        func(pkgGUID string) (v3action.Build, v3action.Warnings, error)
	stageApplicationPackageMutex       sync.RWMutex
	stageApplicationPackageArgsForCall []struct {
		pkgGUID string
	}
	stageApplicationPackageReturns struct {
		result1 v3action.Build
		result2 v3action.Warnings
		result3 error
	}
	stageApplicationPackageReturnsOnCall map[int]struct {
		result1 v3action.Build
		result2 v3action.Warnings
		result3 error
	}
	PollBuildStub        func(buildGUID string, appName string) (v3action.Droplet, v3action.Warnings, error)
	pollBuildMutex       sync.RWMutex
	pollBuildArgsForCall []struct {
		buildGUID string
		appName   string
	}
	pollBuildReturns struct {
		result1 v3action.Droplet
		result2 v3action.Warnings
		result3 error
	}
	pollBuildReturnsOnCall map[int]struct {
		result1 v3action.Droplet
		result2 v3action.Warnings
		result3 error
	}
	UpdateApplicationStub        func(v3action.Application) (v3action.Application, v3action.Warnings, error)
	updateApplicationMutex       sync.RWMutex
	updateApplicationArgsForCall []struct {
		arg1 v3action.Application
	}
	updateApplicationReturns struct {
		result1 v3action.Application
		result2 v3action.Warnings
		result3 error
	}
	updateApplicationReturnsOnCall map[int]struct {
		result1 v3action.Application
		result2 v3action.Warnings
		result3 error
	}
	UploadBitsPackageStub        func(v3action.Package, []sharedaction.Resource, io.Reader, int64) (v3action.Package, v3action.Warnings, error)
	uploadBitsPackageMutex       sync.RWMutex
	uploadBitsPackageArgsForCall []struct {
		arg1 v3action.Package
		arg2 []sharedaction.Resource
		arg3 io.Reader
		arg4 int64
	}
	uploadBitsPackageReturns struct {
		result1 v3action.Package
		result2 v3action.Warnings
		result3 error
	}
	uploadBitsPackageReturnsOnCall map[int]struct {
		result1 v3action.Package
		result2 v3action.Warnings
		result3 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeV3Actor) CloudControllerAPIVersion() string {
	fake.cloudControllerAPIVersionMutex.Lock()
	ret, specificReturn := fake.cloudControllerAPIVersionReturnsOnCall[len(fake.cloudControllerAPIVersionArgsForCall)]
	fake.cloudControllerAPIVersionArgsForCall = append(fake.cloudControllerAPIVersionArgsForCall, struct{}{})
	fake.recordInvocation("CloudControllerAPIVersion", []interface{}{})
	fake.cloudControllerAPIVersionMutex.Unlock()
	if fake.CloudControllerAPIVersionStub != nil {
		return fake.CloudControllerAPIVersionStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.cloudControllerAPIVersionReturns.result1
}

func (fake *FakeV3Actor) CloudControllerAPIVersionCallCount() int {
	fake.cloudControllerAPIVersionMutex.RLock()
	defer fake.cloudControllerAPIVersionMutex.RUnlock()
	return len(fake.cloudControllerAPIVersionArgsForCall)
}

func (fake *FakeV3Actor) CloudControllerAPIVersionReturns(result1 string) {
	fake.CloudControllerAPIVersionStub = nil
	fake.cloudControllerAPIVersionReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeV3Actor) CloudControllerAPIVersionReturnsOnCall(i int, result1 string) {
	fake.CloudControllerAPIVersionStub = nil
	if fake.cloudControllerAPIVersionReturnsOnCall == nil {
		fake.cloudControllerAPIVersionReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.cloudControllerAPIVersionReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *FakeV3Actor) CreateApplicationDeployment(appGUID string) (v3action.Warnings, error) {
	fake.createApplicationDeploymentMutex.Lock()
	ret, specificReturn := fake.createApplicationDeploymentReturnsOnCall[len(fake.createApplicationDeploymentArgsForCall)]
	fake.createApplicationDeploymentArgsForCall = append(fake.createApplicationDeploymentArgsForCall, struct {
		appGUID string
	}{appGUID})
	fake.recordInvocation("CreateApplicationDeployment", []interface{}{appGUID})
	fake.createApplicationDeploymentMutex.Unlock()
	if fake.CreateApplicationDeploymentStub != nil {
		return fake.CreateApplicationDeploymentStub(appGUID)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.createApplicationDeploymentReturns.result1, fake.createApplicationDeploymentReturns.result2
}

func (fake *FakeV3Actor) CreateApplicationDeploymentCallCount() int {
	fake.createApplicationDeploymentMutex.RLock()
	defer fake.createApplicationDeploymentMutex.RUnlock()
	return len(fake.createApplicationDeploymentArgsForCall)
}

func (fake *FakeV3Actor) CreateApplicationDeploymentArgsForCall(i int) string {
	fake.createApplicationDeploymentMutex.RLock()
	defer fake.createApplicationDeploymentMutex.RUnlock()
	return fake.createApplicationDeploymentArgsForCall[i].appGUID
}

func (fake *FakeV3Actor) CreateApplicationDeploymentReturns(result1 v3action.Warnings, result2 error) {
	fake.CreateApplicationDeploymentStub = nil
	fake.createApplicationDeploymentReturns = struct {
		result1 v3action.Warnings
		result2 error
	}{result1, result2}
}

func (fake *FakeV3Actor) CreateApplicationDeploymentReturnsOnCall(i int, result1 v3action.Warnings, result2 error) {
	fake.CreateApplicationDeploymentStub = nil
	if fake.createApplicationDeploymentReturnsOnCall == nil {
		fake.createApplicationDeploymentReturnsOnCall = make(map[int]struct {
			result1 v3action.Warnings
			result2 error
		})
	}
	fake.createApplicationDeploymentReturnsOnCall[i] = struct {
		result1 v3action.Warnings
		result2 error
	}{result1, result2}
}

func (fake *FakeV3Actor) CreateApplicationInSpace(app v3action.Application, spaceGUID string) (v3action.Application, v3action.Warnings, error) {
	fake.createApplicationInSpaceMutex.Lock()
	ret, specificReturn := fake.createApplicationInSpaceReturnsOnCall[len(fake.createApplicationInSpaceArgsForCall)]
	fake.createApplicationInSpaceArgsForCall = append(fake.createApplicationInSpaceArgsForCall, struct {
		app       v3action.Application
		spaceGUID string
	}{app, spaceGUID})
	fake.recordInvocation("CreateApplicationInSpace", []interface{}{app, spaceGUID})
	fake.createApplicationInSpaceMutex.Unlock()
	if fake.CreateApplicationInSpaceStub != nil {
		return fake.CreateApplicationInSpaceStub(app, spaceGUID)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	return fake.createApplicationInSpaceReturns.result1, fake.createApplicationInSpaceReturns.result2, fake.createApplicationInSpaceReturns.result3
}

func (fake *FakeV3Actor) CreateApplicationInSpaceCallCount() int {
	fake.createApplicationInSpaceMutex.RLock()
	defer fake.createApplicationInSpaceMutex.RUnlock()
	return len(fake.createApplicationInSpaceArgsForCall)
}

func (fake *FakeV3Actor) CreateApplicationInSpaceArgsForCall(i int) (v3action.Application, string) {
	fake.createApplicationInSpaceMutex.RLock()
	defer fake.createApplicationInSpaceMutex.RUnlock()
	return fake.createApplicationInSpaceArgsForCall[i].app, fake.createApplicationInSpaceArgsForCall[i].spaceGUID
}

func (fake *FakeV3Actor) CreateApplicationInSpaceReturns(result1 v3action.Application, result2 v3action.Warnings, result3 error) {
	fake.CreateApplicationInSpaceStub = nil
	fake.createApplicationInSpaceReturns = struct {
		result1 v3action.Application
		result2 v3action.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeV3Actor) CreateApplicationInSpaceReturnsOnCall(i int, result1 v3action.Application, result2 v3action.Warnings, result3 error) {
	fake.CreateApplicationInSpaceStub = nil
	if fake.createApplicationInSpaceReturnsOnCall == nil {
		fake.createApplicationInSpaceReturnsOnCall = make(map[int]struct {
			result1 v3action.Application
			result2 v3action.Warnings
			result3 error
		})
	}
	fake.createApplicationInSpaceReturnsOnCall[i] = struct {
		result1 v3action.Application
		result2 v3action.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeV3Actor) CreateBitsPackageByApplication(appGUID string) (v3action.Package, v3action.Warnings, error) {
	fake.createBitsPackageByApplicationMutex.Lock()
	ret, specificReturn := fake.createBitsPackageByApplicationReturnsOnCall[len(fake.createBitsPackageByApplicationArgsForCall)]
	fake.createBitsPackageByApplicationArgsForCall = append(fake.createBitsPackageByApplicationArgsForCall, struct {
		appGUID string
	}{appGUID})
	fake.recordInvocation("CreateBitsPackageByApplication", []interface{}{appGUID})
	fake.createBitsPackageByApplicationMutex.Unlock()
	if fake.CreateBitsPackageByApplicationStub != nil {
		return fake.CreateBitsPackageByApplicationStub(appGUID)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	return fake.createBitsPackageByApplicationReturns.result1, fake.createBitsPackageByApplicationReturns.result2, fake.createBitsPackageByApplicationReturns.result3
}

func (fake *FakeV3Actor) CreateBitsPackageByApplicationCallCount() int {
	fake.createBitsPackageByApplicationMutex.RLock()
	defer fake.createBitsPackageByApplicationMutex.RUnlock()
	return len(fake.createBitsPackageByApplicationArgsForCall)
}

func (fake *FakeV3Actor) CreateBitsPackageByApplicationArgsForCall(i int) string {
	fake.createBitsPackageByApplicationMutex.RLock()
	defer fake.createBitsPackageByApplicationMutex.RUnlock()
	return fake.createBitsPackageByApplicationArgsForCall[i].appGUID
}

func (fake *FakeV3Actor) CreateBitsPackageByApplicationReturns(result1 v3action.Package, result2 v3action.Warnings, result3 error) {
	fake.CreateBitsPackageByApplicationStub = nil
	fake.createBitsPackageByApplicationReturns = struct {
		result1 v3action.Package
		result2 v3action.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeV3Actor) CreateBitsPackageByApplicationReturnsOnCall(i int, result1 v3action.Package, result2 v3action.Warnings, result3 error) {
	fake.CreateBitsPackageByApplicationStub = nil
	if fake.createBitsPackageByApplicationReturnsOnCall == nil {
		fake.createBitsPackageByApplicationReturnsOnCall = make(map[int]struct {
			result1 v3action.Package
			result2 v3action.Warnings
			result3 error
		})
	}
	fake.createBitsPackageByApplicationReturnsOnCall[i] = struct {
		result1 v3action.Package
		result2 v3action.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeV3Actor) GetApplicationByNameAndSpace(appName string, spaceGUID string) (v3action.Application, v3action.Warnings, error) {
	fake.getApplicationByNameAndSpaceMutex.Lock()
	ret, specificReturn := fake.getApplicationByNameAndSpaceReturnsOnCall[len(fake.getApplicationByNameAndSpaceArgsForCall)]
	fake.getApplicationByNameAndSpaceArgsForCall = append(fake.getApplicationByNameAndSpaceArgsForCall, struct {
		appName   string
		spaceGUID string
	}{appName, spaceGUID})
	fake.recordInvocation("GetApplicationByNameAndSpace", []interface{}{appName, spaceGUID})
	fake.getApplicationByNameAndSpaceMutex.Unlock()
	if fake.GetApplicationByNameAndSpaceStub != nil {
		return fake.GetApplicationByNameAndSpaceStub(appName, spaceGUID)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	return fake.getApplicationByNameAndSpaceReturns.result1, fake.getApplicationByNameAndSpaceReturns.result2, fake.getApplicationByNameAndSpaceReturns.result3
}

func (fake *FakeV3Actor) GetApplicationByNameAndSpaceCallCount() int {
	fake.getApplicationByNameAndSpaceMutex.RLock()
	defer fake.getApplicationByNameAndSpaceMutex.RUnlock()
	return len(fake.getApplicationByNameAndSpaceArgsForCall)
}

func (fake *FakeV3Actor) GetApplicationByNameAndSpaceArgsForCall(i int) (string, string) {
	fake.getApplicationByNameAndSpaceMutex.RLock()
	defer fake.getApplicationByNameAndSpaceMutex.RUnlock()
	return fake.getApplicationByNameAndSpaceArgsForCall[i].appName, fake.getApplicationByNameAndSpaceArgsForCall[i].spaceGUID
}

func (fake *FakeV3Actor) GetApplicationByNameAndSpaceReturns(result1 v3action.Application, result2 v3action.Warnings, result3 error) {
	fake.GetApplicationByNameAndSpaceStub = nil
	fake.getApplicationByNameAndSpaceReturns = struct {
		result1 v3action.Application
		result2 v3action.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeV3Actor) GetApplicationByNameAndSpaceReturnsOnCall(i int, result1 v3action.Application, result2 v3action.Warnings, result3 error) {
	fake.GetApplicationByNameAndSpaceStub = nil
	if fake.getApplicationByNameAndSpaceReturnsOnCall == nil {
		fake.getApplicationByNameAndSpaceReturnsOnCall = make(map[int]struct {
			result1 v3action.Application
			result2 v3action.Warnings
			result3 error
		})
	}
	fake.getApplicationByNameAndSpaceReturnsOnCall[i] = struct {
		result1 v3action.Application
		result2 v3action.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeV3Actor) PollPackage(pkg v3action.Package) (v3action.Package, v3action.Warnings, error) {
	fake.pollPackageMutex.Lock()
	ret, specificReturn := fake.pollPackageReturnsOnCall[len(fake.pollPackageArgsForCall)]
	fake.pollPackageArgsForCall = append(fake.pollPackageArgsForCall, struct {
		pkg v3action.Package
	}{pkg})
	fake.recordInvocation("PollPackage", []interface{}{pkg})
	fake.pollPackageMutex.Unlock()
	if fake.PollPackageStub != nil {
		return fake.PollPackageStub(pkg)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	return fake.pollPackageReturns.result1, fake.pollPackageReturns.result2, fake.pollPackageReturns.result3
}

func (fake *FakeV3Actor) PollPackageCallCount() int {
	fake.pollPackageMutex.RLock()
	defer fake.pollPackageMutex.RUnlock()
	return len(fake.pollPackageArgsForCall)
}

func (fake *FakeV3Actor) PollPackageArgsForCall(i int) v3action.Package {
	fake.pollPackageMutex.RLock()
	defer fake.pollPackageMutex.RUnlock()
	return fake.pollPackageArgsForCall[i].pkg
}

func (fake *FakeV3Actor) PollPackageReturns(result1 v3action.Package, result2 v3action.Warnings, result3 error) {
	fake.PollPackageStub = nil
	fake.pollPackageReturns = struct {
		result1 v3action.Package
		result2 v3action.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeV3Actor) PollPackageReturnsOnCall(i int, result1 v3action.Package, result2 v3action.Warnings, result3 error) {
	fake.PollPackageStub = nil
	if fake.pollPackageReturnsOnCall == nil {
		fake.pollPackageReturnsOnCall = make(map[int]struct {
			result1 v3action.Package
			result2 v3action.Warnings
			result3 error
		})
	}
	fake.pollPackageReturnsOnCall[i] = struct {
		result1 v3action.Package
		result2 v3action.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeV3Actor) SetApplicationDroplet(appGUID string, dropletGUID string) (v3action.Warnings, error) {
	fake.setApplicationDropletMutex.Lock()
	ret, specificReturn := fake.setApplicationDropletReturnsOnCall[len(fake.setApplicationDropletArgsForCall)]
	fake.setApplicationDropletArgsForCall = append(fake.setApplicationDropletArgsForCall, struct {
		appGUID     string
		dropletGUID string
	}{appGUID, dropletGUID})
	fake.recordInvocation("SetApplicationDroplet", []interface{}{appGUID, dropletGUID})
	fake.setApplicationDropletMutex.Unlock()
	if fake.SetApplicationDropletStub != nil {
		return fake.SetApplicationDropletStub(appGUID, dropletGUID)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.setApplicationDropletReturns.result1, fake.setApplicationDropletReturns.result2
}

func (fake *FakeV3Actor) SetApplicationDropletCallCount() int {
	fake.setApplicationDropletMutex.RLock()
	defer fake.setApplicationDropletMutex.RUnlock()
	return len(fake.setApplicationDropletArgsForCall)
}

func (fake *FakeV3Actor) SetApplicationDropletArgsForCall(i int) (string, string) {
	fake.setApplicationDropletMutex.RLock()
	defer fake.setApplicationDropletMutex.RUnlock()
	return fake.setApplicationDropletArgsForCall[i].appGUID, fake.setApplicationDropletArgsForCall[i].dropletGUID
}

func (fake *FakeV3Actor) SetApplicationDropletReturns(result1 v3action.Warnings, result2 error) {
	fake.SetApplicationDropletStub = nil
	fake.setApplicationDropletReturns = struct {
		result1 v3action.Warnings
		result2 error
	}{result1, result2}
}

func (fake *FakeV3Actor) SetApplicationDropletReturnsOnCall(i int, result1 v3action.Warnings, result2 error) {
	fake.SetApplicationDropletStub = nil
	if fake.setApplicationDropletReturnsOnCall == nil {
		fake.setApplicationDropletReturnsOnCall = make(map[int]struct {
			result1 v3action.Warnings
			result2 error
		})
	}
	fake.setApplicationDropletReturnsOnCall[i] = struct {
		result1 v3action.Warnings
		result2 error
	}{result1, result2}
}

func (fake *FakeV3Actor) StageApplicationPackage(pkgGUID string) (v3action.Build, v3action.Warnings, error) {
	fake.stageApplicationPackageMutex.Lock()
	ret, specificReturn := fake.stageApplicationPackageReturnsOnCall[len(fake.stageApplicationPackageArgsForCall)]
	fake.stageApplicationPackageArgsForCall = append(fake.stageApplicationPackageArgsForCall, struct {
		pkgGUID string
	}{pkgGUID})
	fake.recordInvocation("StageApplicationPackage", []interface{}{pkgGUID})
	fake.stageApplicationPackageMutex.Unlock()
	if fake.StageApplicationPackageStub != nil {
		return fake.StageApplicationPackageStub(pkgGUID)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	return fake.stageApplicationPackageReturns.result1, fake.stageApplicationPackageReturns.result2, fake.stageApplicationPackageReturns.result3
}

func (fake *FakeV3Actor) StageApplicationPackageCallCount() int {
	fake.stageApplicationPackageMutex.RLock()
	defer fake.stageApplicationPackageMutex.RUnlock()
	return len(fake.stageApplicationPackageArgsForCall)
}

func (fake *FakeV3Actor) StageApplicationPackageArgsForCall(i int) string {
	fake.stageApplicationPackageMutex.RLock()
	defer fake.stageApplicationPackageMutex.RUnlock()
	return fake.stageApplicationPackageArgsForCall[i].pkgGUID
}

func (fake *FakeV3Actor) StageApplicationPackageReturns(result1 v3action.Build, result2 v3action.Warnings, result3 error) {
	fake.StageApplicationPackageStub = nil
	fake.stageApplicationPackageReturns = struct {
		result1 v3action.Build
		result2 v3action.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeV3Actor) StageApplicationPackageReturnsOnCall(i int, result1 v3action.Build, result2 v3action.Warnings, result3 error) {
	fake.StageApplicationPackageStub = nil
	if fake.stageApplicationPackageReturnsOnCall == nil {
		fake.stageApplicationPackageReturnsOnCall = make(map[int]struct {
			result1 v3action.Build
			result2 v3action.Warnings
			result3 error
		})
	}
	fake.stageApplicationPackageReturnsOnCall[i] = struct {
		result1 v3action.Build
		result2 v3action.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeV3Actor) PollBuild(buildGUID string, appName string) (v3action.Droplet, v3action.Warnings, error) {
	fake.pollBuildMutex.Lock()
	ret, specificReturn := fake.pollBuildReturnsOnCall[len(fake.pollBuildArgsForCall)]
	fake.pollBuildArgsForCall = append(fake.pollBuildArgsForCall, struct {
		buildGUID string
		appName   string
	}{buildGUID, appName})
	fake.recordInvocation("PollBuild", []interface{}{buildGUID, appName})
	fake.pollBuildMutex.Unlock()
	if fake.PollBuildStub != nil {
		return fake.PollBuildStub(buildGUID, appName)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	return fake.pollBuildReturns.result1, fake.pollBuildReturns.result2, fake.pollBuildReturns.result3
}

func (fake *FakeV3Actor) PollBuildCallCount() int {
	fake.pollBuildMutex.RLock()
	defer fake.pollBuildMutex.RUnlock()
	return len(fake.pollBuildArgsForCall)
}

func (fake *FakeV3Actor) PollBuildArgsForCall(i int) (string, string) {
	fake.pollBuildMutex.RLock()
	defer fake.pollBuildMutex.RUnlock()
	return fake.pollBuildArgsForCall[i].buildGUID, fake.pollBuildArgsForCall[i].appName
}

func (fake *FakeV3Actor) PollBuildReturns(result1 v3action.Droplet, result2 v3action.Warnings, result3 error) {
	fake.PollBuildStub = nil
	fake.pollBuildReturns = struct {
		result1 v3action.Droplet
		result2 v3action.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeV3Actor) PollBuildReturnsOnCall(i int, result1 v3action.Droplet, result2 v3action.Warnings, result3 error) {
	fake.PollBuildStub = nil
	if fake.pollBuildReturnsOnCall == nil {
		fake.pollBuildReturnsOnCall = make(map[int]struct {
			result1 v3action.Droplet
			result2 v3action.Warnings
			result3 error
		})
	}
	fake.pollBuildReturnsOnCall[i] = struct {
		result1 v3action.Droplet
		result2 v3action.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeV3Actor) UpdateApplication(arg1 v3action.Application) (v3action.Application, v3action.Warnings, error) {
	fake.updateApplicationMutex.Lock()
	ret, specificReturn := fake.updateApplicationReturnsOnCall[len(fake.updateApplicationArgsForCall)]
	fake.updateApplicationArgsForCall = append(fake.updateApplicationArgsForCall, struct {
		arg1 v3action.Application
	}{arg1})
	fake.recordInvocation("UpdateApplication", []interface{}{arg1})
	fake.updateApplicationMutex.Unlock()
	if fake.UpdateApplicationStub != nil {
		return fake.UpdateApplicationStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	return fake.updateApplicationReturns.result1, fake.updateApplicationReturns.result2, fake.updateApplicationReturns.result3
}

func (fake *FakeV3Actor) UpdateApplicationCallCount() int {
	fake.updateApplicationMutex.RLock()
	defer fake.updateApplicationMutex.RUnlock()
	return len(fake.updateApplicationArgsForCall)
}

func (fake *FakeV3Actor) UpdateApplicationArgsForCall(i int) v3action.Application {
	fake.updateApplicationMutex.RLock()
	defer fake.updateApplicationMutex.RUnlock()
	return fake.updateApplicationArgsForCall[i].arg1
}

func (fake *FakeV3Actor) UpdateApplicationReturns(result1 v3action.Application, result2 v3action.Warnings, result3 error) {
	fake.UpdateApplicationStub = nil
	fake.updateApplicationReturns = struct {
		result1 v3action.Application
		result2 v3action.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeV3Actor) UpdateApplicationReturnsOnCall(i int, result1 v3action.Application, result2 v3action.Warnings, result3 error) {
	fake.UpdateApplicationStub = nil
	if fake.updateApplicationReturnsOnCall == nil {
		fake.updateApplicationReturnsOnCall = make(map[int]struct {
			result1 v3action.Application
			result2 v3action.Warnings
			result3 error
		})
	}
	fake.updateApplicationReturnsOnCall[i] = struct {
		result1 v3action.Application
		result2 v3action.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeV3Actor) UploadBitsPackage(arg1 v3action.Package, arg2 []sharedaction.Resource, arg3 io.Reader, arg4 int64) (v3action.Package, v3action.Warnings, error) {
	var arg2Copy []sharedaction.Resource
	if arg2 != nil {
		arg2Copy = make([]sharedaction.Resource, len(arg2))
		copy(arg2Copy, arg2)
	}
	fake.uploadBitsPackageMutex.Lock()
	ret, specificReturn := fake.uploadBitsPackageReturnsOnCall[len(fake.uploadBitsPackageArgsForCall)]
	fake.uploadBitsPackageArgsForCall = append(fake.uploadBitsPackageArgsForCall, struct {
		arg1 v3action.Package
		arg2 []sharedaction.Resource
		arg3 io.Reader
		arg4 int64
	}{arg1, arg2Copy, arg3, arg4})
	fake.recordInvocation("UploadBitsPackage", []interface{}{arg1, arg2Copy, arg3, arg4})
	fake.uploadBitsPackageMutex.Unlock()
	if fake.UploadBitsPackageStub != nil {
		return fake.UploadBitsPackageStub(arg1, arg2, arg3, arg4)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	return fake.uploadBitsPackageReturns.result1, fake.uploadBitsPackageReturns.result2, fake.uploadBitsPackageReturns.result3
}

func (fake *FakeV3Actor) UploadBitsPackageCallCount() int {
	fake.uploadBitsPackageMutex.RLock()
	defer fake.uploadBitsPackageMutex.RUnlock()
	return len(fake.uploadBitsPackageArgsForCall)
}

func (fake *FakeV3Actor) UploadBitsPackageArgsForCall(i int) (v3action.Package, []sharedaction.Resource, io.Reader, int64) {
	fake.uploadBitsPackageMutex.RLock()
	defer fake.uploadBitsPackageMutex.RUnlock()
	return fake.uploadBitsPackageArgsForCall[i].arg1, fake.uploadBitsPackageArgsForCall[i].arg2, fake.uploadBitsPackageArgsForCall[i].arg3, fake.uploadBitsPackageArgsForCall[i].arg4
}

func (fake *FakeV3Actor) UploadBitsPackageReturns(result1 v3action.Package, result2 v3action.Warnings, result3 error) {
	fake.UploadBitsPackageStub = nil
	fake.uploadBitsPackageReturns = struct {
		result1 v3action.Package
		result2 v3action.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeV3Actor) UploadBitsPackageReturnsOnCall(i int, result1 v3action.Package, result2 v3action.Warnings, result3 error) {
	fake.UploadBitsPackageStub = nil
	if fake.uploadBitsPackageReturnsOnCall == nil {
		fake.uploadBitsPackageReturnsOnCall = make(map[int]struct {
			result1 v3action.Package
			result2 v3action.Warnings
			result3 error
		})
	}
	fake.uploadBitsPackageReturnsOnCall[i] = struct {
		result1 v3action.Package
		result2 v3action.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeV3Actor) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.cloudControllerAPIVersionMutex.RLock()
	defer fake.cloudControllerAPIVersionMutex.RUnlock()
	fake.createApplicationDeploymentMutex.RLock()
	defer fake.createApplicationDeploymentMutex.RUnlock()
	fake.createApplicationInSpaceMutex.RLock()
	defer fake.createApplicationInSpaceMutex.RUnlock()
	fake.createBitsPackageByApplicationMutex.RLock()
	defer fake.createBitsPackageByApplicationMutex.RUnlock()
	fake.getApplicationByNameAndSpaceMutex.RLock()
	defer fake.getApplicationByNameAndSpaceMutex.RUnlock()
	fake.pollPackageMutex.RLock()
	defer fake.pollPackageMutex.RUnlock()
	fake.setApplicationDropletMutex.RLock()
	defer fake.setApplicationDropletMutex.RUnlock()
	fake.stageApplicationPackageMutex.RLock()
	defer fake.stageApplicationPackageMutex.RUnlock()
	fake.pollBuildMutex.RLock()
	defer fake.pollBuildMutex.RUnlock()
	fake.updateApplicationMutex.RLock()
	defer fake.updateApplicationMutex.RUnlock()
	fake.uploadBitsPackageMutex.RLock()
	defer fake.uploadBitsPackageMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeV3Actor) recordInvocation(key string, args []interface{}) {
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

var _ pushaction.V3Actor = new(FakeV3Actor)
