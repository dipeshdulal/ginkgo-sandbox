package auth_test

import (
	"errors"

	"github.com/DATA-DOG/go-sqlmock"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"gorm.io/gorm"

	"ginkgo-tests/auth"
	"ginkgo-tests/mocks"
)

var _ = Describe("Repository", func() {
	var (
		db   *gorm.DB
		mock sqlmock.Sqlmock
		repo *auth.Repository
	)

	BeforeEach(func() {
		db, mock = mocks.GetMockDB()
		repo = auth.NewRepository(db)
	})

	AfterEach(func() {
		err := mock.ExpectationsWereMet()
		Expect(err).ShouldNot(HaveOccurred())
	})

	Context("when the database returns an error", func() {
		It("returns the error", func() {
			dbErr := errors.New("db error")
			mock.ExpectQuery("SELECT name FROM `items`").WillReturnError(dbErr)
			items, err := repo.GetItems()
			Expect(err).To(Equal(dbErr))
			Expect(items).To(HaveLen(0))
		})
	})

	Context("when the database returns items", func() {
		It("returns empty items", func() {
			mock.ExpectQuery("SELECT name FROM `items`").WillReturnRows(sqlmock.NewRows([]string{"name"}))
			items, err := repo.GetItems()
			Expect(err).ShouldNot(HaveOccurred())
			Expect(items).To(HaveLen(0))
		})

		It("returns one item", func() {
			mock.ExpectQuery("SELECT name FROM `items`").WillReturnRows(sqlmock.NewRows([]string{"name"}).AddRow("item1"))
			items, err := repo.GetItems()
			Expect(err).ShouldNot(HaveOccurred())
			Expect(items).To(HaveLen(1))
			Expect(items[0]).To(Equal("item1"))
		})
	})
})
