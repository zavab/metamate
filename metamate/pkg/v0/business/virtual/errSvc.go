package virtual

import (
	"context"
	"github.com/metamatex/metamate/asg/pkg/v0/asg/graph"
	"github.com/metamatex/metamate/gen/v0/sdk"

	"github.com/metamatex/metamate/generic/pkg/v0/generic"
	"github.com/metamatex/metamate/generic/pkg/v0/transport/httpjson"
	"github.com/metamatex/metamate/metamate/pkg/v0/types"
	"net/http"
)

func init() {
	handler[Error] = func(f generic.Factory, rn *graph.RootNode, c *http.Client, vSvc types.VirtualSvc) (h http.Handler, t string, err error) {
		h = httpjson.NewServer(httpjson.ServerOpts{
			Root:    rn,
			Factory: f,
			Handler: func(ctx context.Context, gReq generic.Generic) (gRsp generic.Generic) {
				switch gReq.Type().Name() {
				case sdk.LookupServiceRequestName:
					return f.MustFromStruct(sdk.LookupServiceResponse{
						Output: &sdk.LookupServiceOutput{
							Service: &sdk.Service{
								Endpoints: &sdk.Endpoints{
									GetWhatevers: &sdk.GetWhateversEndpoint{
										Filter: &sdk.GetWhateversRequestFilter{
										},
									},
								},
							},
						},
					})
				case sdk.GetWhateversRequestName:
					return f.MustFromStruct(sdk.GetWhateversResponse{
						Errors: []sdk.Error{
							{
								Message: sdk.String("a"),
							},
						},
						Whatevers: []sdk.Whatever{
							{
								Id: &sdk.ServiceId{
									Value: sdk.String("a"),
								},
								StringField: sdk.String("a"),
							},
						},
					})
				}

				return
			},
			LogErr: nil,
		})

		t = sdk.ServiceTransport.HttpJson

		return
	}
}
