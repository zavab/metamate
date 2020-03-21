// generated by metactl sdk gen 
package sdk

const (
	ServiceSortName = "ServiceSort"
)

type ServiceSort struct {
    Id *ServiceIdSort `json:"id,omitempty" yaml:"id,omitempty"`
    IsVirtual *string `json:"isVirtual,omitempty" yaml:"isVirtual,omitempty"`
    Meta *TypeMetaSort `json:"meta,omitempty" yaml:"meta,omitempty"`
    Name *string `json:"name,omitempty" yaml:"name,omitempty"`
    Port *string `json:"port,omitempty" yaml:"port,omitempty"`
    SdkVersion *string `json:"sdkVersion,omitempty" yaml:"sdkVersion,omitempty"`
    Url *UrlSort `json:"url,omitempty" yaml:"url,omitempty"`
}