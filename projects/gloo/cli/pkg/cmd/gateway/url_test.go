package gateway_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/solo-io/gloo-edge/projects/gloo/cli/pkg/testutils"
)

var _ = Describe("Url", func() {
	It("returns the correct url of a proxy pod", func() {

		Skip("this test is temporarily disabled as it relies on an old version of the Helm chart and will " +
			"currently fail. Re-enable it after PR https://github.com/solo-io/gloo-edge/pull/451 has been merged.")

		// install gateway first
		err := testutils.Glooctl("install gateway --release 0.6.19")
		Expect(err).NotTo(HaveOccurred())

		addr, err := testutils.GlooctlOut("proxy url")
		Expect(err).NotTo(HaveOccurred())

		Expect(addr).To(HavePrefix("http://"))
	})
})
