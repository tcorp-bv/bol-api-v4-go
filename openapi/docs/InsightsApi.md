# \InsightsApi

All URIs are relative to *https://api.bol.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetOfferInsights**](InsightsApi.md#GetOfferInsights) | **Get** /retailer/insights/offer | Get offer insights
[**GetSalesForecast**](InsightsApi.md#GetSalesForecast) | **Get** /retailer/insights/sales-forecast | Get sales forecast



## GetOfferInsights

> OfferInsights GetOfferInsights(ctx, offerId, period, numberOfPeriods, name)

Get offer insights

Gets offer insights.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**offerId** | **string**| Unique identifier for an offer. | 
**period** | **string**| The time unit in which the offer insights are grouped. | 
**numberOfPeriods** | **int32**| The number of periods for which the offer insights are requested back in time. | 
**name** | [**[]string**](string.md)| The name of the requested offer insight. | 

### Return type

[**OfferInsights**](OfferInsights.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.retailer.v4+json, application/vnd.retailer.v4+xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSalesForecast

> SalesForecastResponse GetSalesForecast(ctx, offerId, weeksAhead)

Get sales forecast

Gets sales forecast.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**offerId** | **string**| Unique identifier for an offer. | 
**weeksAhead** | **int32**| The number of weeks into the future, starting from today. | 

### Return type

[**SalesForecastResponse**](SalesForecastResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.retailer.v4+json, application/vnd.retailer.v4+xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

