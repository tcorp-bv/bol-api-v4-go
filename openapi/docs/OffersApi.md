# \OffersApi

All URIs are relative to *https://api.bol.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetUnpublishedOfferReport**](OffersApi.md#GetUnpublishedOfferReport) | **Get** /retailer/offers/unpublished/{report-id} | Retrieve an unpublished offer report by report id
[**PostUnpublishedOfferReport**](OffersApi.md#PostUnpublishedOfferReport) | **Post** /retailer/offers/unpublished | Request an unpublished offer report



## GetUnpublishedOfferReport

> GetUnpublishedOfferReport(ctx, reportId)

Retrieve an unpublished offer report by report id

Retrieve an unpublished offer report containing all unpublished offers and reasons.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reportId** | **string**| Unique identifier for unpublished offer report. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.retailer.v4+csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostUnpublishedOfferReport

> ProcessStatus PostUnpublishedOfferReport(ctx, optional)

Request an unpublished offer report

Request an unpublished offer report containing all unpublished offers and reasons.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PostUnpublishedOfferReportOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a PostUnpublishedOfferReportOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of CreateUnpublishedOfferReportRequest**](CreateUnpublishedOfferReportRequest.md)|  | 

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

