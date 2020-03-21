// generated by metactl sdk gen 
package sdk

const (
	ServiceIdFilterName = "ServiceIdFilter"
)

type ServiceIdFilter struct {
    And []ServiceIdFilter `json:"and,omitempty" yaml:"and,omitempty"`
    Not []ServiceIdFilter `json:"not,omitempty" yaml:"not,omitempty"`
    Or []ServiceIdFilter `json:"or,omitempty" yaml:"or,omitempty"`
    ServiceName *StringFilter `json:"serviceName,omitempty" yaml:"serviceName,omitempty"`
    Set *bool `json:"set,omitempty" yaml:"set,omitempty"`
    Value *StringFilter `json:"value,omitempty" yaml:"value,omitempty"`
}