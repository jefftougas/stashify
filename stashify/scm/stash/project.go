package stash

import (
	"fmt"
	log "github.com/Sirupsen/logrus"
	"github.com/franela/goreq"
)

type StashProject struct {
	Name       string `mapstructure:"project"`
	Uri        string `mapstructure:"url"`
	Repository string
	Members    []string
	Username   string
	Password   string
	Key        string
}

type NewStashProject struct {
	Name string `json:"name"`
	Key  string `json:"key"`
}

func (p StashProject) Create(name string, key string) error {
	log.Debug("Attemping to make project called ", name)

	if name == "" {
		name = p.Name
	}

	if key == "" {
		key = p.Key
	}

	project := NewStashProject{Name: name, Key: key}

	req := p.Request("/rest/api/1.0/projects")
	req.Method = "POST"
	req.Body = project

	res, err := req.Do()

	if err != nil {
		log.Error(err.Error())
		return err
	}

	if res.Response.StatusCode != 201 {
		errors := StashAPIErrors{}
		res.Body.FromJsonTo(&errors)

		for _, v := range errors.Errors {
			log.Error(v["message"])
		}

		return nil
	}

	log.Info("Project ", name, " created successfully")
	return nil
}

/* Make HTTP requests on behalf of this project */
func (p StashProject) Request(resource string) *goreq.Request {
	url := fmt.Sprintf("%s%s", p.Uri, resource)
	log.Debug("Initialising request for ", url)

	req := &goreq.Request{Uri: url, BasicAuthUsername: p.Username, BasicAuthPassword: p.Password}
	req.ContentType = "application/json"
	req.Accept = "application/json"
	return req
}
