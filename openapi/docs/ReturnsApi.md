# \ReturnsApi

All URIs are relative to *https://api.bol.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateReturn**](ReturnsApi.md#CreateReturn) | **Post** /retailer/returns | Create return
[**GetReturn**](ReturnsApi.md#GetReturn) | **Get** /retailer/returns/{return-id} | Get a return by return id
[**GetReturns**](ReturnsApi.md#GetReturns) | **Get** /retailer/returns | Get returns
[**HandleReturn**](ReturnsApi.md#HandleReturn) | **Put** /retailer/returns/{rma-id} | Handle a return



## CreateReturn

> ProcessStatus CreateReturn(ctx, optional)

Create return

Create a return, and automatically handle it with the provided handling result. When successfully created, the resulting return id is provided in the process status.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CreateReturnOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateReturnOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of CreateReturnRequest**](CreateReturnRequest.md)|  | 

### Return type

[**ProcessStatus**](ProcessStatus.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.retailer.v4+json, application/vnd.retailer.v4+xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetReturn

> Return GetReturn(ctx, returnId)

Get a return by return id

Retrieve a return based on the return id.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**returnId** | **int64**| The return id that identifies this return. | 

### Return type

[**Return**](Return.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.retailer.v4+json, application/vnd.retailer.v4+xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetReturns

> ReturnsResponse GetReturns(ctx, optional)

Get returns

Get a paginated list of multi-item returns, which are sorted by date in descending order.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetReturnsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetReturnsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **optional.Int32**| The page to get with a page size of 50. | [default to 1]
 **handled** | **optional.Bool**| The status of the returns you wish to see, shows either handled or unhandled returns. | 
 **fulfilmentMethod** | **optional.String**| The fulfilment method. Fulfilled by the retailer (FBR) or fulfilled by bol.com (FBB). | [default to FBR]

### Return type

[**ReturnsResponse**](ReturnsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.retailer.v4+json, application/vnd.retailer.v4+xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## HandleReturn

> ProcessStatus HandleReturn(ctx, rmaId, optional)

Handle a return

Allows the user to handle a return. This can be to either handle an open return, or change the handlingResult of an already handled return. The latter is only possible in limited scenarios, please check the returns documentation for a complete list.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**rmaId** | **int64**| The RMA (Return Merchandise Authorization) id that identifies this particular return. | 
 **optional** | ***HandleReturnOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a HandleReturnOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of ReturnRequest**](ReturnRequest.md)| The handling result requested by the retailer. | 

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

