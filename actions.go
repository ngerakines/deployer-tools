package main

import (
	"errors"
	"fmt"
)

type UpdateServiceCommand struct {
	Dry          bool
	Mapping      Mapping
	Event        *Event
	HostTemplate string
	WithAuth     bool
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

	authInclude := ""
	if c.WithAuth == true {
		authInclude = "--with-registry-auth "
	}

	for cluster, services := range branchMapping {
		host, err := BuildHostTemplate(c.HostTemplate, cluster)
		if err != nil {
			return err
		}
		for _, service := range services {
			fmt.Printf("%sdocker -H %s service update %s--image %s %s\n", prefix, host, authInclude, c.Event.Container, service)
		}
	}

	return nil
}
