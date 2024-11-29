package resource

import "github.com/myback/oca/proto-public/pbresource"

var (
	TypeV1Tombstone = &pbresource.Type{
		Group:        "internal",
		GroupVersion: "v1",
		Kind:         "Tombstone",
	}
)
