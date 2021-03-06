// generated by metactl sdk gen 
package transport

import (
    "context"

    "github.com/metamatex/metamate/gen/v0/sdk"
)

type Client interface {
	GetBlueWhatevers(context.Context, sdk.GetBlueWhateversRequest) (*sdk.GetBlueWhateversResponse, error)
	GetPostFeeds(context.Context, sdk.GetPostFeedsRequest) (*sdk.GetPostFeedsResponse, error)
	GetPosts(context.Context, sdk.GetPostsRequest) (*sdk.GetPostsResponse, error)
	GetServices(context.Context, sdk.GetServicesRequest) (*sdk.GetServicesResponse, error)
	GetSocialAccounts(context.Context, sdk.GetSocialAccountsRequest) (*sdk.GetSocialAccountsResponse, error)
	GetWhatevers(context.Context, sdk.GetWhateversRequest) (*sdk.GetWhateversResponse, error)
	LookupService(context.Context, sdk.LookupServiceRequest) (*sdk.LookupServiceResponse, error)
	PipeWhatevers(context.Context, sdk.PipeWhateversRequest) (*sdk.PipeWhateversResponse, error)
}