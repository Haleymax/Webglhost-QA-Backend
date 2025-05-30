package models

type Watcher struct {
	ID       string   `json:"id" bson:"_id,omitempty"`
	Name     string   `json:"name" bson:"name"`
	Resource string   `json:"resource" bson:"resource"`
	Click    string   `json:"click" bson:"click"`
	Brand    []string `json:"brand" bson:"brand"`
	Tag      []string `json:"tag" bson:"tag"`
}

type WatcherRequest struct {
	Env     string `json:"env"`
	Runtime string `json:"runtime"`
}
