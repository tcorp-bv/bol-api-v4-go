# \OrdersApi

All URIs are relative to *https://api.bol.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CancelOrderItem**](OrdersApi.md#CancelOrderItem) | **Put** /retailer/orders/cancellation | Cancel an order item by an order item id
[**ShipOrderItem**](OrdersApi.md#ShipOrderItem) | **Put** /retailer/orders/shipment | Ship order item



## CancelOrderItem

> ProcessStatus CancelOrderItem(ctx, optional)

Cancel an order item by an order item id

This endpoint can be used to either confirm a cancellation request by the customer or to cancel an order item you yourself are unable to fulfil.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CancelOrderItemOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CancelOrderItemOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of ContainerForTheOrderItemsThatHaveToBeCancelled**](ContainerForTheOrderItemsThatHaveToBeCancelled.md)|  | 

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


## ShipOrderItem

> ProcessStatus ShipOrderItem(ctx, optional)

Ship order item

Ship a single order item within a customer order by providing shipping information. In case you purchased a shipping label you can add the shippingLabelId to this message; in that case the transport element must be left empty. If you ship the item(s) using your own transporter method you must omit the shippingLabelId entirely.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ShipOrderItemOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ShipOrderItemOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of ShipmentRequest**](ShipmentRequest.md)|  | 

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

