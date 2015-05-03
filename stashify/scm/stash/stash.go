package stash

type StashAPIErrors struct {
	Errors []map[string]string `json:"errors"`
}
