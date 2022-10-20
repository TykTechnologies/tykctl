package cmd

import (
	"context"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

type Builder interface {
	WithExample(comment string) Builder
	Hidden() Builder
	WithDescription(description string) Builder
	WithLongDescription(long string) Builder
	WithCommands(cmds ...*cobra.Command) *cobra.Command
	NoArgs(action func(context.Context, cobra.Command) error) *cobra.Command
	AddFlags([]*Flag) Builder
	WithFlagAdder(persistent bool, adder func(*pflag.FlagSet)) Builder
}

type builder struct {
	cmd cobra.Command
	key string
}

func (b *builder) WithFlagAdder(persistent bool, adder func(*pflag.FlagSet)) Builder {
	if persistent {
		adder(b.cmd.PersistentFlags())
	} else {
		adder(b.cmd.Flags())
	}

	return b
}

func (b *builder) AddFlags(flags []*Flag) Builder {
	for _, f := range flags {
		fl := f.flag()
		b.cmd.Flags().AddFlag(fl)

	}
	return b
}

// NewCmd Creates a new command builder
// key is unique for each command and will be used for global flags
func NewCmd(key, use string) Builder {
	return &builder{
		key: key,
		cmd: cobra.Command{
			Use: use,
		},
	}

}

func (b *builder) WithExample(comment string) Builder {
	if b.cmd.Example != "" {
		b.cmd.Example += "\n"
	}
	b.cmd.Example += comment
	return b
}

func (b *builder) Hidden() Builder {
	b.cmd.Hidden = true
	return b
}

func (b *builder) WithDescription(description string) Builder {
	b.cmd.Short = description
	return b
}

func (b *builder) WithLongDescription(long string) Builder {
	b.cmd.Long = long
	return b
}

func (b *builder) WithCommands(cmds ...*cobra.Command) *cobra.Command {
	for _, cmd := range cmds {
		b.cmd.AddCommand(cmd)

	}
	return &b.cmd
}

func (b *builder) NoArgs(action func(context.Context, cobra.Command) error) *cobra.Command {
	b.cmd.RunE = func(*cobra.Command, []string) error {
		return action(b.cmd.Context(), b.cmd)
	}
	return &b.cmd
}
