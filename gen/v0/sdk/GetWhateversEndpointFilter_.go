// generated by metactl sdk gen 
package sdk

const (
	GetWhateversEndpointFilterName = "GetWhateversEndpointFilter"
)

type GetWhateversEndpointFilter struct {
    And []GetWhateversEndpointFilter `json:"and,omitempty" yaml:"and,omitempty"`
    Not []GetWhateversEndpointFilter `json:"not,omitempty" yaml:"not,omitempty"`
    Or []GetWhateversEndpointFilter `json:"or,omitempty" yaml:"or,omitempty"`
    Set *bool `json:"set,omitempty" yaml:"set,omitempty"`
}