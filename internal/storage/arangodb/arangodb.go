package arangodb

type (
	ArangoDBMeta struct {
		ID          string `json:"_id,omitempty"`
		Key         string `json:"_key,omitempty"`
		Revision    string `json:"_rev,omitempty"`
		OldRevision string `json:"_old_rev,omitempty"`
	}
)
