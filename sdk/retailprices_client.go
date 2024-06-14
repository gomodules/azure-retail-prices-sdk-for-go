package sdk

import (
	"context"
	"net/http"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
)

// https://learn.microsoft.com/en-us/rest/api/cost-management/retail-prices/azure-retail-prices

const (
	moduleName    = "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute"
	moduleVersion = "v5.7.0"
)

// RetailPricesClient contains the methods for the RetailPrices group.
// Don't use this type directly, use NewRetailPricesClient() instead.
type RetailPricesClient struct {
	internal *Client
}

// NewRetailPricesClient creates a new instance of RetailPricesClient with the specified values.
//   - subscriptionID - Subscription credentials which uniquely identify Microsoft Azure subscription. The subscription ID forms
//     part of the URI for every service call.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewRetailPricesClient(options *arm.ClientOptions) (*RetailPricesClient, error) {
	cl, err := NewClient(moduleName, moduleVersion, options)
	if err != nil {
		return nil, err
	}
	client := &RetailPricesClient{
		internal: cl,
	}
	return client, nil
}

// NewListPager - Gets the list of Microsoft.Compute SKUs available for your Subscription.
//
// Generated from API version 2021-07-01
//   - options - RetailPricesClientListOptions contains the optional parameters for the RetailPricesClient.NewListPager method.
func (client *RetailPricesClient) NewListPager(options *RetailPricesClientListOptions) *runtime.Pager[RetailPricesClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[RetailPricesClientListResponse]{
		More: func(page RetailPricesClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *RetailPricesClientListResponse) (RetailPricesClientListResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "RetailPricesClient.NewListPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listCreateRequest(ctx, options)
			}, nil)
			if err != nil {
				return RetailPricesClientListResponse{}, err
			}
			return client.listHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listCreateRequest creates the List request.
func (client *RetailPricesClient) listCreateRequest(ctx context.Context, options *RetailPricesClientListOptions) (*policy.Request, error) {
	urlPath := "/api/retail/prices"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.APIVersion != nil {
		reqQP.Set("api-version", *options.APIVersion)
	}
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	if options != nil && options.MeterRegion != nil {
		reqQP.Set("meterRegion", *options.MeterRegion)
	}
	if options != nil && options.CurrencyCode != nil {
		reqQP.Set("currencyCode", *options.CurrencyCode)
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *RetailPricesClient) listHandleResponse(resp *http.Response) (RetailPricesClientListResponse, error) {
	result := RetailPricesClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.RetailPricesResult); err != nil {
		return RetailPricesClientListResponse{}, err
	}
	return result, nil
}
