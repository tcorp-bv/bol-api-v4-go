/*
 * v4
 *
 * The bol.com API for Retailers.
 *
 * API version: 4.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package client
// ShipmentRequest struct for ShipmentRequest
type ShipmentRequest struct {
	// List of order items to ship. Order item id's must belong to the same order.
	OrderItems []OrderItem `json:"orderItems"`
	// Used for administration purposes.
	ShipmentReference string `json:"shipmentReference,omitempty"`
	// Specifies shipping label to be used for this shipment. Can be retrieved through the shipping label endpoint.
	ShippingLabelId string `json:"shippingLabelId,omitempty"`
	Transport TransportInstruction `json:"transport,omitempty"`
}