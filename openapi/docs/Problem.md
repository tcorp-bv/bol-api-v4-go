# Problem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Type URI for this problem. Fixed value: https://api.bol.com/problems. | [optional] 
**Title** | **string** | Title describing the nature of the problem. | [optional] 
**Status** | **int32** | HTTP status returned from the endpoint causing the problem. | [optional] 
**Detail** | **string** | Detailed error message describing in additional detail what caused the service to return this problem. | [optional] 
**Host** | **string** | Host identifier describing the server instance that reported the problem. | [optional] 
**Instance** | **string** | Full URI path of the resource that reported the problem. | [optional] 
**Violations** | [**[]Violation**](Violation.md) |  | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


