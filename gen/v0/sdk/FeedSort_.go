// generated by metactl sdk gen 
package sdk

const (
	FeedSortName = "FeedSort"
)

type FeedSort struct {
    Id *ServiceIdSort `json:"id,omitempty" yaml:"id,omitempty"`
    Info *InfoSort `json:"info,omitempty" yaml:"info,omitempty"`
    Meta *TypeMetaSort `json:"meta,omitempty" yaml:"meta,omitempty"`
}