# \PricingApi

All URIs are relative to *https://api.bol.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetRetailPrice**](PricingApi.md#GetRetailPrice) | **Get** /retailer/pricing/retail-prices/{ean} | Retrieve retail price information for an EAN.



## GetRetailPrice

> RetailPriceResponse GetRetailPrice(ctx, ean)

Retrieve retail price information for an EAN.

Currently this endpoint only supports the allowable retail price and can support the following use cases:<br /><br />                     1) EANs that have been unpublished due to price related reasons can be checked against this endpoint.<br />                     2) Requesting the allowable retail price for EANs that are not yet in your assortment can help inform price setting.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ean** | **string**| The EAN number associated with this product. | 

### Return type

[**RetailPriceResponse**](RetailPriceResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.retailer.v4+json, application/vnd.retailer.v4+xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

