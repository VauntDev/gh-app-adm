package cmd

import (
	"context"

	"github.com/elewis787/boa"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"golang.org/x/sync/errgroup"
)

func Execute() error {
	// root command entry to application
	rootCmd := &cobra.Command{
		Use:     "gh-app-adm",
		Version: "v0.0.1",
		Short:   "Github Application admin cli",
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			// Bind Cobra flags with viper
			if err := viper.BindPFlags(cmd.Flags()); err != nil {
				return err
			}
			return nil
		},
	}

	// Add start command
	rootCmd.AddCommand(jwtCmd())
	rootCmd.AddCommand(installations())

	rootCmd.SetUsageFunc(boa.UsageFunc)
	rootCmd.SetHelpFunc(boa.HelpFunc)

	ctx, cancel := context.WithCancel(context.Background())
	errGrp, errctx := errgroup.WithContext(ctx)
	errGrp.Go(func() error {
		defer cancel()
		if err := rootCmd.ExecuteContext(errctx); err != nil {
			return err
		}
		return nil
	})
	return errGrp.Wait()
}
