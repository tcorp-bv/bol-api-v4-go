# \SubscriptionsApi

All URIs are relative to *https://api.bol.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeletePushNotificationSubscription**](SubscriptionsApi.md#DeletePushNotificationSubscription) | **Delete** /retailer/subscriptions/{subscription-id} | Delete push notification subscription
[**GetPushNotificationSubscription**](SubscriptionsApi.md#GetPushNotificationSubscription) | **Get** /retailer/subscriptions/{subscription-id} | Get push notification subscription by id
[**GetPushNotificationSubscriptions**](SubscriptionsApi.md#GetPushNotificationSubscriptions) | **Get** /retailer/subscriptions | Get push notification subscriptions
[**PostPushNotificationSubscription**](SubscriptionsApi.md#PostPushNotificationSubscription) | **Post** /retailer/subscriptions | Create push notification subscription
[**PostTestPushNotification**](SubscriptionsApi.md#PostTestPushNotification) | **Post** /retailer/subscriptions/test | Send test push notification for subscriptions
[**PutPushNotificationSubscription**](SubscriptionsApi.md#PutPushNotificationSubscription) | **Put** /retailer/subscriptions/{subscription-id} | Update push notification subscription



## DeletePushNotificationSubscription

> ProcessStatus DeletePushNotificationSubscription(ctx, subscriptionId)

Delete push notification subscription

Delete a push notification subscription with the provided id.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subscriptionId** | **int64**| A unique identifier for the subscription | 

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


## GetPushNotificationSubscription

> SubscriptionResponse GetPushNotificationSubscription(ctx, subscriptionId)

Get push notification subscription by id

Retrieve a configured and active push notification subscription with the provided id.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subscriptionId** | **int64**| A unique identifier for the subscription | 

### Return type

[**SubscriptionResponse**](SubscriptionResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.retailer.v4+json, application/vnd.retailer.v4+xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPushNotificationSubscriptions

> SubscriptionsResponse GetPushNotificationSubscriptions(ctx, )

Get push notification subscriptions

Retrieve a list of all configured and active push notification subscriptions.

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**SubscriptionsResponse**](SubscriptionsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.retailer.v4+json, application/vnd.retailer.v4+xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostPushNotificationSubscription

> ProcessStatus PostPushNotificationSubscription(ctx, optional)

Create push notification subscription

Create a push notification subscription for one (or more) of the available resources. The configured URL has to support https scheme.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PostPushNotificationSubscriptionOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a PostPushNotificationSubscriptionOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of CreateSubscriptionRequest**](CreateSubscriptionRequest.md)|  | 

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


## PostTestPushNotification

> ProcessStatus PostTestPushNotification(ctx, )

Send test push notification for subscriptions

Send a test push notification to all subscriptions for the provided event.

### Required Parameters

This endpoint does not need any parameter.

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


## PutPushNotificationSubscription

> ProcessStatus PutPushNotificationSubscription(ctx, subscriptionId, optional)

Update push notification subscription

Update an existing push notification subscription with the supplied id. The configured URL has to support https scheme.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subscriptionId** | **int64**| A unique identifier for the subscription | 
 **optional** | ***PutPushNotificationSubscriptionOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a PutPushNotificationSubscriptionOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of UpdateSubscriptionRequest**](UpdateSubscriptionRequest.md)|  | 

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

