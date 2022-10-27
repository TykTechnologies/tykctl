package cmd

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewCmdExample(t *testing.T) {
	example := "hello./aha testing example usage"
	cmd := NewCmd("cloud").WithExample(example).NoArgs(nil)
	assert.Equal(t, example, cmd.Example, "test example")
	assert.Equal(t, "cloud", cmd.Use, "should return cloud")
}

func TestNewCmdWithMultipleExample(t *testing.T) {
	expected := "testing double\n this is second linen\ncomment!"
	cmd := NewCmd("").WithExample("testing double\n this is second linen").WithExample("comment!").NoArgs(nil)
	assert.Equal(t, expected, cmd.Example)
}

func TestNewCmdNoArgs(t *testing.T) {
	cmd := NewCmd("hello").NoArgs(nil)
	err := cmd.Args(cmd, []string{})
	assert.Nil(t, err, "expected nil error")
	err = cmd.Args(cmd, []string{"test me"})
	assert.Error(t, err, "expected an error")

}

func TestNewCmdHidden(t *testing.T) {
	cmd := NewCmd("hidden").NoArgs(nil)
	assert.Equal(t, false, cmd.Hidden)
	cmd = NewCmd("hidden").Hidden().NoArgs(nil)
	assert.Equal(t, true, cmd.Hidden)
	assert.Equal(t, "hidden", cmd.Use)
}

func TestNewCmdDescription(t *testing.T) {
	description := "team are here test description"
	cmd := NewCmd("team").WithDescription(description).NoArgs(nil)
	assert.Equal(t, description, cmd.Short)
	assert.Equal(t, "team", cmd.Use)

}

func TestNewCmdLongDescription(t *testing.T) {
	longDescription := "testing long description"
	cmd := NewCmd("").WithLongDescription(longDescription).NoArgs(nil)
	assert.Equal(t, longDescription, cmd.Long)
	assert.Equal(t, "", cmd.Use)

}
