// generated by metactl sdk gen 
package sdk

const (
	DeleteWhateversEndpointFilterName = "DeleteWhateversEndpointFilter"
)

type DeleteWhateversEndpointFilter struct {
    And []DeleteWhateversEndpointFilter `json:"and,omitempty" yaml:"and,omitempty"`
    Not []DeleteWhateversEndpointFilter `json:"not,omitempty" yaml:"not,omitempty"`
    Or []DeleteWhateversEndpointFilter `json:"or,omitempty" yaml:"or,omitempty"`
    Set *bool `json:"set,omitempty" yaml:"set,omitempty"`
}