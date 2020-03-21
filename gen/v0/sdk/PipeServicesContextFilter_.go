// generated by metactl sdk gen 
package sdk

const (
	PipeServicesContextFilterName = "PipeServicesContextFilter"
)

type PipeServicesContextFilter struct {
    And []PipeServicesContextFilter `json:"and,omitempty" yaml:"and,omitempty"`
    Delete *PipeDeleteServicesContextFilter `json:"delete,omitempty" yaml:"delete,omitempty"`
    Get *PipeGetServicesContextFilter `json:"get,omitempty" yaml:"get,omitempty"`
    Not []PipeServicesContextFilter `json:"not,omitempty" yaml:"not,omitempty"`
    Or []PipeServicesContextFilter `json:"or,omitempty" yaml:"or,omitempty"`
    Post *PipePostServicesContextFilter `json:"post,omitempty" yaml:"post,omitempty"`
    Put *PipePutServicesContextFilter `json:"put,omitempty" yaml:"put,omitempty"`
    Set *bool `json:"set,omitempty" yaml:"set,omitempty"`
}