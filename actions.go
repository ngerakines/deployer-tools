package main

import (
  "errors"
)

func ServiceUpdate(mapping *Mapping, event *Event) error {
  if mapping == nil {
    return errors.New("Invalid mapping")
  }
  if event == nil {
    return errors.New("Invalid event")
  }
  return nil
}
