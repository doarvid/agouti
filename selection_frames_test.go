package agouti_test

import (
	"errors"

	. "github.com/doarvid/agouti"
	"github.com/doarvid/agouti/api"
	. "github.com/doarvid/agouti/internal/matchers"
	"github.com/doarvid/agouti/internal/mocks"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Selection Frames", func() {
	var (
		selection         *Selection
		session           *mocks.Session
		elementRepository *mocks.ElementRepository
	)

	BeforeEach(func() {
		session = &mocks.Session{}
		elementRepository = &mocks.ElementRepository{}
		selection = NewTestSelection(session, elementRepository, "#selector")
	})

	Describe("#SwitchToFrame", func() {
		var apiElement *api.Element

		BeforeEach(func() {
			apiElement = &api.Element{}
			elementRepository.GetExactlyOneCall.ReturnElement = apiElement
		})

		It("should successfully switch to the frame indicated by the selection", func() {
			Expect(selection.SwitchToFrame()).To(Succeed())
			Expect(session.FrameCall.Frame).To(ExactlyEqual(apiElement))
		})

		Context("when there is an error retrieving exactly one element", func() {
			It("should return an error", func() {
				elementRepository.GetExactlyOneCall.Err = errors.New("some error")
				err := selection.SwitchToFrame()
				Expect(err).To(MatchError("failed to select element from selection 'CSS: #selector [single]': some error"))
			})
		})

		Context("when the session fails to switch frames", func() {
			It("should return an error", func() {
				session.FrameCall.Err = errors.New("some error")
				err := selection.SwitchToFrame()
				Expect(err).To(MatchError("failed to switch to frame referred to by selection 'CSS: #selector [single]': some error"))
			})
		})
	})
})
