// generated by metactl sdk gen 
package sdk

const (
	GetWhateversResponseFilterName = "GetWhateversResponseFilter"
)

type GetWhateversResponseFilter struct {
    And []GetWhateversResponseFilter `json:"and,omitempty" yaml:"and,omitempty"`
    Meta *CollectionMetaFilter `json:"meta,omitempty" yaml:"meta,omitempty"`
    Not []GetWhateversResponseFilter `json:"not,omitempty" yaml:"not,omitempty"`
    Or []GetWhateversResponseFilter `json:"or,omitempty" yaml:"or,omitempty"`
    Set *bool `json:"set,omitempty" yaml:"set,omitempty"`
    Whatevers *WhateverListFilter `json:"whatevers,omitempty" yaml:"whatevers,omitempty"`
}