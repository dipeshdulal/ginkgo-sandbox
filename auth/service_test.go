package auth_test

import (
	"errors"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"ginkgo-tests/auth"
	"ginkgo-tests/mocks/ginkgo-tests/contracts"
)

var _ = Describe("Service", func() {
	var (
		t       GinkgoTInterface
		service *auth.Service
		repo    *contracts.MockAuthRepository
	)

	BeforeEach(func() {
		t = GinkgoT()
		repo = contracts.NewMockAuthRepository(t)
		service = auth.NewService(repo)
	})

	Describe("GetItems", func() {
		It("returns the items", func() {
			repo.EXPECT().GetItems().Return([]string{"item1", "item2"}, nil)
			items, err := service.GetItems()
			Expect(err).To(BeNil())
			Expect(items).To(Equal([]string{"item1", "item2"}))
		})

		Context("when the repository returns an error", func() {
			It("returns the error", func() {
				repoErr := errors.New("repo error")
				repo.EXPECT().GetItems().Return(nil, repoErr)
				items, err := service.GetItems()
				Expect(err).To(Equal(repoErr))
				Expect(items).To(BeNil())
			})
		})
	})
})
