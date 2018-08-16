package v3action

import "code.cloudfoundry.org/cli/api/cloudcontroller/ccv3"

func (actor Actor) CreateApplicationDeployment(appGUID string) (Warnings, error) {
	dep := ccv3.Deployment{}
	warnings, err := actor.CloudControllerClient.CreateApplicationDeployment(dep)

	return Warnings(warnings), err
}
