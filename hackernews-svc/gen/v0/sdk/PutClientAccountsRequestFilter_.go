// generated by metactl sdk gen 
package sdk

const (
	PutClientAccountsRequestFilterName = "PutClientAccountsRequestFilter"
)

type PutClientAccountsRequestFilter struct {
    And []PutClientAccountsRequestFilter `json:"and,omitempty" yaml:"and,omitempty"`
    ClientAccounts *ClientAccountListFilter `json:"clientAccounts,omitempty" yaml:"clientAccounts,omitempty"`
    Meta *RequestMetaFilter `json:"meta,omitempty" yaml:"meta,omitempty"`
    Mode *PutModeFilter `json:"mode,omitempty" yaml:"mode,omitempty"`
    Not []PutClientAccountsRequestFilter `json:"not,omitempty" yaml:"not,omitempty"`
    Or []PutClientAccountsRequestFilter `json:"or,omitempty" yaml:"or,omitempty"`
    Set *bool `json:"set,omitempty" yaml:"set,omitempty"`
}