package e2e

import (
	"time"

	. "github.com/onsi/ginkgo"
	"github.com/solo-io/gloo/pkg/api/types/v1"
	"github.com/solo-io/gloo/pkg/plugins/rest"
	. "github.com/solo-io/gloo/test/helpers"
)

var _ = Describe("Swagger Function Discovery", func() {
	petstoreUpstreamName := namespace + "-petstore-8080"
	Context("creating a vhost with a route to a function generated by"+
		" gloo-function-discovery", func() {
		getPath := "/api/pets"
		functionPath := "/function"
		functionPathWithParams := "/function-with-params"
		vhostName := "one-route"
		BeforeEach(func() {
			_, err := gloo.V1().VirtualHosts().Create(&v1.VirtualHost{
				Name: vhostName,
				Routes: []*v1.Route{
					{
						Matcher: &v1.Route_RequestMatcher{
							RequestMatcher: &v1.RequestMatcher{
								Path: &v1.RequestMatcher_PathPrefix{
									PathPrefix: getPath,
								},
								Verbs: []string{"GET"},
							},
						},
						SingleDestination: &v1.Destination{
							DestinationType: &v1.Destination_Upstream{
								Upstream: &v1.UpstreamDestination{
									Name: petstoreUpstreamName,
								},
							},
						},
					},
					{
						Matcher: &v1.Route_RequestMatcher{
							RequestMatcher: &v1.RequestMatcher{
								Path: &v1.RequestMatcher_PathExact{
									PathExact: functionPath,
								},
							},
						},
						SingleDestination: &v1.Destination{
							DestinationType: &v1.Destination_Function{
								Function: &v1.FunctionDestination{
									FunctionName: "addPet",
									UpstreamName: petstoreUpstreamName,
								},
							},
						},
					},
					{
						Matcher: &v1.Route_RequestMatcher{
							RequestMatcher: &v1.RequestMatcher{
								Path: &v1.RequestMatcher_PathExact{
									PathExact: functionPathWithParams,
								},
							},
						},
						SingleDestination: &v1.Destination{
							DestinationType: &v1.Destination_Function{
								Function: &v1.FunctionDestination{
									FunctionName: "addPet",
									UpstreamName: petstoreUpstreamName,
								},
							},
						},
						Extensions: rest.EncodeRouteExtension(rest.RouteExtension{
							Parameters: &rest.Parameters{
								Headers: map[string]string{
									"x-id":   "{id}",
									"x-name": "{name}",
									"x-tag":  "{tag}",
								},
							},
						}),
					},
				},
			})
			Must(err)
		})
		AfterEach(func() {
			gloo.V1().Upstreams().Delete(petstoreUpstreamName)
			gloo.V1().VirtualHosts().Delete(vhostName)
		})
		It("should route to the petstore function", func() {
			curlEventuallyShouldRespond(curlOpts{
				path: functionPath,
				body: `{"id": 3, "tag": "donkey", "name": "videogamedunkey"}`,
			}, "< HTTP/1.1 200", time.Minute*5)
			curlEventuallyShouldRespond(curlOpts{
				path: getPath + "/3",
			}, "< HTTP/1.1 200", time.Minute*5)
			curlEventuallyShouldRespond(curlOpts{
				path: getPath + "/3",
			}, `{"id":3,"name":"videogamedunkey"}`, time.Minute*5)
		})
		It("using params: should route to the petstore function", func() {
			curlEventuallyShouldRespond(curlOpts{
				path: functionPathWithParams,
				headers: map[string]string{
					"x-id":   "4",
					"x-name": "spatula",
					"x-tag":  "dolphin",
				},
			}, "< HTTP/1.1 200", time.Minute*5)
			curlEventuallyShouldRespond(curlOpts{
				path: getPath + "/4",
			}, "< HTTP/1.1 200", time.Minute*5)
			curlEventuallyShouldRespond(curlOpts{
				path: getPath + "/4",
			}, `{"id":4,"name":"spatula"}`, time.Minute*5)
		})
	})
})
