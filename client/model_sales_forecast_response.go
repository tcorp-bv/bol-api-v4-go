/*
 * v4
 *
 * The bol.com API for Retailers.
 *
 * API version: 4.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package client
// SalesForecastResponse struct for SalesForecastResponse
type SalesForecastResponse struct {
	// Indicator name.
	Name string `json:"name,omitempty"`
	// Interpretation of the data that applies to this measurement.
	Type string `json:"type,omitempty"`
	Total TotalResponse `json:"total"`
	Countries []CountryTotalResponse `json:"countries"`
	Periods []PeriodResponse `json:"periods"`
}
