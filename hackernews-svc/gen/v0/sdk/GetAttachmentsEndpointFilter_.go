// generated by metactl sdk gen 
package sdk

const (
	GetAttachmentsEndpointFilterName = "GetAttachmentsEndpointFilter"
)

type GetAttachmentsEndpointFilter struct {
    And []GetAttachmentsEndpointFilter `json:"and,omitempty" yaml:"and,omitempty"`
    Not []GetAttachmentsEndpointFilter `json:"not,omitempty" yaml:"not,omitempty"`
    Or []GetAttachmentsEndpointFilter `json:"or,omitempty" yaml:"or,omitempty"`
    Set *bool `json:"set,omitempty" yaml:"set,omitempty"`
}