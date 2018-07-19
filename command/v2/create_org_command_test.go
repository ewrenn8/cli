package v2_test

import (
	"code.cloudfoundry.org/cli/actor/actionerror"
	"code.cloudfoundry.org/cli/command/commandfakes"
	"code.cloudfoundry.org/cli/command/flag"
	. "code.cloudfoundry.org/cli/command/v2"
	"code.cloudfoundry.org/cli/command/v2/v2fakes"
	"code.cloudfoundry.org/cli/util/ui"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/onsi/gomega/gbytes"
)

var _ = FDescribe("CreateOrgCommand", func() {
	var (
		fakeConfig      *commandfakes.FakeConfig
		fakeActor       *v2fakes.FakeCreateOrgActor
		fakeSharedActor *commandfakes.FakeSharedActor
		testUI          *ui.UI

		cmd CreateOrgCommand

		executeErr error
	)

	BeforeEach(func() {
		testUI = ui.NewTestUI(nil, NewBuffer(), NewBuffer())
		fakeConfig = new(commandfakes.FakeConfig)
		fakeActor = new(v2fakes.FakeCreateOrgActor)
		fakeSharedActor = new(commandfakes.FakeSharedActor)

		cmd = CreateOrgCommand{
			UI:           testUI,
			Config:       fakeConfig,
			Actor:        fakeActor,
			SharedActor:  fakeSharedActor,
			RequiredArgs: flag.Organization{Organization: "some-org"},
		}
	})

	JustBeforeEach(func() {
		executeErr = cmd.Execute(nil)
	})

	Context("when checking the target fails", func() {
		var binaryName string

		BeforeEach(func() {
			binaryName = "faceman"
			fakeConfig.BinaryNameReturns(binaryName)
			fakeSharedActor.CheckTargetReturns(actionerror.NotLoggedInError{BinaryName: binaryName})
		})

		It("returns an error", func() {
			Expect(executeErr).To(MatchError(actionerror.NotLoggedInError{BinaryName: "faceman"}))

			Expect(fakeSharedActor.CheckTargetCallCount()).To(Equal(1))
			checkTargetedOrg, checkTargetedSpace := fakeSharedActor.CheckTargetArgsForCall(0)
			Expect(checkTargetedOrg).To(BeFalse())
			Expect(checkTargetedSpace).To(BeFalse())
		})
	})

	Context("when checking the target succeeds", func() {
		It("creates the org", func() {
			Expect(executeErr).ToNot(HaveOccurred())
			Expect(fakeActor.CreateOrganizationCallCount()).To(Equal(1))
			Expect(fakeActor.CreateOrganizationArgsForCall(0)).To(Equal("some-org"))
		})
	})
})
