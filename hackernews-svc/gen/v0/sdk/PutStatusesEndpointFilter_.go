// generated by metactl sdk gen 
package sdk

const (
	PutStatusesEndpointFilterName = "PutStatusesEndpointFilter"
)

type PutStatusesEndpointFilter struct {
    And []PutStatusesEndpointFilter `json:"and,omitempty" yaml:"and,omitempty"`
    Not []PutStatusesEndpointFilter `json:"not,omitempty" yaml:"not,omitempty"`
    Or []PutStatusesEndpointFilter `json:"or,omitempty" yaml:"or,omitempty"`
    Set *bool `json:"set,omitempty" yaml:"set,omitempty"`
}