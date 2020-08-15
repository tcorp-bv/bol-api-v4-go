# ReturnItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RmaId** | **int64** | The RMA (Return Merchandise Authorization) id that identifies this particular return. | [optional] 
**OrderId** | **string** | The id of the customer order this return item is in. | [optional] 
**Ean** | **string** | The EAN number associated with this product. | [optional] 
**Title** | **string** | The product title. | [optional] 
**ExpectedQuantity** | **int32** | The quantity that is expected to be returned by the customer. Note: this can be greater than 1 in case the customer ordered a quantity greater than 1 of the same product in the same customer order. | [optional] 
**ReturnReason** | **string** | The reason why the customer returned this product. | [optional] 
**ReturnReasonComments** | **string** | Additional details from the customer as to why this item was returned. | [optional] 
**TrackAndTrace** | **string** | The track and trace code that is associated with this transport. | [optional] 
**TransporterName** | **string** | The name of the transporter. | [optional] 
**Handled** | **bool** | Indicates if this return item has been handled (by the retailer). | [optional] 
**ProcessingResults** | [**[]ReturnProcessingResult**](ReturnProcessingResult.md) |  | 
**CustomerDetails** | [**CustomerDetails**](CustomerDetails.md) |  | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


