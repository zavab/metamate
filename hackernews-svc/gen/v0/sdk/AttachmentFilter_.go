// generated by metactl sdk gen 
package sdk

const (
	AttachmentFilterName = "AttachmentFilter"
)

type AttachmentFilter struct {
    AlternativeIds *IdListFilter `json:"alternativeIds,omitempty" yaml:"alternativeIds,omitempty"`
    And []AttachmentFilter `json:"and,omitempty" yaml:"and,omitempty"`
    Description *StringFilter `json:"description,omitempty" yaml:"description,omitempty"`
    Id *ServiceIdFilter `json:"id,omitempty" yaml:"id,omitempty"`
    Meta *TypeMetaFilter `json:"meta,omitempty" yaml:"meta,omitempty"`
    Not []AttachmentFilter `json:"not,omitempty" yaml:"not,omitempty"`
    Or []AttachmentFilter `json:"or,omitempty" yaml:"or,omitempty"`
    Set *bool `json:"set,omitempty" yaml:"set,omitempty"`
}