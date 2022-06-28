package investapi

import (
	"context"
	"fmt"
	investapi "homework-2/pkg/api"
)

func GetSharePrice(share *investapi.Share) (lastPrice []*investapi.LastPrice, err error) {
	conn, err := initConnection()
	if err != nil {
		return
	}
	getLastPriceResponse, err := investapi.NewMarketDataServiceClient(conn).GetLastPrices(context.Background(), &investapi.GetLastPricesRequest{
		Figi: []string{share.Figi},
	})
	if err != nil {
		return []*investapi.LastPrice{}, fmt.Errorf("NewMarketDataServiceClient Error: %s", err)
	}
	if getLastPriceResponse == nil {
		return []*investapi.LastPrice{}, fmt.Errorf("instrument %s has no found", share.Figi)
	}
	err = conn.Close()
	if err != nil {
		return []*investapi.LastPrice{}, fmt.Errorf("conn.Close() Error: %s", err)
	}
	return getLastPriceResponse.LastPrices, nil
}
