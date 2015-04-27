package stash

import (
	"testing"
)

func TestProject(t *testing.T) {
	project := StashProject{Name: "Test"}
	if project.Name != "Test" {
		t.Error("StashProject Must have name field")
	}
}
