// generated by metactl sdk gen 
package sdk

const (
	DeleteTransactionsResponseFilterName = "DeleteTransactionsResponseFilter"
)

type DeleteTransactionsResponseFilter struct {
    And []DeleteTransactionsResponseFilter `json:"and,omitempty" yaml:"and,omitempty"`
    Meta *ResponseMetaFilter `json:"meta,omitempty" yaml:"meta,omitempty"`
    Not []DeleteTransactionsResponseFilter `json:"not,omitempty" yaml:"not,omitempty"`
    Or []DeleteTransactionsResponseFilter `json:"or,omitempty" yaml:"or,omitempty"`
    Set *bool `json:"set,omitempty" yaml:"set,omitempty"`
}