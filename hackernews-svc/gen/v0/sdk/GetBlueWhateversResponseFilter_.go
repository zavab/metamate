// generated by metactl sdk gen 
package sdk

const (
	GetBlueWhateversResponseFilterName = "GetBlueWhateversResponseFilter"
)

type GetBlueWhateversResponseFilter struct {
    And []GetBlueWhateversResponseFilter `json:"and,omitempty" yaml:"and,omitempty"`
    BlueWhatevers *BlueWhateverListFilter `json:"blueWhatevers,omitempty" yaml:"blueWhatevers,omitempty"`
    Meta *CollectionMetaFilter `json:"meta,omitempty" yaml:"meta,omitempty"`
    Not []GetBlueWhateversResponseFilter `json:"not,omitempty" yaml:"not,omitempty"`
    Or []GetBlueWhateversResponseFilter `json:"or,omitempty" yaml:"or,omitempty"`
    Set *bool `json:"set,omitempty" yaml:"set,omitempty"`
}