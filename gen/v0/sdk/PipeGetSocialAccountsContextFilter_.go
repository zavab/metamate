// generated by metactl sdk gen 
package sdk

const (
	PipeGetSocialAccountsContextFilterName = "PipeGetSocialAccountsContextFilter"
)

type PipeGetSocialAccountsContextFilter struct {
    And []PipeGetSocialAccountsContextFilter `json:"and,omitempty" yaml:"and,omitempty"`
    ClientRequest *GetSocialAccountsRequestFilter `json:"clientRequest,omitempty" yaml:"clientRequest,omitempty"`
    ClientResponse *GetSocialAccountsResponseFilter `json:"clientResponse,omitempty" yaml:"clientResponse,omitempty"`
    Not []PipeGetSocialAccountsContextFilter `json:"not,omitempty" yaml:"not,omitempty"`
    Or []PipeGetSocialAccountsContextFilter `json:"or,omitempty" yaml:"or,omitempty"`
    ServiceRequest *GetSocialAccountsRequestFilter `json:"serviceRequest,omitempty" yaml:"serviceRequest,omitempty"`
    ServiceResponse *GetSocialAccountsResponseFilter `json:"serviceResponse,omitempty" yaml:"serviceResponse,omitempty"`
    Set *bool `json:"set,omitempty" yaml:"set,omitempty"`
}