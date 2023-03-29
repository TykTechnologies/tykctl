package internal

import (
	"context"
	"errors"
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

var ErrCurrentUserNotFound = errors.New("current user not set/found")

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
	WithValidArgs(args []string) Builder
	AddPreRunFuncs(...func(cmd *cobra.Command, args []string) error) Builder
}

type builder struct {
	cmd                     cobra.Command
	runOnPreRun             []func(cmd *cobra.Command, args []string) error
	bindOnPreRun            []BindFlag
	bindWithContextOnPreRun []BindFlag
}

// NewCmd Creates a new command builder.
func NewCmd(use string) Builder {
	return &builder{
		cmd: cobra.Command{
			Use:          use,
			SilenceUsage: true,
			Version:      Version,
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
	b.bindOnPreRun = append(b.bindOnPreRun, flags...)
	return b
}

func (b *builder) bindFlagonPreRun() error {
	for _, flag := range b.bindOnPreRun {
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

// WithBindFlagWithCurrentUserContext will help us get the current user flag since viper is initialized on cobra PreRun.
func (b *builder) WithBindFlagWithCurrentUserContext(flags []BindFlag) Builder {
	b.bindWithContextOnPreRun = append(b.bindWithContextOnPreRun, flags...)
	return b
}

// bindFlagonPreRunWithCurrentContext which get the parameters from the current logged user
// this will allow us to support different users in the config.
func (b *builder) bindFlagonPreRunWithCurrentContext() error {
	if len(b.bindWithContextOnPreRun) == 0 {
		return nil
	}

	currentUser := viper.GetString(currentCloudUser)
	if currentUser == "" {
		return ErrCurrentUserNotFound
	}

	for _, flag := range b.bindWithContextOnPreRun {
		currentUserCtx := fmt.Sprintf("cloud.%s.%s", currentUser, flag.Name)

		var err error

		if flag.Persistent {
			err = viper.BindPFlag(currentUserCtx, b.cmd.PersistentFlags().Lookup(flag.Name))
		} else {
			err = viper.BindPFlag(currentUserCtx, b.cmd.Flags().Lookup(flag.Name))
		}

		if err != nil {
			return err
		}
	}

	return nil
}

// executePreRunFuncs will run all the function scheduled to be run during cmd preRun.
func (b *builder) executePreRunFuncs(cmd *cobra.Command, args []string) error {
	for _, f := range b.runOnPreRun {
		err := f(cmd, args)
		if err != nil {
			return err
		}
	}

	return nil
}

func (b *builder) PreRun(cmd *cobra.Command, args []string) error {
	err := b.bindFlagonPreRun()
	if err != nil {
		return err
	}

	err = b.bindFlagonPreRunWithCurrentContext()
	if err != nil {
		return err
	}

	return b.executePreRunFuncs(cmd, args)
}

// NoArgs is for when you want to execute the cloudCmd with zero args.
func (b *builder) NoArgs(action func(context.Context, cobra.Command) error) *cobra.Command {
	b.cmd.Args = cobra.NoArgs
	b.cmd.PreRunE = b.PreRun
	b.cmd.RunE = func(*cobra.Command, []string) error {
		return action(b.cmd.Context(), b.cmd)
	}

	return &b.cmd
}

// MaximumArgs fails if you pass args that are more than the specified maxArgCount.
func (b *builder) MaximumArgs(maxArgCount int, action func(context.Context, cobra.Command, []string) error) *cobra.Command {
	b.cmd.Args = cobra.MaximumNArgs(maxArgCount)
	b.cmd.PreRunE = b.PreRun
	b.cmd.RunE = func(cmd *cobra.Command, args []string) error {
		return action(b.cmd.Context(), b.cmd, args)
	}

	return &b.cmd
}

// ExactArgs fails if you pass args that are more or less than the specified argCount.
func (b *builder) ExactArgs(argCount int, action func(context.Context, cobra.Command, []string) error) *cobra.Command {
	b.cmd.Args = cobra.ExactArgs(argCount)
	b.cmd.PreRunE = b.PreRun
	b.cmd.SilenceUsage = true
	b.cmd.RunE = func(cmd *cobra.Command, args []string) error {
		return action(b.cmd.Context(), b.cmd, args)
	}

	return &b.cmd
}

func (b *builder) WithAliases(aliases []string) Builder {
	b.cmd.Aliases = aliases
	return b
}

func (b *builder) WithValidArgs(args []string) Builder {
	b.cmd.ValidArgs = args
	return b
}

func (b *builder) AddPreRunFuncs(items ...func(cmd *cobra.Command, args []string) error) Builder {
	b.runOnPreRun = append(b.runOnPreRun, items...)
	return b
}
