package auth_test

import (
	"errors"
	"ginkgo-tests/auth"
	"ginkgo-tests/mocks/ginkgo-tests/contracts"
	"net/http"

	"github.com/gin-gonic/gin"
	. "github.com/onsi/ginkgo/v2"
	"github.com/steinfletcher/apitest"
)

var _ = Describe("Controller", func() {

	var (
		t          GinkgoTInterface
		router     *gin.Engine
		service    *contracts.MockAuthService
		controller *auth.Controller
	)

	BeforeEach(func() {
		// manual dependency injections
		t = GinkgoT()
		router = auth.NewRouter()

		service = contracts.NewMockAuthService(t)
		controller = auth.NewController(service)
		auth.RegisterRoutes(router, controller)
	})

	Describe("GetItems", func() {
		It("returns the items", func() {
			service.EXPECT().GetItems().Return([]string{"item1", "item2"}, nil)
			apitest.New().
				Handler(router).
				Get("/items").
				Expect(t).
				Body(`{"message":"pong","items":["item1","item2"]}`).
				Status(http.StatusOK).
				End()
		})

		It("returns empty items", func() {
			service.EXPECT().GetItems().Return([]string{}, nil)
			apitest.New().
				Handler(router).
				Get("/items").
				Expect(t).
				Body(`{"message":"pong","items":[]}`).
				Status(http.StatusOK).
				End()
		})

		It("returns error from service", func() {
			service.EXPECT().GetItems().Return(nil, errors.New("my error"))
			apitest.New().
				Handler(router).
				Get("/items").
				Expect(t).
				Body(`{"message":"my error"}`).
				Status(http.StatusInternalServerError).
				End()
		})
	})
})
