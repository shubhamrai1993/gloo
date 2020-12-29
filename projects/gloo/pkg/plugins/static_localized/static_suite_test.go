package static_localized_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestStatic(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Static_Localized Suite")
}
