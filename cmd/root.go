package cmd

import (
	"fmt"
	"github.com/TwinProduction/dyr/config"
	"github.com/spf13/cobra"
	"os"
)

func NewRootCmd() *cobra.Command {
	var rootCmd = &cobra.Command{
		Use:     "dyr",
		Short:   "dyr",
		Long:    "dyr",
		Version: "v0.0.1",
		Run: func(cmd *cobra.Command, args []string) {
			_ = cmd.Help()
		},
	}
	return rootCmd
}

func Execute() {
	rootCmd := NewRootCmd()

	err := config.LoadConfiguration()
	if err != nil {
		panic(err)
	}

	if err := rootCmd.Execute(); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, err.Error())
		os.Exit(1)
	}
}
