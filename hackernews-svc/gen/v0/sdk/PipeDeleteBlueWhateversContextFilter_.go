// generated by metactl sdk gen 
package sdk

const (
	PipeDeleteBlueWhateversContextFilterName = "PipeDeleteBlueWhateversContextFilter"
)

type PipeDeleteBlueWhateversContextFilter struct {
    And []PipeDeleteBlueWhateversContextFilter `json:"and,omitempty" yaml:"and,omitempty"`
    ClientRequest *DeleteBlueWhateversRequestFilter `json:"clientRequest,omitempty" yaml:"clientRequest,omitempty"`
    ClientResponse *DeleteBlueWhateversResponseFilter `json:"clientResponse,omitempty" yaml:"clientResponse,omitempty"`
    Not []PipeDeleteBlueWhateversContextFilter `json:"not,omitempty" yaml:"not,omitempty"`
    Or []PipeDeleteBlueWhateversContextFilter `json:"or,omitempty" yaml:"or,omitempty"`
    ServiceRequest *DeleteBlueWhateversRequestFilter `json:"serviceRequest,omitempty" yaml:"serviceRequest,omitempty"`
    ServiceResponse *DeleteBlueWhateversResponseFilter `json:"serviceResponse,omitempty" yaml:"serviceResponse,omitempty"`
    Set *bool `json:"set,omitempty" yaml:"set,omitempty"`
}