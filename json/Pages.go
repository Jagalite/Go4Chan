package json

type PagesJSON []struct {
	Page    int `json:"page"`
	Threads []struct {
		No           int `json:"no"`
		LastModified int `json:"last_modified"`
	} `json:"threads"`
}
