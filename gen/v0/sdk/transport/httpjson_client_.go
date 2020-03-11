// generated by metactl sdk gen 
package transport

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/metamatex/metamate/gen/v0/sdk"
	"net/http"
	"reflect"
	
)

type HttpJsonClient struct {
	opts HttpJsonClientOpts
}

type HttpJsonClientOpts struct {
	HttpClient	*http.Client
	Token	string
	Addr	string
}

func NewHttpJsonClient(opts HttpJsonClientOpts) (Client) {
	return HttpJsonClient{opts: opts}
}

func (c HttpJsonClient) send(req interface{}, rsp interface{}) (err error) {
	b := new(bytes.Buffer)
	err = json.NewEncoder(b).Encode(req)
	if err != nil {
		return
	}

	httpReq, err := http.NewRequest(http.MethodPost, c.opts.Addr, b)
	if err != nil {
		return
	}

	httpReq.Header.Set(ContentTypeHeader, ContentTypeJson)
	httpReq.Header.Set(MetamateTypeHeader, reflect.TypeOf(req).Name())
	httpReq.Header.Set(AuthorizationHeader, "Bearer " + c.opts.Token)

	res, err := c.opts.HttpClient.Do(httpReq)
	if err != nil {
		return
	}

	err = json.NewDecoder(res.Body).Decode(rsp)
	if err != nil {
		return
	}

	return
}
func (c HttpJsonClient) AuthenticateClientAccount(ctx context.Context, req sdk.AuthenticateClientAccountRequest) (rsp *sdk.AuthenticateClientAccountResponse, err error) {
	err = c.send(req, &rsp)

	return
}
func (c HttpJsonClient) DeleteBlueWhatevers(ctx context.Context, req sdk.DeleteBlueWhateversRequest) (rsp *sdk.DeleteBlueWhateversResponse, err error) {
	err = c.send(req, &rsp)

	return
}
func (c HttpJsonClient) DeleteStatuses(ctx context.Context, req sdk.DeleteStatusesRequest) (rsp *sdk.DeleteStatusesResponse, err error) {
	err = c.send(req, &rsp)

	return
}
func (c HttpJsonClient) DeleteWhatevers(ctx context.Context, req sdk.DeleteWhateversRequest) (rsp *sdk.DeleteWhateversResponse, err error) {
	err = c.send(req, &rsp)

	return
}
func (c HttpJsonClient) GetBlueWhatevers(ctx context.Context, req sdk.GetBlueWhateversRequest) (rsp *sdk.GetBlueWhateversResponse, err error) {
	err = c.send(req, &rsp)

	return
}
func (c HttpJsonClient) GetClientAccounts(ctx context.Context, req sdk.GetClientAccountsRequest) (rsp *sdk.GetClientAccountsResponse, err error) {
	err = c.send(req, &rsp)

	return
}
func (c HttpJsonClient) GetFeeds(ctx context.Context, req sdk.GetFeedsRequest) (rsp *sdk.GetFeedsResponse, err error) {
	err = c.send(req, &rsp)

	return
}
func (c HttpJsonClient) GetPeople(ctx context.Context, req sdk.GetPeopleRequest) (rsp *sdk.GetPeopleResponse, err error) {
	err = c.send(req, &rsp)

	return
}
func (c HttpJsonClient) GetServiceAccounts(ctx context.Context, req sdk.GetServiceAccountsRequest) (rsp *sdk.GetServiceAccountsResponse, err error) {
	err = c.send(req, &rsp)

	return
}
func (c HttpJsonClient) GetServices(ctx context.Context, req sdk.GetServicesRequest) (rsp *sdk.GetServicesResponse, err error) {
	err = c.send(req, &rsp)

	return
}
func (c HttpJsonClient) GetStatuses(ctx context.Context, req sdk.GetStatusesRequest) (rsp *sdk.GetStatusesResponse, err error) {
	err = c.send(req, &rsp)

	return
}
func (c HttpJsonClient) GetWhatevers(ctx context.Context, req sdk.GetWhateversRequest) (rsp *sdk.GetWhateversResponse, err error) {
	err = c.send(req, &rsp)

	return
}
func (c HttpJsonClient) LookupService(ctx context.Context, req sdk.LookupServiceRequest) (rsp *sdk.LookupServiceResponse, err error) {
	err = c.send(req, &rsp)

	return
}
func (c HttpJsonClient) PipeClientAccounts(ctx context.Context, req sdk.PipeClientAccountsRequest) (rsp *sdk.PipeClientAccountsResponse, err error) {
	err = c.send(req, &rsp)

	return
}
func (c HttpJsonClient) PipeWhatevers(ctx context.Context, req sdk.PipeWhateversRequest) (rsp *sdk.PipeWhateversResponse, err error) {
	err = c.send(req, &rsp)

	return
}
func (c HttpJsonClient) PostBlueWhatevers(ctx context.Context, req sdk.PostBlueWhateversRequest) (rsp *sdk.PostBlueWhateversResponse, err error) {
	err = c.send(req, &rsp)

	return
}
func (c HttpJsonClient) PostClientAccounts(ctx context.Context, req sdk.PostClientAccountsRequest) (rsp *sdk.PostClientAccountsResponse, err error) {
	err = c.send(req, &rsp)

	return
}
func (c HttpJsonClient) PostPeople(ctx context.Context, req sdk.PostPeopleRequest) (rsp *sdk.PostPeopleResponse, err error) {
	err = c.send(req, &rsp)

	return
}
func (c HttpJsonClient) PostServiceAccounts(ctx context.Context, req sdk.PostServiceAccountsRequest) (rsp *sdk.PostServiceAccountsResponse, err error) {
	err = c.send(req, &rsp)

	return
}
func (c HttpJsonClient) PostStatuses(ctx context.Context, req sdk.PostStatusesRequest) (rsp *sdk.PostStatusesResponse, err error) {
	err = c.send(req, &rsp)

	return
}
func (c HttpJsonClient) PostWhatevers(ctx context.Context, req sdk.PostWhateversRequest) (rsp *sdk.PostWhateversResponse, err error) {
	err = c.send(req, &rsp)

	return
}
func (c HttpJsonClient) PutBlueWhatevers(ctx context.Context, req sdk.PutBlueWhateversRequest) (rsp *sdk.PutBlueWhateversResponse, err error) {
	err = c.send(req, &rsp)

	return
}
func (c HttpJsonClient) PutPeople(ctx context.Context, req sdk.PutPeopleRequest) (rsp *sdk.PutPeopleResponse, err error) {
	err = c.send(req, &rsp)

	return
}
func (c HttpJsonClient) PutServiceAccounts(ctx context.Context, req sdk.PutServiceAccountsRequest) (rsp *sdk.PutServiceAccountsResponse, err error) {
	err = c.send(req, &rsp)

	return
}
func (c HttpJsonClient) PutStatuses(ctx context.Context, req sdk.PutStatusesRequest) (rsp *sdk.PutStatusesResponse, err error) {
	err = c.send(req, &rsp)

	return
}
func (c HttpJsonClient) PutWhatevers(ctx context.Context, req sdk.PutWhateversRequest) (rsp *sdk.PutWhateversResponse, err error) {
	err = c.send(req, &rsp)

	return
}
func (c HttpJsonClient) VerifyToken(ctx context.Context, req sdk.VerifyTokenRequest) (rsp *sdk.VerifyTokenResponse, err error) {
	err = c.send(req, &rsp)

	return
}