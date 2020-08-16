# Bol-api-v4-go
Golang API client generated from the [bol.com v4](https://developers.bol.com/beta-program-retailer-api/) api [swagger specification](https://api.bol.com/retailer/public/apispec/v4).

TCorp does not hold copyright over the API specification and [types.json](types.json) is from [bol.com](https://api.bol.com/retailer/public/apispec/v3).

## Using the API
Add `github.com/tcorp-bv/bol-api-v4-go v1.0.0` to go.mod.

```go
import (
    bolapi "github.com/tcorp-bv/bol-api-v4-go"
)

// Setup the authentication
api, err := bolapi.New(&auth.EnvironmentProvider{ClientIdKey: "CLIENT_ID", ClientSecretKey: "CLIENT_SECRET"})
if err != nil {
	//handle error
}

// Get the client (you should do this once)
client := api.GetClient()
	
// Use the API
res, err := client.Shipments.GetShipments(&shipments.GetShipmentsParams{Context:context.Background()})
if err != nil {
	// handle error
}

for _, s := range res.Payload.Shipments {
	println(s.ShipmentID)
}
```


## Regenerating the API
```shell script
make generate
```

## Notes on rate limiting
Bol.com shares your rate limits over all of your OAuth2.0 keys. The limits are extremely low. You should have a single service that consumes and caches the API for your other services.

This library has retry-logic build in. When a request fails due to rate-limiting , the api sends the request again. A request could thus take multiple minutes to complete. The default behavior is to retry 10 times.
