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
	WithFlagAdder(persistent bool, adder func(*pflag.FlagSet)) Builder
}

type builder struct {
	cmd cobra.Command
}

// NewCmd Creates a new command builder.
func NewCmd(use string) Builder {
	return &builder{

		cmd: cobra.Command{
			Use: use,
		},
	}

}

// WithFlagAdder adds flags to the cmd flags
// if persistent is set to true the flags will be added as PersistentFlags
// otherwise they will be added as local flags.
func (b *builder) WithFlagAdder(persistent bool, adder func(*pflag.FlagSet)) Builder {
	if persistent {
		adder(b.cmd.PersistentFlags())
	} else {
		adder(b.cmd.Flags())
	}

	return b
}

// WithExample will set an example of how to use the command.
func (b *builder) WithExample(comment string) Builder {
	if b.cmd.Example != "" {
		b.cmd.Example += "\n"
	}
	b.cmd.Example += comment
	return b
}

// Hidden defines whether to remove the command from the list of available commands.
func (b *builder) Hidden() Builder {
	b.cmd.Hidden = true
	return b
}

// WithDescription sets a short description to our command.
func (b *builder) WithDescription(description string) Builder {
	b.cmd.Short = description
	return b
}

// WithLongDescription sets a long description to our cmd.
func (b *builder) WithLongDescription(long string) Builder {
	b.cmd.Long = long
	return b
}

// WithCommands adds subcommand to this parent command.
func (b *builder) WithCommands(cmds ...*cobra.Command) *cobra.Command {
	for _, cmd := range cmds {
		b.cmd.AddCommand(cmd)
	}
	return &b.cmd
}

// NoArgs is for when you want to execute the cmd with zero args.
func (b *builder) NoArgs(action func(context.Context, cobra.Command) error) *cobra.Command {
	b.cmd.Args = cobra.NoArgs
	b.cmd.RunE = func(*cobra.Command, []string) error {
		return action(b.cmd.Context(), b.cmd)
	}
	return &b.cmd
}
