// generated by metactl sdk gen 
package sdk

const (
	GetPostFeedsResponseName = "GetPostFeedsResponse"
)

type GetPostFeedsResponse struct {
    Meta *CollectionMeta `json:"meta,omitempty" yaml:"meta,omitempty"`
    PostFeeds []PostFeed `json:"postFeeds,omitempty" yaml:"postFeeds,omitempty"`
}