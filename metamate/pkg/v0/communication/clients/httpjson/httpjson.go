package httpjson

import (
	"context"
	"github.com/metamatex/metamate/generic/pkg/v0/generic"
	"github.com/metamatex/metamate/generic/pkg/v0/transport/httpjson"
	"github.com/metamatex/metamate/metamate/pkg/v0/types"
	"net/http"
)

func GetRequestHandler(f generic.Factory, client *http.Client) types.RequestHandler {
	return func(ctx context.Context, addr string, gReq generic.Generic) (generic.Generic, error) {
		return httpjson.Send(f, client, addr, "", gReq)
	}
}
