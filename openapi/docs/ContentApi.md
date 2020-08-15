# \ContentApi

All URIs are relative to *https://api.bol.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetValidationReport**](ContentApi.md#GetValidationReport) | **Get** /retailer/content/validation-report/{uploadId} | Get validation report
[**PostProductContent**](ContentApi.md#PostProductContent) | **Post** /retailer/content/product | Post product content



## GetValidationReport

> ValidationReportResponse GetValidationReport(ctx, uploadId)

Get validation report

Retrieve a validation report of the product content upload based on the upload id.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uploadId** | **string**| The identifier of the product content upload. | 

### Return type

[**ValidationReportResponse**](ValidationReportResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.retailer.v4+json, application/vnd.retailer.v4+xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostProductContent

> ProcessStatus PostProductContent(ctx, optional)

Post product content

Create content for existing products or new products.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PostProductContentOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a PostProductContentOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of CreateProductContentRequest**](CreateProductContentRequest.md)|  | 

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

