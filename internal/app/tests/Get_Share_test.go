package tests

import (
	"context"
	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/assert"
	"homework-2/internal/app"
	pb "homework-2/pkg/api"
	"testing"
)

func TestGetShare(t *testing.T) {
	mc := minimock.NewController(t)
	defer mc.Finish()

	share := &pb.Share{
		Figi:          "BBG000BSJK37",
		ClassCode:     "SPBXM",
		Currency:      "usd",
		Name:          "AT&T",
		CountryOfRisk: "US",
	}

	mrepo := app.NewRepositoryMock(mc)
	mrepo.GetShareMock.Return(share, nil)
	service := app.New(mrepo)

	resp, err := service.GetShare(context.Background(), &pb.Share{Figi: "BBG000BSJK37"})
	assert.Nil(t, err)
	assert.Equal(t, resp, share)
}
