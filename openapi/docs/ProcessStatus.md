# ProcessStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int64** | The process status id. | [optional] 
**EntityId** | **string** | The id of the object being processed. E.g. in case of a shipment process id, you will receive the id of the order item being processed. | [optional] 
**EventType** | **string** | Name of the requested action that is being processed. | [optional] 
**Description** | **string** | Describes the action that is being processed. | [optional] 
**Status** | **string** | Status of the action being processed. | [optional] 
**ErrorMessage** | **string** | Shows error message if applicable. | [optional] 
**CreateTimestamp** | [**time.Time**](time.Time.md) | Time of creation of the response. | [optional] 
**Links** | [**[]Link**](Link.md) | Lists available actions applicable to this endpoint. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


