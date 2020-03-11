// generated by metactl sdk gen 
package kubernetes

import (
	"encoding/json"
	"github.com/metamatex/metamate/gen/v0/sdk"
	"github.com/metamatex/metamate/gen/v0/sdk/utils/ptr"
	"github.com/metamatex/metamate/gen/v0/sdk/transport"
	"net/http"
	"reflect"
	
)

type HttpJsonServer struct {
	opts HttpJsonServerOpts
}

type HttpJsonServerOpts struct {
	Service Service
}

func NewHttpJsonServer(opts HttpJsonServerOpts) (http.Handler) {
	return HttpJsonServer{opts: opts}
}

func (s HttpJsonServer) send(w http.ResponseWriter, rsp interface{}) (err error) {
	w.Header().Set(transport.ContentTypeHeader, transport.ContentTypeJson)
	w.Header().Set(transport.MetamateTypeHeader, reflect.TypeOf(rsp).Name())

	err = json.NewEncoder(w).Encode(rsp)
	if err != nil {
	    return
	}

	return
}

func (s HttpJsonServer) getService() (sdk.Service) {
	getServicesEndpoint := s.opts.Service.GetGetServicesEndpoint()

	return sdk.Service{
		Name: ptr.String(s.opts.Service.Name()),
		SdkVersion: ptr.String(sdk.Version),
		Endpoints: &sdk.Endpoints{
			LookupService: &sdk.LookupServiceEndpoint{},
			GetServices: &getServicesEndpoint,
		},
	}
}

func (s HttpJsonServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Header.Get(transport.MetamateTypeHeader) {
	case sdk.LookupServiceRequestName:
			var req sdk.LookupServiceRequest
			err := json.NewDecoder(r.Body).Decode(&req)
			if err != nil {
				return
			}
	
			svc := s.getService()
			rsp := sdk.LookupServiceResponse{
				Output: &sdk.LookupServiceOutput{
					Service: &svc,
				},
			}
	
			err = s.send(w, rsp)
			if err != nil {
				return
			}
    case sdk.GetServicesRequestName:
        var req sdk.GetServicesRequest
        err := json.NewDecoder(r.Body).Decode(&req)
        if err != nil {
            return
        }

        rsp := s.opts.Service.GetServices(r.Context(), req)

        err = s.send(w, rsp)
        if err != nil {
            return
        }
	}
}