// generated by metactl sdk gen 
package sdk

const (
	PostBankAccountsRequestFilterName = "PostBankAccountsRequestFilter"
)

type PostBankAccountsRequestFilter struct {
    And []PostBankAccountsRequestFilter `json:"and,omitempty" yaml:"and,omitempty"`
    BankAccounts *BankAccountListFilter `json:"bankAccounts,omitempty" yaml:"bankAccounts,omitempty"`
    Meta *RequestMetaFilter `json:"meta,omitempty" yaml:"meta,omitempty"`
    Mode *PostModeFilter `json:"mode,omitempty" yaml:"mode,omitempty"`
    Not []PostBankAccountsRequestFilter `json:"not,omitempty" yaml:"not,omitempty"`
    Or []PostBankAccountsRequestFilter `json:"or,omitempty" yaml:"or,omitempty"`
    Set *bool `json:"set,omitempty" yaml:"set,omitempty"`
}