package v2_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/onsi/gomega/gbytes"

	"code.cloudfoundry.org/cli/api/cloudcontroller/ccversion"
	"code.cloudfoundry.org/cli/command/commandfakes"
	"code.cloudfoundry.org/cli/command/translatableerror"
	. "code.cloudfoundry.org/cli/command/v2"
	"code.cloudfoundry.org/cli/command/v2/v2fakes"
	"code.cloudfoundry.org/cli/util/ui"
)

var _ = Describe("DeleteBuildpackCommand", func() {
	var (
		cmd             DeleteBuildpackCommand
		testUI          *ui.UI
		fakeConfig      *commandfakes.FakeConfig
		fakeSharedActor *commandfakes.FakeSharedActor
		fakeActor       *v2fakes.FakeDeleteBuildpackActor

		executeErr error
	)

	BeforeEach(func() {
		testUI = ui.NewTestUI(NewBuffer(), NewBuffer(), NewBuffer())
		fakeConfig = new(commandfakes.FakeConfig)
		fakeSharedActor = new(commandfakes.FakeSharedActor)
		fakeActor = new(v2fakes.FakeDeleteBuildpackActor)

		cmd = DeleteBuildpackCommand{
			UI:          testUI,
			Config:      fakeConfig,
			SharedActor: fakeSharedActor,
			Actor:       fakeActor,
		}

		cmd.RequiredArgs.Buildpack = "bp-name"
	})

	JustBeforeEach(func() {
		executeErr = cmd.Execute(nil)
	})

	When("a stack is specified", func() {
		BeforeEach(func() {
			cmd.Stack = "some-stack"
		})

		When("the api version is below minimum for stack association", func() {
			BeforeEach(func() {
				fakeActor.CloudControllerAPIVersionReturns("2.113.9")
			})

			It("returns a version error", func() {
				Expect(executeErr).To(MatchError(translatableerror.MinimumAPIVersionNotMetError{
					Command:        "Option `-s`",
					CurrentVersion: fakeActor.CloudControllerAPIVersion(),
					MinimumVersion: ccversion.MinVersionForStackFlagV2,
				}))
			})
		})

		When("the api version is at or above minimum for stack association", func() {
			BeforeEach(func() {
				fakeActor.CloudControllerAPIVersionReturns("2.114.0")
			})

			It("returns the unrefactored command error", func() {
				Expect(executeErr).To(MatchError(translatableerror.UnrefactoredCommandError{}))
			})
		})
	})
})
