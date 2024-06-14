package main

import (
	"context"
	"fmt"

	"gomodules.xyz/azure-retail-prices-sdk-for-go/sdk"

	"github.com/Azure/go-autorest/autorest/to"
)

func main() {
	ctx := context.Background()

	client, err := sdk.NewRetailPricesClient(nil)
	if err != nil {
		panic(err)
	}
	pager := client.NewListPager(&sdk.RetailPricesClientListOptions{
		APIVersion: to.StringPtr("2023-01-01-preview"), // 2023-01-01-preview
		Filter:     to.StringPtr(`armRegionName eq 'eastus' and serviceName eq 'Virtual Machines'`),
		// query = "armRegionName eq 'southcentralus' and armSkuName eq 'Standard_NP20s' and priceType eq 'Consumption' and contains(meterName, 'Spot')"
		MeterRegion:  to.StringPtr(`'primary'`),
		CurrencyCode: nil,
	})

	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			panic(err)
		}
		for _, v := range page.Items {
			fmt.Println(v.Location)
		}
	}
}
