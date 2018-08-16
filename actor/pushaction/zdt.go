package pushaction

import "code.cloudfoundry.org/cli/actor/v3action"

func (actor Actor) DeployApplication(appGUID string) (Warnings, error) {
	warnings, err := actor.V3Actor.CreateApplicationDeployment(appGUID)

	return Warnings(warnings), err
}

func (actor Actor) ZdtPollStart(appGUID string, warningsChannel chan<- v3action.Warnings) error {
	return nil
}
