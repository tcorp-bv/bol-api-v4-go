/*
 * v4
 *
 * The bol.com API for Retailers.
 *
 * API version: 4.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package client
// OfferInsight struct for OfferInsight
type OfferInsight struct {
	// The name of the requested offer insight.
	Name string `json:"name,omitempty"`
	// Interpretation of the data that applies to this measurement.
	Type string `json:"type,omitempty"`
	// Total number of customer visits on the product page when the offer had the buy box over the requested period (excluding the current day).
	Total float64 `json:"total,omitempty"`
	Countries []Country `json:"countries"`
	Periods []Periods `json:"periods"`
}
