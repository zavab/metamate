// generated by metactl sdk gen 
package sdk

const (
	PipeClientAccountsContextName = "PipeClientAccountsContext"
)

type PipeClientAccountsContext struct {
    Delete *PipeDeleteClientAccountsContext `json:"delete,omitempty" yaml:"delete,omitempty"`
    Get *PipeGetClientAccountsContext `json:"get,omitempty" yaml:"get,omitempty"`
    Post *PipePostClientAccountsContext `json:"post,omitempty" yaml:"post,omitempty"`
    Put *PipePutClientAccountsContext `json:"put,omitempty" yaml:"put,omitempty"`
}