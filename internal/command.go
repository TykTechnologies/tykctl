package internal

import (
	"context"
	"errors"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

var (
	ErrCurrentUserNotFound = errors.New("you have not yet set the current user")
)

type Builder interface {
	WithExample(comment string) Builder
	Hidden() Builder
	WithDescription(description string) Builder
	WithLongDescription(long string) Builder
	WithCommands(cmds ...*cobra.Command) *cobra.Command
	NoArgs(action func(context.Context, cobra.Command) error) *cobra.Command
	MaximumArgs(maxArgCount int, action func(context.Context, cobra.Command, []string) error) *cobra.Command
	ExactArgs(argCount int, action func(context.Context, cobra.Command, []string) error) *cobra.Command
	WithFlagAdder(persistent bool, adder func(*pflag.FlagSet)) Builder
	WithBindFlagOnPreRun(flags []BindFlag) Builder
	WithBindFlagWithCurrentUserContext([]BindFlag) Builder
	WithAliases(aliases []string) Builder
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

// WithFlagAdder adds flags to the cloudcmd flags
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

// WithLongDescription sets a long description to our cloudcmd.
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

// WithBindFlagOnPreRun helps us bind flags before preRun
// this help us solve ths cobra issue https://github.com/spf13/viper/issues/233.
func (b *builder) WithBindFlagOnPreRun(flags []BindFlag) Builder {
	b.cmd.PreRunE = func(cmd *cobra.Command, args []string) error {
		for _, flag := range flags {
			if flag.Persistent {
				err := viper.BindPFlag(flag.Name, b.cmd.PersistentFlags().Lookup(flag.Name))
				if err != nil {
					return err
				}
			} else {
				err := viper.BindPFlag(flag.Name, b.cmd.Flags().Lookup(flag.Name))
				if err != nil {
					return err
				}
			}
		}
		return nil
	}
	return b
}

// WithBindFlagWithCurrentUserContext will help us get the current user flag since viper is initialized on cobra PreRun.
func (b *builder) WithBindFlagWithCurrentUserContext(flags []BindFlag) Builder {
	b.cmd.PreRunE = func(cmd *cobra.Command, args []string) error {
		currentUser := viper.GetString(currentCloudUser)
		if currentUser == "" {
			return ErrCurrentUserNotFound
		}
		for _, flag := range flags {
			currentUserCtx := fmt.Sprintf("cloud.%s.%s", currentUser, flag.Name)
			if flag.Persistent {
				err := viper.BindPFlag(currentUserCtx, b.cmd.PersistentFlags().Lookup(flag.Name))
				if err != nil {
					return err
				}
			} else {
				err := viper.BindPFlag(currentUserCtx, b.cmd.Flags().Lookup(flag.Name))
				if err != nil {
					return err
				}
			}
			return nil
		}
		return nil
	}
	return b
}

// NoArgs is for when you want to execute the cloudcmd with zero args.
func (b *builder) NoArgs(action func(context.Context, cobra.Command) error) *cobra.Command {
	b.cmd.Args = cobra.NoArgs
	b.cmd.RunE = func(*cobra.Command, []string) error {
		return action(b.cmd.Context(), b.cmd)
	}
	return &b.cmd
}

// MaximumArgs fails if you pass args that are more than the specified maxArgCount.
func (b *builder) MaximumArgs(maxArgCount int, action func(context.Context, cobra.Command, []string) error) *cobra.Command {
	b.cmd.Args = cobra.MaximumNArgs(maxArgCount)
	b.cmd.RunE = func(cmd *cobra.Command, args []string) error {
		return action(b.cmd.Context(), b.cmd, args)
	}
	return &b.cmd
}

// ExactArgs fails if you pass args that are more or less than the specified argCount.
func (b *builder) ExactArgs(argCount int, action func(context.Context, cobra.Command, []string) error) *cobra.Command {
	b.cmd.Args = cobra.ExactArgs(argCount)
	b.cmd.RunE = func(cmd *cobra.Command, args []string) error {
		return action(b.cmd.Context(), b.cmd, args)
	}
	return &b.cmd
}

func (b *builder) WithAliases(aliases []string) Builder {
	b.cmd.Aliases = aliases
	return b
}
