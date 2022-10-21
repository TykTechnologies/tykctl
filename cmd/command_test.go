package cmd

import (
	"github.com/TykTechnologies/tykctl/testutil"
	"testing"
)

func TestNewCmdExample(t *testing.T) {
	example := "hello./aha testing example usage"
	cmd := NewCmd("cloud").WithExample(example).NoArgs(nil)
	testutil.Equal(t, example, cmd.Example)
	testutil.Equal(t, "cloud", cmd.Use)
}
func TestNewCmdWithMultipleExample(t *testing.T) {
	expected := "testing double\n this is second linen\ncomment!"

	cmd := NewCmd("").WithExample("testing double\n this is second linen").WithExample("comment!").NoArgs(nil)
	if expected != cmd.Example {
		t.Errorf("Expected \n %s  got \n %s ", expected, cmd.Example)
	}
	testutil.Equal(t, "", cmd.Use)
}

func TestNewCmdNoArgs(t *testing.T) {
	cmd := NewCmd("hello").NoArgs(nil)
	err := cmd.Args(cmd, []string{})
	if err != nil {
		t.Errorf("Expected nil error got %s", err)
	}
	err = cmd.Args(cmd, []string{"test me"})
	if err == nil {
		t.Errorf("Expected error got nil")
	}
	testutil.Equal(t, "hello", cmd.Use)

}

func TestNewCmdDescription(t *testing.T) {
	description := "team are here test description"
	cmd := NewCmd("team").WithDescription(description).NoArgs(nil)
	if cmd.Short != description {
		t.Errorf("Expected %s error got %s", description, cmd.Short)
	}
	testutil.Equal(t, "team", cmd.Use)
}

func TestNewCmdLongDescription(t *testing.T) {
	longDescription := "testing long description"
	cmd := NewCmd("").WithLongDescription(longDescription).NoArgs(nil)
	testutil.Equal(t, longDescription, cmd.Long)
	testutil.Equal(t, "", cmd.Use)
}
