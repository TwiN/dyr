package delete

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

func NewDeleteCmd() *cobra.Command {
	var cmd = &cobra.Command{
		Use:     "delete NOTE_ID",
		Aliases: []string{"del", "remove"},
		Short:   "Delete a note",
		Long:    "Delete a note",
		Example: "dyr delete 1",
		Args:    cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			id, err := strconv.Atoi(args[0])
			if err != nil {
				panic(err)
			}
			fmt.Printf("Deleting note with id %d\n", id)
		},
	}

	return cmd
}
