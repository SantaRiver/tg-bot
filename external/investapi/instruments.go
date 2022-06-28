package investapi

import (
	"context"
	"fmt"
	"homework-2/internal/dict"
	investapi "homework-2/pkg/api"
)

func GetShare(Ticker string) (instrument investapi.Share, err error) {
	conn, err := initConnection()
	if err != nil {
		return investapi.Share{}, err
	}
	instr, err := investapi.NewInstrumentsServiceClient(conn).ShareBy(context.Background(), &investapi.InstrumentRequest{
		IdType:    investapi.InstrumentIdType_INSTRUMENT_ID_TYPE_TICKER,
		ClassCode: dict.ClassCodesSPB,
		Id:        Ticker,
	})
	if err != nil {
		return investapi.Share{}, fmt.Errorf("NewInstrumentsServiceClient Error: %s", err)
	}
	if instr == nil {
		return investapi.Share{}, fmt.Errorf("instrument %s has no found", Ticker)
	}
	err = conn.Close()
	if err != nil {
		return investapi.Share{}, fmt.Errorf("conn.Close() Error: %s", err)
	}
	return *instr.GetInstrument(), nil
}
