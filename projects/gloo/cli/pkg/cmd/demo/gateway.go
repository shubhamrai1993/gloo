package demo

import (
	"os"

	"github.com/solo-io/gloo/projects/gloo/cli/pkg/cmd/options"
	"github.com/solo-io/gloo/projects/gloo/cli/pkg/common"
	"github.com/solo-io/gloo/projects/gloo/cli/pkg/constants"
	"github.com/spf13/cobra"
)

func gateway(opts *options.Options) *cobra.Command {
	cmd := &cobra.Command{
		Use:   constants.DEMO_GATEWAY_COMMAND.Use,
		Short: constants.DEMO_GATEWAY_COMMAND.Short,
		Long:  constants.DEMO_FEDERATION_COMMAND.Long,
		RunE: func(cmd *cobra.Command, args []string) error {
			runner := common.NewShellRunner(os.Stdin, os.Stdout)
			return runner.Run("bash", "-c", initGlooEdgeDemoScript, "init-demo.sh")
		},
	}
	return cmd
}


const (
	initGlooEdgeDemoScript = `
#!/bin/bash

# Create the kind cluster
if kind get clusters;
then
    echo "Found a kind cluster already"
else
    echo "Creating a kind cluster"
#    kind create cluster
fi

which glooctl
glooctl version
glooctl install gateway

kubectl -n gloo-system rollout status deployment gloo --timeout=2m || true
kubectl -n gloo-system rollout status deployment discovery --timeout=2m || true
kubectl -n gloo-system rollout status deployment gateway-proxy --timeout=2m || true
kubectl -n gloo-system rollout status deployment gateway --timeout=2m || true

kubectl apply -f https://raw.githubusercontent.com/solo-io/gloo/v1.2.9/example/petstore/petstore.yaml
glooctl add route \
  --path-exact /all-pets \
  --dest-name default-petstore-8080 \
  --prefix-rewrite /api/pets

# Instructions for Gloo Edge demo
cat << EOF
# We now have Gloo Edge set up correctly!

# To view the upstreams, run:
glooctl get upstreams

# To view the virtual services, run:
glooctl get virtualservices

# Wait for the Virtual Service to be ACCEPTED

# For this section, use two terminals, one for the port-forward command and one for the curl command.

# Curl the route to reach the petstore upstream.

curl $(glooctl proxy url)/all-pets
EOF
`
)