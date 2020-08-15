/*
 * v4
 *
 * The bol.com API for Retailers.
 *
 * API version: 4.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// ContainerForTheOrderItemsThatHaveToBeCancelled struct for ContainerForTheOrderItemsThatHaveToBeCancelled
type ContainerForTheOrderItemsThatHaveToBeCancelled struct {
	// List of order items to cancel. Order item id's must belong to the same order.
	OrderItems []OrderItemCancellation `json:"orderItems"`
}
