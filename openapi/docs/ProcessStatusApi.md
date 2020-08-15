# \ProcessStatusApi

All URIs are relative to *https://api.bol.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetProcessStatus**](ProcessStatusApi.md#GetProcessStatus) | **Get** /retailer/process-status/{process-status-id} | Get the status of an asynchronous process by id
[**GetProcessStatusBulk**](ProcessStatusApi.md#GetProcessStatusBulk) | **Post** /retailer/process-status | Gets the status of multiple asynchronous processes by an array of process status id&#39;s for a retailer
[**GetProcessStatusEntityId**](ProcessStatusApi.md#GetProcessStatusEntityId) | **Get** /retailer/process-status | Gets the status of an asynchronous process by entity id and event type for a retailer



## GetProcessStatus

> ProcessStatus GetProcessStatus(ctx, processStatusId)

Get the status of an asynchronous process by id

Retrieve a specific process-status, which shows information regarding a previously executed PUT/POST/DELETE request. All PUT/POST/DELETE requests on the other endpoints will supply a process-status-id in the related response. You can use this id to retrieve a status by using the endpoint below. Please note: process status instances are only retained for a limited period of time after completion. Outside of this period, a 404 will be returned for missing process statuses. Please handle this accordingly, by stopping any active polling for these statuses.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**processStatusId** | **int64**| The id of the process status being requested. This id is supplied in every response to a PUT/POST/DELETE request on the other endpoints. | 

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


## GetProcessStatusBulk

> ProcessStatusResponse GetProcessStatusBulk(ctx, optional)

Gets the status of multiple asynchronous processes by an array of process status id's for a retailer

Retrieve a list of process statuses, which shows information regarding previously executed PUT/POST/DELETE requests. No more than 1000 process status id's can be sent in a single request.Please note: process status instances are only retained for a limited period of time after completion. Outside of this period, deleted process statuses will no longer be returned. Please handle this accordingly, by stopping any active polling for these statuses.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetProcessStatusBulkOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetProcessStatusBulkOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of BulkProcessStatusRequest**](BulkProcessStatusRequest.md)|  | 

### Return type

[**ProcessStatusResponse**](ProcessStatusResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/vnd.retailer.v4+json, application/vnd.retailer.v4+xml
- **Accept**: application/vnd.retailer.v4+json, application/vnd.retailer.v4+xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProcessStatusEntityId

> ProcessStatusResponse GetProcessStatusEntityId(ctx, entityId, eventType, optional)

Gets the status of an asynchronous process by entity id and event type for a retailer

Retrieve a list of process statuses, which shows information regarding previously executed PUT/POST/DELETE requests in descending order. You need to supply an entity id and event type. Please note: process status instances are only retained for a limited period of time after completion. Outside of this period, deleted process statuses will no longer be returned. Please handle this accordingly, by stopping any active polling for these statuses.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**entityId** | **string**| The entity id is not unique so you need to provide an event type. The entity id can either be order item id, transport id, return number or inbound reference. | 
**eventType** | **string**| The event type can only be used in combination with the entity id. | 
 **optional** | ***GetProcessStatusEntityIdOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetProcessStatusEntityIdOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **page** | **optional.Int32**| The requested page number with a pagesize of 50 | [default to 1]

### Return type

[**ProcessStatusResponse**](ProcessStatusResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.retailer.v4+json, application/vnd.retailer.v4+xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

