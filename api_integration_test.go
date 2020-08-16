//+build integration

package bolapiv4

import (
	"context"
	"testing"
	"time"
)

func TestAPI(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Minute)
	defer cancel()
	api := *GetTestAPI(t, true, false)
	client := api.GetClient()
	testOfferID := "8aa42288-6a4d-395f-e053-828b620a1175"
	res, _, err := client.InsightsApi.GetSalesForecast(ctx, testOfferID, 4)
	if err != nil {
		t.Fatal(err)
	}
	if res.Total.Maximum == 0 {
		t.Log("No forecast found")
		return
	}
	t.Logf("%s forecasted to sell %f-%f in 4 weeks", testOfferID, res.Total.Minimum, res.Total.Maximum)
}
