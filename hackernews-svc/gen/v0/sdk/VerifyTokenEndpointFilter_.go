// generated by metactl sdk gen 
package sdk

const (
	VerifyTokenEndpointFilterName = "VerifyTokenEndpointFilter"
)

type VerifyTokenEndpointFilter struct {
    And []VerifyTokenEndpointFilter `json:"and,omitempty" yaml:"and,omitempty"`
    Not []VerifyTokenEndpointFilter `json:"not,omitempty" yaml:"not,omitempty"`
    Or []VerifyTokenEndpointFilter `json:"or,omitempty" yaml:"or,omitempty"`
    Set *bool `json:"set,omitempty" yaml:"set,omitempty"`
}