// generated by metactl sdk gen 
package sdk

const (
	VerifyTokenRequestFilterName = "VerifyTokenRequestFilter"
)

type VerifyTokenRequestFilter struct {
    And []VerifyTokenRequestFilter `json:"and,omitempty" yaml:"and,omitempty"`
    Input *VerifyTokenInputFilter `json:"input,omitempty" yaml:"input,omitempty"`
    Meta *RequestMetaFilter `json:"meta,omitempty" yaml:"meta,omitempty"`
    Not []VerifyTokenRequestFilter `json:"not,omitempty" yaml:"not,omitempty"`
    Or []VerifyTokenRequestFilter `json:"or,omitempty" yaml:"or,omitempty"`
    Set *bool `json:"set,omitempty" yaml:"set,omitempty"`
}