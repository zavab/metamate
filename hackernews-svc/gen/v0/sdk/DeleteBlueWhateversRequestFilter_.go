// generated by metactl sdk gen 
package sdk

const (
	DeleteBlueWhateversRequestFilterName = "DeleteBlueWhateversRequestFilter"
)

type DeleteBlueWhateversRequestFilter struct {
    And []DeleteBlueWhateversRequestFilter `json:"and,omitempty" yaml:"and,omitempty"`
    Ids *ServiceIdListFilter `json:"ids,omitempty" yaml:"ids,omitempty"`
    Meta *RequestMetaFilter `json:"meta,omitempty" yaml:"meta,omitempty"`
    Mode *DeleteModeFilter `json:"mode,omitempty" yaml:"mode,omitempty"`
    Not []DeleteBlueWhateversRequestFilter `json:"not,omitempty" yaml:"not,omitempty"`
    Or []DeleteBlueWhateversRequestFilter `json:"or,omitempty" yaml:"or,omitempty"`
    Set *bool `json:"set,omitempty" yaml:"set,omitempty"`
}