package main

import (
	"bytes"
	"errors"
	"fmt"
	"text/template"
)

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

func ServiceUpdate(mapping Mapping, event *Event, hostTemplate string) error {
	if mapping == nil {
		return errors.New("Invalid mapping")
	}
	if event == nil {
		return errors.New("Invalid event")
	}

	projectMapping, hasProjectMapping := mapping[event.Project]
	if !hasProjectMapping {
		return nil
		// return errors.New("No project mapping for event.")
	}

	branchMapping, hasBranchMapping := projectMapping[event.Branch]
	if !hasBranchMapping {
		return nil
		// return errors.New("No branch mapping for event.")
	}

	for cluster, services := range branchMapping {
		host, err := BuildHostTemplate(hostTemplate, cluster)
		if err != nil {
			return err
		}
		for _, service := range services {
			fmt.Println("docker -H", host, "service update --image", event.Container, service)
		}
	}

	return nil
}
