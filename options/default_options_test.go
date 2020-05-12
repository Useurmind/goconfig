package options

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

type dos_test struct {
	name string
}

func TestDefaultOptions(t*  testing.T) {
	name := "klsfjgöfd-ghlsd"
	source := NewDefaultOptionsSource(dos_test{
		name: name,
	})

	options, err := source.GetOptions()
	assert.Nil(t, err)

	dosOptions := options.(dos_test)
	assert.Equal(t, name, dosOptions.name)
}