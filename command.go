package main

import (
	"fmt"
)

type Command interface {
	Run() error
}

type CommandOption func(*CommandOptions)

type CommandOptions struct {
	Dry          bool
	WithAuth     bool
	Mapping      Mapping
	Event        *Event
	HostTemplate string
}

func WithDry(dry bool) CommandOption {
	return func(o *CommandOptions) {
		o.Dry = dry
	}
}

func WithMapping(m Mapping) CommandOption {
	return func(o *CommandOptions) {
		o.Mapping = m
	}
}

func WithEvent(e *Event) CommandOption {
	return func(o *CommandOptions) {
		o.Event = e
	}
}

func WithAuth(b bool) CommandOption {
	return func(o *CommandOptions) {
		o.WithAuth = b
	}
}

func WithHostTemplate(h string) CommandOption {
	return func(o *CommandOptions) {
		o.HostTemplate = h
	}
}

func NewCommand(opts ...CommandOption) Command {
	options := CommandOptions{}
	for _, o := range opts {
		o(&options)
	}
	switch options.Event.Type {
	case "service.update":
		return &UpdateServiceCommand{
			Dry:          options.Dry,
			Mapping:      options.Mapping,
			Event:        options.Event,
			HostTemplate: options.HostTemplate,
			WithAuth:     options.WithAuth,
		}
	default:
		fmt.Printf("Unknown event type '%s'.", options.Event.Type)
	}
	return nil
}
