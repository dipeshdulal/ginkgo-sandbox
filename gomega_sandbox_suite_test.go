package gomega_sandbox_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestGomegaSandbox(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "GomegaSandbox Suite")
}
