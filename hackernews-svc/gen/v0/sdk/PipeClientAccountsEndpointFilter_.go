// generated by metactl sdk gen 
package sdk

const (
	PipeClientAccountsEndpointFilterName = "PipeClientAccountsEndpointFilter"
)

type PipeClientAccountsEndpointFilter struct {
    And []PipeClientAccountsEndpointFilter `json:"and,omitempty" yaml:"and,omitempty"`
    Not []PipeClientAccountsEndpointFilter `json:"not,omitempty" yaml:"not,omitempty"`
    Or []PipeClientAccountsEndpointFilter `json:"or,omitempty" yaml:"or,omitempty"`
    Set *bool `json:"set,omitempty" yaml:"set,omitempty"`
}