// generated by metactl sdk gen 
package sdk

const (
	ServiceName = "Service"
)

type Service struct {
    AlternativeIds []Id `json:"alternativeIds,omitempty" yaml:"alternativeIds,omitempty"`
    CreatedAt *Timestamp `json:"createdAt,omitempty" yaml:"createdAt,omitempty"`
    Endpoints *Endpoints `json:"endpoints,omitempty" yaml:"endpoints,omitempty"`
    Id *ServiceId `json:"id,omitempty" yaml:"id,omitempty"`
    IsVirtual *bool `json:"isVirtual,omitempty" yaml:"isVirtual,omitempty"`
    Name *string `json:"name,omitempty" yaml:"name,omitempty"`
    Port *int32 `json:"port,omitempty" yaml:"port,omitempty"`
    SdkVersion *string `json:"sdkVersion,omitempty" yaml:"sdkVersion,omitempty"`
    Transport *string `json:"transport,omitempty" yaml:"transport,omitempty"`
    Url *Url `json:"url,omitempty" yaml:"url,omitempty"`
}