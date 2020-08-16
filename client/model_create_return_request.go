/*
 * v4
 *
 * The bol.com API for Retailers.
 *
 * API version: 4.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package client
// CreateReturnRequest struct for CreateReturnRequest
type CreateReturnRequest struct {
	// The id for the order item (1 order can have multiple order items).
	OrderItemId string `json:"orderItemId"`
	// The quantity of items returned.
	QuantityReturned int32 `json:"quantityReturned"`
	// The handling result requested by the user.
	HandlingResult string `json:"handlingResult"`
}
