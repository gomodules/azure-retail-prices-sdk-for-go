package sdk

import (
	armpolicy "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	azruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/tracing"
)

// ClientOptions contains configuration settings for a client's pipeline.
type ClientOptions = armpolicy.ClientOptions

// Client is a HTTP client for use with ARM endpoints.  It consists of an endpoint, pipeline, and tracing provider.
type Client struct {
	ep string
	pl runtime.Pipeline
	tr tracing.Tracer
}

// NewClient creates a new Client instance with the provided values.
// This client is intended to be used with Azure Resource Manager endpoints.
//   - moduleName - the fully qualified name of the module where the client is defined; used by the telemetry policy and tracing provider.
//   - moduleVersion - the semantic version of the module; used by the telemetry policy and tracing provider.
//   - cred - the TokenCredential used to authenticate the request
//   - options - optional client configurations; pass nil to accept the default values
func NewClient(moduleName, moduleVersion string, options *ClientOptions) (*Client, error) {
	if options == nil {
		options = &ClientOptions{}
	}

	ep := "https://prices.azure.com"
	pl := azruntime.NewPipeline(moduleName, moduleVersion, runtime.PipelineOptions{}, &options.ClientOptions)
	tr := options.TracingProvider.NewTracer(moduleName, moduleVersion)
	return &Client{ep: ep, pl: pl, tr: tr}, nil
}

// Endpoint returns the service's base URL for this client.
func (c *Client) Endpoint() string {
	return c.ep
}

// Pipeline returns the pipeline for this client.
func (c *Client) Pipeline() runtime.Pipeline {
	return c.pl
}

// Tracer returns the tracer for this client.
func (c *Client) Tracer() tracing.Tracer {
	return c.tr
}
