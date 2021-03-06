// generated by metactl sdk gen 
package sdk

const (
	EndpointsFilterName = "EndpointsFilter"
)

type EndpointsFilter struct {
    And []EndpointsFilter `json:"and,omitempty" yaml:"and,omitempty"`
    GetAttachments *GetAttachmentsEndpointFilter `json:"getAttachments,omitempty" yaml:"getAttachments,omitempty"`
    GetBlueWhatevers *GetBlueWhateversEndpointFilter `json:"getBlueWhatevers,omitempty" yaml:"getBlueWhatevers,omitempty"`
    GetPostFeeds *GetPostFeedsEndpointFilter `json:"getPostFeeds,omitempty" yaml:"getPostFeeds,omitempty"`
    GetPosts *GetPostsEndpointFilter `json:"getPosts,omitempty" yaml:"getPosts,omitempty"`
    GetServices *GetServicesEndpointFilter `json:"getServices,omitempty" yaml:"getServices,omitempty"`
    GetSocialAccounts *GetSocialAccountsEndpointFilter `json:"getSocialAccounts,omitempty" yaml:"getSocialAccounts,omitempty"`
    GetWhatevers *GetWhateversEndpointFilter `json:"getWhatevers,omitempty" yaml:"getWhatevers,omitempty"`
    LookupService *LookupServiceEndpointFilter `json:"lookupService,omitempty" yaml:"lookupService,omitempty"`
    Not []EndpointsFilter `json:"not,omitempty" yaml:"not,omitempty"`
    Or []EndpointsFilter `json:"or,omitempty" yaml:"or,omitempty"`
    PipeAttachments *PipeAttachmentsEndpointFilter `json:"pipeAttachments,omitempty" yaml:"pipeAttachments,omitempty"`
    PipeBlueWhatevers *PipeBlueWhateversEndpointFilter `json:"pipeBlueWhatevers,omitempty" yaml:"pipeBlueWhatevers,omitempty"`
    PipePostFeeds *PipePostFeedsEndpointFilter `json:"pipePostFeeds,omitempty" yaml:"pipePostFeeds,omitempty"`
    PipePosts *PipePostsEndpointFilter `json:"pipePosts,omitempty" yaml:"pipePosts,omitempty"`
    PipeServices *PipeServicesEndpointFilter `json:"pipeServices,omitempty" yaml:"pipeServices,omitempty"`
    PipeSocialAccounts *PipeSocialAccountsEndpointFilter `json:"pipeSocialAccounts,omitempty" yaml:"pipeSocialAccounts,omitempty"`
    PipeWhatevers *PipeWhateversEndpointFilter `json:"pipeWhatevers,omitempty" yaml:"pipeWhatevers,omitempty"`
    Set *bool `json:"set,omitempty" yaml:"set,omitempty"`
}