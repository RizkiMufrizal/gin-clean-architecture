package client

import (
	"context"
	"github.com/RizkiMufrizal/gin-clean-architecture/model"
)

type HttpBinClient interface {
	PostMethod(ctx context.Context, requestBody *model.HttpBin, response *map[string]interface{})
}
