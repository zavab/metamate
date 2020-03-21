// generated by metactl sdk gen 
package sdk

const (
	PipeDeleteFeedsContextFilterName = "PipeDeleteFeedsContextFilter"
)

type PipeDeleteFeedsContextFilter struct {
    And []PipeDeleteFeedsContextFilter `json:"and,omitempty" yaml:"and,omitempty"`
    ClientRequest *DeleteFeedsRequestFilter `json:"clientRequest,omitempty" yaml:"clientRequest,omitempty"`
    ClientResponse *DeleteFeedsResponseFilter `json:"clientResponse,omitempty" yaml:"clientResponse,omitempty"`
    Not []PipeDeleteFeedsContextFilter `json:"not,omitempty" yaml:"not,omitempty"`
    Or []PipeDeleteFeedsContextFilter `json:"or,omitempty" yaml:"or,omitempty"`
    ServiceRequest *DeleteFeedsRequestFilter `json:"serviceRequest,omitempty" yaml:"serviceRequest,omitempty"`
    ServiceResponse *DeleteFeedsResponseFilter `json:"serviceResponse,omitempty" yaml:"serviceResponse,omitempty"`
    Set *bool `json:"set,omitempty" yaml:"set,omitempty"`
}