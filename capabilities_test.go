package agouti_test

import (
	. "github.com/doarvid/agouti"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Capabilities", func() {
	var capabilities Capabilities

	BeforeEach(func() {
		capabilities = NewCapabilities("firstEnabled", "secondEnabled")
	})

	It("should successfully encode all provided options into JSON", func() {
		capabilities.Browser("some-browser").Version("v100").Platform("some-os")
		capabilities.With("withEnabled").Without("withoutDisabled")
		capabilities.Set("deviceName", "some-device-name").Set("udid", "some-udid")
		capabilities.Proxy(ProxyConfig{
			ProxyType: "manual",
			HTTPProxy: "some-http-proxy",
			SSLProxy:  "some-http-proxy",
		})
		Expect(capabilities.JSON()).To(MatchJSON(`{
			"browserName": "some-browser",
			"version": "v100",
			"platform": "some-os",
			"withEnabled": true,
			"withoutDisabled": false,
			"deviceName": "some-device-name",
			"udid": "some-udid",
			"firstEnabled": true,
			"secondEnabled": true,
			"proxy": {
				"proxyType": "manual",
				"httpProxy": "some-http-proxy",
				"sslProxy": "some-http-proxy"
			}
		}`))
	})

	Context("when the provided options cannot be converted to JSON", func() {
		It("should return an error", func() {
			capabilities["some-feature"] = func() {}
			_, err := capabilities.JSON()
			Expect(err).To(MatchError("json: unsupported type: func()"))
		})
	})
})
