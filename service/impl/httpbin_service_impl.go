package impl

import (
	"context"
	"github.com/RizkiMufrizal/gin-clean-architecture/client"
	"github.com/RizkiMufrizal/gin-clean-architecture/common"
	"github.com/RizkiMufrizal/gin-clean-architecture/model"
	"github.com/RizkiMufrizal/gin-clean-architecture/service"
)

func NewHttpBinServiceImpl(httpBinClient *client.HttpBinClient) service.HttpBinService {
	return &httpBinServiceImpl{HttpBinClient: *httpBinClient}
}

type httpBinServiceImpl struct {
	client.HttpBinClient
}

func (h *httpBinServiceImpl) PostMethod(ctx context.Context) {
	httpBin := model.HttpBin{
		Name: "rizki",
	}
	var response map[string]interface{}
	h.HttpBinClient.PostMethod(ctx, &httpBin, &response)
	common.NewLogger().Info("log response service ", response)
}
