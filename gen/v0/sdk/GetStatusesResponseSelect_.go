// generated by metactl sdk gen 
package sdk

const (
	GetStatusesResponseSelectName = "GetStatusesResponseSelect"
)

type GetStatusesResponseSelect struct {
    All *bool `json:"all,omitempty" yaml:"all,omitempty"`
    Meta *CollectionMetaSelect `json:"meta,omitempty" yaml:"meta,omitempty"`
    Statuses *StatusSelect `json:"statuses,omitempty" yaml:"statuses,omitempty"`
}