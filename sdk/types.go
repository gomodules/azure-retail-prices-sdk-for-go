package sdk

import "time"

// RetailPricesClientListOptions contains the optional parameters for the RetailPricesClient.NewListPager method.
type RetailPricesClientListOptions struct {
	// Filters are supported for the following fields:
	// - armRegionName
	// - Location
	// - meterId
	// - meterName
	// - productid
	// - skuId
	// - productName
	// - skuName
	// - serviceName
	// - serviceId
	// - serviceFamily
	// - priceType
	// - armSkuName
	Filter       *string
	APIVersion   *string
	MeterRegion  *string
	CurrencyCode *string
}

// RetailPricesClientListResponse contains the response from method RetailPricesClient.NewListPager.
type RetailPricesClientListResponse struct {
	// The List Retail Prices operation response.
	RetailPricesResult
}

// RetailPricesResult - The List Retail Prices operation response.
type RetailPricesResult struct {
	BillingCurrency    string        `json:"BillingCurrency"`
	CustomerEntityID   string        `json:"CustomerEntityId"`
	CustomerEntityType string        `json:"CustomerEntityType"`
	Items              []ResourceSKU `json:"Items"`
	Count              int           `json:"Count"`

	// The URI to fetch the next page of Retail Prices. Call ListNext() with this URI to fetch the next page of Retail Prices
	NextLink *string
}

type ResourceSKU struct {
	CurrencyCode         string        `json:"currencyCode"`
	TierMinimumUnits     float64       `json:"tierMinimumUnits"`
	RetailPrice          float64       `json:"retailPrice"`
	UnitPrice            float64       `json:"unitPrice"`
	ArmRegionName        string        `json:"armRegionName"`
	Location             string        `json:"location"`
	EffectiveStartDate   time.Time     `json:"effectiveStartDate"`
	MeterID              string        `json:"meterId"`
	MeterName            string        `json:"meterName"`
	ProductID            string        `json:"productId"`
	SkuID                string        `json:"skuId"`
	ProductName          string        `json:"productName"`
	SkuName              string        `json:"skuName"`
	ServiceName          string        `json:"serviceName"`
	ServiceID            string        `json:"serviceId"`
	ServiceFamily        string        `json:"serviceFamily"`
	UnitOfMeasure        string        `json:"unitOfMeasure"`
	Type                 string        `json:"type"`
	IsPrimaryMeterRegion bool          `json:"isPrimaryMeterRegion"`
	ArmSkuName           string        `json:"armSkuName"`
	ReservationTerm      string        `json:"reservationTerm,omitempty"`
	SavingsPlan          []SavingsPlan `json:"savingsPlan,omitempty"`
	EffectiveEndDate     time.Time     `json:"effectiveEndDate,omitempty"`
}

type SavingsPlan struct {
	UnitPrice   float64 `json:"unitPrice"`
	RetailPrice float64 `json:"retailPrice"`
	Term        string  `json:"term"`
}
