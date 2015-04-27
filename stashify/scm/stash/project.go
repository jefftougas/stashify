package stash

import (
	log "github.com/Sirupsen/logrus"
)

type StashProject struct {
	Name       string `mapstructure:"project"`
	Url        string
	Repository string
	Users      []string
}

func (p StashProject) Create(name string) error {
	log.Debug("Creating project %s", name)
	return nil
}
