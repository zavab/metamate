// generated by metactl sdk gen 
package sdk

const (
	PostFeedsResponseFilterName = "PostFeedsResponseFilter"
)

type PostFeedsResponseFilter struct {
    And []PostFeedsResponseFilter `json:"and,omitempty" yaml:"and,omitempty"`
    Feeds *FeedListFilter `json:"feeds,omitempty" yaml:"feeds,omitempty"`
    Meta *ResponseMetaFilter `json:"meta,omitempty" yaml:"meta,omitempty"`
    Not []PostFeedsResponseFilter `json:"not,omitempty" yaml:"not,omitempty"`
    Or []PostFeedsResponseFilter `json:"or,omitempty" yaml:"or,omitempty"`
    Set *bool `json:"set,omitempty" yaml:"set,omitempty"`
}