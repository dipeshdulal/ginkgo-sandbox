package gomega_sandbox_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	gomega_sandbox "ginkgo-tests"
)

var _ = Describe("Books", func() {
	Describe("GetBookName", func() {
		It("returns the name of the book", func() {
			book := gomega_sandbox.Books{
				Name: "The Little Prince",
			}
			Expect(book.GetBookName()).To(Equal("The Little Prince"))
		})
	})
})
