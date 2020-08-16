/*
 * v4
 *
 * The bol.com API for Retailers.
 *
 * API version: 4.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package client
// ShippingLabelRequest The configuration of order items to get delivery options for.
type ShippingLabelRequest struct {
	// Order items for which the delivery options are requested.
	OrderItems []OrderItem `json:"orderItems"`
	// Shipping label offer id for which you request a shipping label.
	ShippingLabelOfferId string `json:"shippingLabelOfferId"`
}