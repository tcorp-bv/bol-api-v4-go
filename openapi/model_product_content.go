/*
 * v4
 *
 * The bol.com API for Retailers.
 *
 * API version: 4.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// ProductContent struct for ProductContent
type ProductContent struct {
	// A user defined unique reference to identify the products in the upload.
	InternalReference string `json:"internalReference"`
	// A list of attributes.
	Attributes []Attribute `json:"attributes"`
}
