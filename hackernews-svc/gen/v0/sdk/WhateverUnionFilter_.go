// generated by metactl sdk gen 
package sdk

const (
	WhateverUnionFilterName = "WhateverUnionFilter"
)

type WhateverUnionFilter struct {
    And []WhateverUnionFilter `json:"and,omitempty" yaml:"and,omitempty"`
    BoolField *BoolFilter `json:"boolField,omitempty" yaml:"boolField,omitempty"`
    EnumField *EnumFilter `json:"enumField,omitempty" yaml:"enumField,omitempty"`
    Float64Field *Float64Filter `json:"float64Field,omitempty" yaml:"float64Field,omitempty"`
    Int32Field *Int32Filter `json:"int32Field,omitempty" yaml:"int32Field,omitempty"`
    Kind *EnumFilter `json:"kind,omitempty" yaml:"kind,omitempty"`
    Not []WhateverUnionFilter `json:"not,omitempty" yaml:"not,omitempty"`
    Or []WhateverUnionFilter `json:"or,omitempty" yaml:"or,omitempty"`
    Set *bool `json:"set,omitempty" yaml:"set,omitempty"`
    StringField *StringFilter `json:"stringField,omitempty" yaml:"stringField,omitempty"`
}