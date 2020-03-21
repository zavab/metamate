// generated by metactl sdk gen 
package sdk

const (
	StatusRelationshipsName = "StatusRelationships"
)

type StatusRelationships struct {
    AuthoredByMe *bool `json:"authoredByMe,omitempty" yaml:"authoredByMe,omitempty"`
    FavoredByMe *bool `json:"favoredByMe,omitempty" yaml:"favoredByMe,omitempty"`
    MentionsMe *bool `json:"mentionsMe,omitempty" yaml:"mentionsMe,omitempty"`
    MutedByMe *bool `json:"mutedByMe,omitempty" yaml:"mutedByMe,omitempty"`
    NotReadByMe *bool `json:"notReadByMe,omitempty" yaml:"notReadByMe,omitempty"`
    ReadByMe *bool `json:"readByMe,omitempty" yaml:"readByMe,omitempty"`
    RebloggedByMe *bool `json:"rebloggedByMe,omitempty" yaml:"rebloggedByMe,omitempty"`
    RepliesToMe *bool `json:"repliesToMe,omitempty" yaml:"repliesToMe,omitempty"`
}