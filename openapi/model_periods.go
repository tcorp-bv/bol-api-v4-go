/*
 * v4
 *
 * The bol.com API for Retailers.
 *
 * API version: 4.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// Periods struct for Periods
type Periods struct {
	Period Period `json:"period"`
	// Total number of customer visits on the product page when the offer had the buy box over the requested period (excluding the current day).
	Total float64 `json:"total,omitempty"`
	Countries []Country `json:"countries,omitempty"`
}
