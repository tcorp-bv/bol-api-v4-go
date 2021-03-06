/*
 * v4
 *
 * The bol.com API for Retailers.
 *
 * API version: 4.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package client
// OrderItemCancellation struct for OrderItemCancellation
type OrderItemCancellation struct {
	// The id for the order item (1 order can have multiple order items).
	OrderItemId string `json:"orderItemId"`
	// The code representing the reason for cancellation of this item.
	ReasonCode string `json:"reasonCode"`
}
