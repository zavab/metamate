// generated by metactl sdk gen 
package sdk

const (
	DeleteWhateversResponseFilterName = "DeleteWhateversResponseFilter"
)

type DeleteWhateversResponseFilter struct {
    And []DeleteWhateversResponseFilter `json:"and,omitempty" yaml:"and,omitempty"`
    Meta *ResponseMetaFilter `json:"meta,omitempty" yaml:"meta,omitempty"`
    Not []DeleteWhateversResponseFilter `json:"not,omitempty" yaml:"not,omitempty"`
    Or []DeleteWhateversResponseFilter `json:"or,omitempty" yaml:"or,omitempty"`
    Set *bool `json:"set,omitempty" yaml:"set,omitempty"`
}