// generated by metactl sdk gen 
package sdk

const (
	OffsetPageName = "OffsetPage"
)

type OffsetPage struct {
    Limit *int32 `json:"limit,omitempty" yaml:"limit,omitempty"`
    Offset *int32 `json:"offset,omitempty" yaml:"offset,omitempty"`
}