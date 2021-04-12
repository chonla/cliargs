package cliargs_test

import (
	"cliargs"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCmdParseWithoutCommand(t *testing.T) {
	os.Args = []string{"cli"}

	cmd, err := cliargs.Parse()

	assert.Nil(t, err)
	assert.NotNil(t, cmd)
	assert.Equal(t, "", cmd.Command)
}

func TestCmdParseWithCommand(t *testing.T) {
	os.Args = []string{"cli", "hello"}

	cmd, err := cliargs.Parse()

	assert.Nil(t, err)
	assert.NotNil(t, cmd)
	assert.Equal(t, "hello", cmd.Command)
}

func TestCmdParseWithSubcommand(t *testing.T) {
	os.Args = []string{"cli", "hello", "world"}

	cmd, err := cliargs.Parse()

	assert.Nil(t, err)
	assert.NotNil(t, cmd)
	assert.Equal(t, "hello", cmd.Command)
	assert.Equal(t, "world", cmd.Subcommand)
}

func TestCmdParseCommandWithOneOption(t *testing.T) {
	os.Args = []string{"cli", "hello", "--name=john"}

	cmd, err := cliargs.Parse()

	assert.Nil(t, err)
	assert.NotNil(t, cmd)
	assert.Equal(t, "hello", cmd.Command)
	assert.Equal(t, "john", cmd.Options.First("name"))
}

func TestCmdParseCommandWithTwoOption(t *testing.T) {
	os.Args = []string{"cli", "hello", "--name=john", "--surname=doe"}

	cmd, err := cliargs.Parse()

	assert.Nil(t, err)
	assert.NotNil(t, cmd)
	assert.Equal(t, "hello", cmd.Command)
	assert.Equal(t, "john", cmd.Options.First("name"))
	assert.Equal(t, "doe", cmd.Options.First("surname"))
}

func TestCmdParseCommandWithRepetibleOption(t *testing.T) {
	os.Args = []string{"cli", "hello", "--name=john", "--name=jane"}

	cmd, err := cliargs.Parse()

	assert.Nil(t, err)
	assert.NotNil(t, cmd)
	assert.Equal(t, "hello", cmd.Command)
	assert.Equal(t, []string{"john", "jane"}, cmd.Options.Get("name"))
}
