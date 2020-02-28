package cmd

import (
	"fmt"
	"os"

	"github.com/TwinProduction/dyr/cmd/create"
	"github.com/TwinProduction/dyr/cmd/get"
	"github.com/TwinProduction/dyr/cmd/remove"
	"github.com/TwinProduction/dyr/config"
	"github.com/spf13/cobra"
)

func NewRootCmd() *cobra.Command {
	var rootCmd = &cobra.Command{
		Use:     "dyr",
		Short:   "dyr",
		Long:    "dyr - Do you remember?",
		Version: "v0.0.1",
		Run: func(cmd *cobra.Command, args []string) {
			_ = cmd.Help()
		},
	}
	return rootCmd
}

func Execute() {
	rootCmd := NewRootCmd()

	err := config.Load()
	if err != nil {
		panic(err)
	}

	rootCmd.AddCommand(create.NewCreateCmd())
	rootCmd.AddCommand(get.NewGetCmd())
	rootCmd.AddCommand(remove.NewRemoveCmd())

	if err := rootCmd.Execute(); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
}
