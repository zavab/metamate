// generated by metactl sdk gen 
package sdk

const (
	PostBlueWhateversResponseSelectName = "PostBlueWhateversResponseSelect"
)

type PostBlueWhateversResponseSelect struct {
    All *bool `json:"all,omitempty" yaml:"all,omitempty"`
    BlueWhatevers *BlueWhateverSelect `json:"blueWhatevers,omitempty" yaml:"blueWhatevers,omitempty"`
    Meta *ResponseMetaSelect `json:"meta,omitempty" yaml:"meta,omitempty"`
}