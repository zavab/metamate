// generated by metactl sdk gen 
package sdk

const (
	PipeGetClientAccountsContextName = "PipeGetClientAccountsContext"
)

type PipeGetClientAccountsContext struct {
    ClientRequest *GetClientAccountsRequest `json:"clientRequest,omitempty" yaml:"clientRequest,omitempty"`
    ClientResponse *GetClientAccountsResponse `json:"clientResponse,omitempty" yaml:"clientResponse,omitempty"`
    ServiceRequest *GetClientAccountsRequest `json:"serviceRequest,omitempty" yaml:"serviceRequest,omitempty"`
    ServiceResponse *GetClientAccountsResponse `json:"serviceResponse,omitempty" yaml:"serviceResponse,omitempty"`
}