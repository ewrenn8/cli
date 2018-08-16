package v3action

func (actor Actor) CreateApplicationDeployment(appGUID string) (Warnings, error) {
	warnings, err := actor.CloudControllerClient.CreateApplicationDeployment(appGUID)

	return Warnings(warnings), err
}
