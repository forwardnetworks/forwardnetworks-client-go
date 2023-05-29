package forwardnetworks

type Version struct {
	Version string `json:"version"`
}

type ExternalId struct {
	ExternalId string `json:"externalId"`
}

type AwsPolicy struct {
	Version   string      `json:"Version"`
	Statement []Statement `json:"Statement"`
}

type Statement struct {
	Effect   string   `json:"Effect"`
	Action   []string `json:"Action"`
	Resource string   `json:"Resource"`
}

type Network struct {
    ID        string `json:"id"`
    ParentID  string `json:"parentId,omitempty"`
    Name      string `json:"name"`
    OrgID     string `json:"orgId"`
    Creator   string `json:"creator,omitempty"`
    CreatorID string `json:"creatorId,omitempty"`
    CreatedAt int64  `json:"createdAt"`
    Note      string `json:"note,omitempty"`
}