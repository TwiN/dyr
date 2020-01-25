package create

import (
	"fmt"
	"github.com/TwinProduction/dyr/core"
	"github.com/spf13/cobra"
	"strings"
)

func NewCreateCmd() *cobra.Command {
	note := &core.Note{
		Id:   0,
		Text: "",
		Tags: []string{},
	}

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
				note.Text = strings.Join(args, " ")
			}
			SaveNote(note)
		},
	}

	cmd.Flags().StringSliceVarP(&note.Tags, "tags", "t", note.Tags, "Comma-separated list of tags")

	return cmd
}

func SaveNote(note *core.Note) {
	fmt.Printf("%v\n", note)
}
