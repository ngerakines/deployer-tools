package main

import (
	"bytes"
	"errors"
	"fmt"
	"text/template"
)

type Command interface {
	Run() error
}

type UpdateServiceCommand struct {
	Dry          bool
	Mapping      Mapping
	Event        *Event
	HostTemplate string
}

type TemplateParams struct {
	Cluster string
}

func BuildHostTemplate(hostTemplate, cluster string) (string, error) {
	t, err := template.New("host").Parse(hostTemplate)
	if err != nil {
		return "", err
	}

	params := TemplateParams{cluster}

	var doc bytes.Buffer
	t.Execute(&doc, params)
	s := doc.String()
	return s, nil
}

func (c *UpdateServiceCommand) Run() error {
	// mapping Mapping, event *Event, hostTemplate string
	if c.Mapping == nil {
		return errors.New("Invalid mapping")
	}
	if c.Event == nil {
		return errors.New("Invalid event")
	}

	var prefix string
	if c.Dry == true {
		prefix = "# "
	}

	projectMapping, hasProjectMapping := c.Mapping[c.Event.Project]
	if !hasProjectMapping {
		return nil
	}

	branchMapping, hasBranchMapping := projectMapping[c.Event.Branch]
	if !hasBranchMapping {
		return nil
	}

	for cluster, services := range branchMapping {
		host, err := BuildHostTemplate(c.HostTemplate, cluster)
		if err != nil {
			return err
		}
		for _, service := range services {
			fmt.Printf("%sdocker -H %s service update --image %s %s\n", prefix, host, c.Event.Container, service)
		}
	}

	return nil
}
