# ShipmentRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OrderItems** | [**[]OrderItem**](OrderItem.md) | List of order items to ship. Order item id&#39;s must belong to the same order. | 
**ShipmentReference** | **string** | Used for administration purposes. | [optional] 
**ShippingLabelId** | **string** | Specifies shipping label to be used for this shipment. Can be retrieved through the shipping label endpoint. | [optional] 
**Transport** | [**TransportInstruction**](TransportInstruction.md) |  | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


