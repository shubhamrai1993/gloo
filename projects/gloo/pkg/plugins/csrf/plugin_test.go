package csrf_test

import (
	envoy_config_core "github.com/envoyproxy/go-control-plane/envoy/config/core/v3"
	envoy_config_route "github.com/envoyproxy/go-control-plane/envoy/config/route/v3"
	envoycsrf "github.com/envoyproxy/go-control-plane/envoy/extensions/filters/http/csrf/v3"
	envoyhcm "github.com/envoyproxy/go-control-plane/envoy/extensions/filters/network/http_connection_manager/v3"
	envoy_type_matcher "github.com/envoyproxy/go-control-plane/envoy/type/matcher/v3"
	envoytype "github.com/envoyproxy/go-control-plane/envoy/type/v3"

	gloo_config_core "github.com/solo-io/gloo/projects/gloo/pkg/api/external/envoy/config/core/v3"
	gloocsrf "github.com/solo-io/gloo/projects/gloo/pkg/api/external/envoy/extensions/filters/http/csrf/v3"
	gloo_type_matcher "github.com/solo-io/gloo/projects/gloo/pkg/api/external/envoy/type/matcher/v3"
	glootype "github.com/solo-io/gloo/projects/gloo/pkg/api/external/envoy/type/v3"
	v1 "github.com/solo-io/gloo/projects/gloo/pkg/api/v1"
	"github.com/solo-io/gloo/projects/gloo/pkg/plugins"
	. "github.com/solo-io/gloo/projects/gloo/pkg/plugins/csrf"
	"github.com/solo-io/gloo/projects/gloo/pkg/utils"

	"github.com/golang/protobuf/ptypes"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/solo-io/solo-kit/test/matchers"
)

var _ = Describe("plugin", func() {

	var (
		glooRfp *gloo_config_core.RuntimeFractionalPercent
		envoyRfp *envoy_config_core.RuntimeFractionalPercent
		envoyZeroRfp *envoy_config_core.RuntimeFractionalPercent
		glooAdditionalOrigins []*gloo_type_matcher.StringMatcher
		envoyAdditionalOrigins []*envoy_type_matcher.StringMatcher
	)

	BeforeEach(func() {
		glooRfp = &gloo_config_core.RuntimeFractionalPercent{
			DefaultValue: &glootype.FractionalPercent{
				Numerator:   uint32(100),
				Denominator: glootype.FractionalPercent_HUNDRED,
			},
			RuntimeKey: "csrf.runtime_key",
		}

		envoyRfp = &envoy_config_core.RuntimeFractionalPercent{
			DefaultValue: &envoytype.FractionalPercent{
				Numerator:   uint32(100),
				Denominator: envoytype.FractionalPercent_HUNDRED,
			},
			RuntimeKey: "csrf.runtime_key",
		}

		envoyZeroRfp = &envoy_config_core.RuntimeFractionalPercent{
			DefaultValue: &envoytype.FractionalPercent{
				Numerator:   uint32(0),
				Denominator: envoytype.FractionalPercent_HUNDRED,
			},
		}

		glooAdditionalOrigins = []*gloo_type_matcher.StringMatcher{{
			MatchPattern: &gloo_type_matcher.StringMatcher_Exact{
				Exact: "test",
			},
			IgnoreCase: true,
		}}

		envoyAdditionalOrigins = []*envoy_type_matcher.StringMatcher{{
			MatchPattern: &envoy_type_matcher.StringMatcher_Exact{
				Exact: "test",
			},
			IgnoreCase: true,
		}}
	})

	Context("HttpFilterPlugin", func() {

		It("copies config with filters enabled", func() {
			filters, err := NewPlugin().HttpFilters(plugins.Params{}, &v1.HttpListener{
				Options: &v1.HttpListenerOptions{
					Csrf: &gloocsrf.CsrfPolicy{
						FilterEnabled:     glooRfp,
						AdditionalOrigins: glooAdditionalOrigins,
					},
				},
			})

			Expect(err).NotTo(HaveOccurred())
			expectedStageFilter := plugins.StagedHttpFilter{
				HttpFilter: &envoyhcm.HttpFilter{
					Name: FilterName,
					ConfigType: &envoyhcm.HttpFilter_TypedConfig{
						TypedConfig: utils.MustMessageToAny(&envoycsrf.CsrfPolicy{
							FilterEnabled:     envoyRfp,
							ShadowEnabled: envoyZeroRfp,
							AdditionalOrigins: envoyAdditionalOrigins,
						}),
					},
				},
				Stage: plugins.FilterStage{
					RelativeTo: 8,
					Weight:     0,
				},
			}

			Expect(filters[0].HttpFilter).To(matchers.MatchProto(expectedStageFilter.HttpFilter))
			Expect(filters[0].Stage).To(Equal(expectedStageFilter.Stage))
		})

		It("copies config with shadow enabled", func() {
			filters, err := NewPlugin().HttpFilters(plugins.Params{}, &v1.HttpListener{
				Options: &v1.HttpListenerOptions{
					Csrf: &gloocsrf.CsrfPolicy{
						ShadowEnabled:     glooRfp,
						AdditionalOrigins: glooAdditionalOrigins,
					},
				},
			})

			Expect(err).NotTo(HaveOccurred())
			expectedStageFilter := plugins.StagedHttpFilter{
				HttpFilter: &envoyhcm.HttpFilter{
					Name: FilterName,
					ConfigType: &envoyhcm.HttpFilter_TypedConfig{
						TypedConfig: utils.MustMessageToAny(&envoycsrf.CsrfPolicy{
							FilterEnabled: envoyZeroRfp,
							ShadowEnabled:     envoyRfp,
							AdditionalOrigins: envoyAdditionalOrigins,
						}),
					},
				},
				Stage: plugins.FilterStage{
					RelativeTo: 8,
					Weight:     0,
				},
			}

			Expect(filters[0].HttpFilter).To(matchers.MatchProto(expectedStageFilter.HttpFilter))
			Expect(filters[0].Stage).To(Equal(expectedStageFilter.Stage))
		})

		It("copies config with both enabled and shadow mode fields", func() {
			filters, err := NewPlugin().HttpFilters(plugins.Params{}, &v1.HttpListener{
				Options: &v1.HttpListenerOptions{
					Csrf: &gloocsrf.CsrfPolicy{
						FilterEnabled:     glooRfp,
						ShadowEnabled:     glooRfp,
						AdditionalOrigins: glooAdditionalOrigins,
					},
				},
			})

			Expect(err).NotTo(HaveOccurred())
			Expect(filters).To(Equal([]plugins.StagedHttpFilter{{
				HttpFilter: &envoyhcm.HttpFilter{
					Name: FilterName,
					ConfigType: &envoyhcm.HttpFilter_TypedConfig{
						TypedConfig: utils.MustMessageToAny(&envoycsrf.CsrfPolicy{
							FilterEnabled:     envoyRfp,
							ShadowEnabled:     envoyRfp,
							AdditionalOrigins: envoyAdditionalOrigins,
						}),
					},
				},
				Stage: plugins.FilterStage{
					RelativeTo: 8,
					Weight:     0,
				},
			}}))
		})

	})

	Context("VirtualHostPlugin", func() {

		It("allows vhost specific csrf config", func() {
			p := NewPlugin()
			out := &envoy_config_route.VirtualHost{}
			err := p.ProcessVirtualHost(plugins.VirtualHostParams{}, &v1.VirtualHost{
				Options: &v1.VirtualHostOptions{
					Csrf: &gloocsrf.CsrfPolicy{
						FilterEnabled:     glooRfp,
						ShadowEnabled:     glooRfp,
						AdditionalOrigins: glooAdditionalOrigins,
					},
				},
			}, out)
			Expect(err).NotTo(HaveOccurred())

			var cfg envoycsrf.CsrfPolicy
			err = ptypes.UnmarshalAny(out.GetTypedPerFilterConfig()[FilterName], &cfg)

			Expect(cfg.GetAdditionalOrigins()).To(Equal(envoyAdditionalOrigins))
			Expect(cfg.GetFilterEnabled()).To(Equal(envoyRfp))
			Expect(cfg.GetShadowEnabled()).To(Equal(envoyRfp))
		})

	})

	Context("WeightedDestinationPlugin", func() {

		It("allows weighted destination specific csrf config", func() {
			p := NewPlugin()
			out := &envoy_config_route.WeightedCluster_ClusterWeight{}
			err := p.ProcessWeightedDestination(plugins.RouteParams{}, &v1.WeightedDestination{
				Options: &v1.WeightedDestinationOptions{
					Csrf: &gloocsrf.CsrfPolicy{
						FilterEnabled:     glooRfp,
						ShadowEnabled:     glooRfp,
						AdditionalOrigins: glooAdditionalOrigins,
					},
				},
			}, out)
			Expect(err).NotTo(HaveOccurred())

			var cfg envoycsrf.CsrfPolicy
			err = ptypes.UnmarshalAny(out.GetTypedPerFilterConfig()[FilterName], &cfg)

			Expect(cfg.GetAdditionalOrigins()).To(Equal(envoyAdditionalOrigins))
			Expect(cfg.GetFilterEnabled()).To(Equal(envoyRfp))
			Expect(cfg.GetShadowEnabled()).To(Equal(envoyRfp))
		})

	})

	Context("RoutePlugin", func() {

		It("allows route specific csrf config", func() {
			p := NewPlugin()
			out := &envoy_config_route.Route{}
			err := p.ProcessRoute(plugins.RouteParams{}, &v1.Route{
				Options: &v1.RouteOptions{
					Csrf: &gloocsrf.CsrfPolicy{
						FilterEnabled:     glooRfp,
						ShadowEnabled:     glooRfp,
						AdditionalOrigins: glooAdditionalOrigins,
					},
				},
			}, out)
			Expect(err).NotTo(HaveOccurred())

			var cfg envoycsrf.CsrfPolicy
			err = ptypes.UnmarshalAny(out.GetTypedPerFilterConfig()[FilterName], &cfg)

			Expect(cfg.GetAdditionalOrigins()).To(Equal(envoyAdditionalOrigins))
			Expect(cfg.GetFilterEnabled()).To(Equal(envoyRfp))
			Expect(cfg.GetShadowEnabled()).To(Equal(envoyRfp))
		})

	})

})
