/*
 * v4
 *
 * The bol.com API for Retailers.
 *
 * API version: 4.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// DeliveryOption A delivery option shows how and the costs of a transport for a shippable configuration
type DeliveryOption struct {
	// Unique identifier for the shipping label offer.
	ShippingLabelOfferId string `json:"shippingLabelOfferId,omitempty"`
	// A code representing the transporter which is being used for transportation.
	TransporterCode string `json:"transporterCode,omitempty"`
	// The type of the label, representing the way an item is being transported.
	LabelType string `json:"labelType,omitempty"`
	LabelPrice LabelPrice `json:"labelPrice"`
	PackageRestrictions PackageRestrictions `json:"packageRestrictions"`
}