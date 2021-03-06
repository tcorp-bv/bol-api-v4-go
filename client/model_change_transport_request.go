/*
 * v4
 *
 * The bol.com API for Retailers.
 *
 * API version: 4.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package client
// ChangeTransportRequest struct for ChangeTransportRequest
type ChangeTransportRequest struct {
	TransporterCode string `json:"transporterCode,omitempty"`
	// The track and trace code that is associated with this transport.
	TrackAndTrace string `json:"trackAndTrace,omitempty"`
}
