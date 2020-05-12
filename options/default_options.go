package options

import (
	"fmt"
)

type DefaultOptionsSource struct {
	defaultOptions interface{}
}

func NewDefaultOptionsSource(defaultOptions interface{}) DefaultOptionsSource {
	return DefaultOptionsSource{
		defaultOptions: defaultOptions,
	}
}

func (s *DefaultOptionsSource) GetOptions() (interface{}, error) {
	if s.defaultOptions == nil {
		return nil, fmt.Errorf("Default options object not set.")
	}

	return s.defaultOptions, nil
}
