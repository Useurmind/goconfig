package options

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type CmdTest struct {
	Name   string `opt:"name,n"`
	Number int    `opt:"number,nr"`
	Flag   bool   `opt:"flag,f"`
}

func TestCMDLineOptionsFullName(t *testing.T) {
	name := "klsfjgöfd-ghlsd"
	number := 243425
	flag := true

	args := make([]string, 6)
	args[1] = "-name"
	args[2] = name
	args[3] = "-number"
	args[4] = fmt.Sprintf("%d", number)
	args[5] = "-flag"

	source := NewCMDLineOptionsSource(args)
	options := CmdTest{}

	err := source.FillOptions(&options)
	assert.Nil(t, err)

	assert.Equal(t, name, options.Name)
	assert.Equal(t, number, options.Number)
	assert.Equal(t, flag, options.Flag)
}

func TestCMDLineOptionsAlias(t *testing.T) {
	name := "klsfjgöfd-ghlsd"
	number := 243425
	flag := true

	args := make([]string, 6)
	args[1] = "-n"
	args[2] = name
	args[3] = "-nr"
	args[4] = fmt.Sprintf("%d", number)
	args[5] = "-f"

	source := NewCMDLineOptionsSource(args)
	options := CmdTest{}

	err := source.FillOptions(&options)
	assert.Nil(t, err)

	assert.Equal(t, name, options.Name)
	assert.Equal(t, number, options.Number)
	assert.Equal(t, flag, options.Flag)
}

func TestCMDLineOptionsMixed(t *testing.T) {
	name := "klsfjgöfd-ghlsd"
	number := 243425
	flag := true

	args := make([]string, 6)
	args[1] = "-nr"
	args[2] = fmt.Sprintf("%d", number)
	args[3] = "-flag"
	args[4] = "-name"
	args[5] = name

	source := NewCMDLineOptionsSource(args)
	options := CmdTest{}

	err := source.FillOptions(&options)
	assert.Nil(t, err)

	assert.Equal(t, name, options.Name)
	assert.Equal(t, number, options.Number)
	assert.Equal(t, flag, options.Flag)
}
