// generated by metactl sdk gen 
package transport

import (
    "context"

    "github.com/metamatex/metamate/gen/v0/sdk"
)

type Client interface {
	AuthenticateClientAccount(context.Context, sdk.AuthenticateClientAccountRequest) (*sdk.AuthenticateClientAccountResponse, error)
	DeleteBlueWhatevers(context.Context, sdk.DeleteBlueWhateversRequest) (*sdk.DeleteBlueWhateversResponse, error)
	DeleteStatuses(context.Context, sdk.DeleteStatusesRequest) (*sdk.DeleteStatusesResponse, error)
	DeleteWhatevers(context.Context, sdk.DeleteWhateversRequest) (*sdk.DeleteWhateversResponse, error)
	GetBlueWhatevers(context.Context, sdk.GetBlueWhateversRequest) (*sdk.GetBlueWhateversResponse, error)
	GetClientAccounts(context.Context, sdk.GetClientAccountsRequest) (*sdk.GetClientAccountsResponse, error)
	GetFeeds(context.Context, sdk.GetFeedsRequest) (*sdk.GetFeedsResponse, error)
	GetServiceAccounts(context.Context, sdk.GetServiceAccountsRequest) (*sdk.GetServiceAccountsResponse, error)
	GetServices(context.Context, sdk.GetServicesRequest) (*sdk.GetServicesResponse, error)
	GetSocialAccounts(context.Context, sdk.GetSocialAccountsRequest) (*sdk.GetSocialAccountsResponse, error)
	GetStatuses(context.Context, sdk.GetStatusesRequest) (*sdk.GetStatusesResponse, error)
	GetWhatevers(context.Context, sdk.GetWhateversRequest) (*sdk.GetWhateversResponse, error)
	LookupService(context.Context, sdk.LookupServiceRequest) (*sdk.LookupServiceResponse, error)
	PipeClientAccounts(context.Context, sdk.PipeClientAccountsRequest) (*sdk.PipeClientAccountsResponse, error)
	PipeWhatevers(context.Context, sdk.PipeWhateversRequest) (*sdk.PipeWhateversResponse, error)
	PostBlueWhatevers(context.Context, sdk.PostBlueWhateversRequest) (*sdk.PostBlueWhateversResponse, error)
	PostClientAccounts(context.Context, sdk.PostClientAccountsRequest) (*sdk.PostClientAccountsResponse, error)
	PostServiceAccounts(context.Context, sdk.PostServiceAccountsRequest) (*sdk.PostServiceAccountsResponse, error)
	PostSocialAccounts(context.Context, sdk.PostSocialAccountsRequest) (*sdk.PostSocialAccountsResponse, error)
	PostStatuses(context.Context, sdk.PostStatusesRequest) (*sdk.PostStatusesResponse, error)
	PostWhatevers(context.Context, sdk.PostWhateversRequest) (*sdk.PostWhateversResponse, error)
	PutBlueWhatevers(context.Context, sdk.PutBlueWhateversRequest) (*sdk.PutBlueWhateversResponse, error)
	PutServiceAccounts(context.Context, sdk.PutServiceAccountsRequest) (*sdk.PutServiceAccountsResponse, error)
	PutSocialAccounts(context.Context, sdk.PutSocialAccountsRequest) (*sdk.PutSocialAccountsResponse, error)
	PutStatuses(context.Context, sdk.PutStatusesRequest) (*sdk.PutStatusesResponse, error)
	PutWhatevers(context.Context, sdk.PutWhateversRequest) (*sdk.PutWhateversResponse, error)
	VerifyToken(context.Context, sdk.VerifyTokenRequest) (*sdk.VerifyTokenResponse, error)
}