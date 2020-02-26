package remove

import (
	"fmt"
	"strconv"

	"github.com/TwinProduction/dyr/storage"
	"github.com/spf13/cobra"
)

func NewRemoveCmd() *cobra.Command {
	var cmd = &cobra.Command{
		Use:     "remove NOTE_ID",
		Aliases: []string{"delete"},
		Short:   "Delete a note",
		Long:    "Delete a note",
		Example: "dyr remove 1",
		Args:    cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			id, err := strconv.Atoi(args[0])
			if err != nil {
				panic(err)
			}
			fmt.Printf("Deleting note with id %d\n", id)
			err = storage.DeleteNoteById(uint64(id))
			if err != nil {
				panic(err)
			}
		},
	}

	return cmd
}
