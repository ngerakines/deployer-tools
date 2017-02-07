package main

import (
	"bytes"
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
