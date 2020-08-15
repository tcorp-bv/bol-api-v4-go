# \ShippingLabelsApi

All URIs are relative to *https://api.bol.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetDeliveryOptions**](ShippingLabelsApi.md#GetDeliveryOptions) | **Post** /retailer/shipping-labels/delivery-options | Get delivery options for a shippable configuration of a number of order items within an order.
[**GetShippingLabel**](ShippingLabelsApi.md#GetShippingLabel) | **Get** /retailer/shipping-labels/{shipping-label-id} | Get a shipping label
[**PostShippingLabel**](ShippingLabelsApi.md#PostShippingLabel) | **Post** /retailer/shipping-labels | Create a shipping label



## GetDeliveryOptions

> DeliveryOptionsResponse GetDeliveryOptions(ctx, optional)

Get delivery options for a shippable configuration of a number of order items within an order.

Retrieves all available delivery options based on the supplied configuration of order items that has to be shipped.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetDeliveryOptionsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetDeliveryOptionsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of DeliveryOptionsRequest**](DeliveryOptionsRequest.md)|  | 

### Return type

[**DeliveryOptionsResponse**](DeliveryOptionsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/vnd.retailer.v4+json, application/vnd.retailer.v4+xml
- **Accept**: application/vnd.retailer.v4+json, application/vnd.retailer.v4+xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetShippingLabel

> []string GetShippingLabel(ctx, shippingLabelId)

Get a shipping label

Gets a shipping label by shipping label id.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**shippingLabelId** | **string**| The shipping label id. | 

### Return type

**[]string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.retailer.v4+pdf

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostShippingLabel

> ProcessStatus PostShippingLabel(ctx, optional)

Create a shipping label

Create a shipping label with a shipping label offer id retrieved from get delivery options.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PostShippingLabelOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a PostShippingLabelOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of ShippingLabelRequest**](ShippingLabelRequest.md)|  | 

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

