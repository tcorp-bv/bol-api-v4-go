# \TransportsApi

All URIs are relative to *https://api.bol.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddTransportInformationByTransportId**](TransportsApi.md#AddTransportInformationByTransportId) | **Put** /retailer/transports/{transport-id} | Add transport information by transport id



## AddTransportInformationByTransportId

> ProcessStatus AddTransportInformationByTransportId(ctx, transportId, optional)

Add transport information by transport id

Add information to an existing transport. The transport id is part of the shipment. You can retrieve the transport id through the GET shipment list request.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**transportId** | **int64**| The transport id. | 
 **optional** | ***AddTransportInformationByTransportIdOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a AddTransportInformationByTransportIdOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of ChangeTransportRequest**](ChangeTransportRequest.md)| The change transport requested by the user. | 

### Return type

[**ProcessStatus**](ProcessStatus.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/vnd.retailer.v4+json, application/vnd.retailer.v4+xml
- **Accept**: application/vnd.retailer.v4+json, application/vnd.retailer.v4+xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

