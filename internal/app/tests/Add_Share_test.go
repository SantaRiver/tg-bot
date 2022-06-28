package tests

import (
	"context"
	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/assert"
	"homework-2/internal/app"
	pb "homework-2/pkg/api"
	"testing"
)

func TestAddShare(t *testing.T) {
	mc := minimock.NewController(t)
	defer mc.Finish()

	mrepo := app.NewRepositoryMock(mc)
	mrepo.AddShareMock.Return(nil)
	service := app.New(mrepo)

	share := &pb.Share{
		Figi:          "BBG000BSJK38",
		ClassCode:     "SPBXM",
		Name:          "Amazon.com",
		CountryOfRisk: "US",
	}
	_, err := service.AddShare(context.Background(), share)
	assert.Nil(t, err)
}

func TestAddShareDuplicate(t *testing.T) {
	mc := minimock.NewController(t)
	defer mc.Finish()

	mrepo := app.NewRepositoryMock(mc)
	mrepo.AddShareMock.Return(app.AlreadyExistErr)
	service := app.New(mrepo)

	share := &pb.Share{
		Figi:          "BBG000BVPV84",
		ClassCode:     "SPBXM",
		Name:          "Amazon.com",
		CountryOfRisk: "US",
	}
	_, err := service.AddShare(context.Background(), share)
	_, err = service.AddShare(context.Background(), share)
	assert.Equal(t, err, app.AlreadyExistErr)
}
