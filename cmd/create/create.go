package create

import (
	"github.com/TwinProduction/dyr/storage"
	"github.com/spf13/cobra"
	"strings"
)

func NewCreateCmd() *cobra.Command {
	var tags []string

	var cmd = &cobra.Command{
		Use:     "create NOTE_TO_ADD",
		Aliases: []string{"add", "save"},
		Short:   "Create a new note",
		Long:    "Create a new note",
		Example: "dyr create :term allows you to open a terminal in vim --tags programming,vim",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) == 0 {
				_ = cmd.Help()
				return
			} else {
				SaveNote(strings.Join(args, " "), tags)
			}
		},
	}

	cmd.Flags().StringSliceVarP(&tags, "tags", "t", tags, "Comma-separated list of tags")

	return cmd
}

func SaveNote(text string, tags []string) {
	err := storage.SaveNote(text, tags)
	if err != nil {
		panic(err)
	}
}
