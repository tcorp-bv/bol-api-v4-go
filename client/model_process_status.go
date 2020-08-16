/*
 * v4
 *
 * The bol.com API for Retailers.
 *
 * API version: 4.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package client
import (
	"time"
)
// ProcessStatus struct for ProcessStatus
type ProcessStatus struct {
	// The process status id.
	Id int64 `json:"id,omitempty"`
	// The id of the object being processed. E.g. in case of a shipment process id, you will receive the id of the order item being processed.
	EntityId string `json:"entityId,omitempty"`
	// Name of the requested action that is being processed.
	EventType string `json:"eventType,omitempty"`
	// Describes the action that is being processed.
	Description string `json:"description,omitempty"`
	// Status of the action being processed.
	Status string `json:"status,omitempty"`
	// Shows error message if applicable.
	ErrorMessage string `json:"errorMessage,omitempty"`
	// Time of creation of the response.
	CreateTimestamp time.Time `json:"createTimestamp,omitempty"`
	// Lists available actions applicable to this endpoint.
	Links []Link `json:"links,omitempty"`
}
