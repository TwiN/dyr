package get

import (
	"fmt"
	"github.com/TwinProduction/dyr/config"
	"github.com/TwinProduction/dyr/storage"
	"github.com/spf13/cobra"
	"strconv"
)

type Options struct {
	TextOnly bool
	Banner   bool
}

func NewGetCmd() *cobra.Command {
	options := Options{
		Banner: true,
	}

	var cmd = &cobra.Command{
		Use:     "get NOTE_ID",
		Short:   "Get an entry by id",
		Long:    "Get an entry by id",
		Example: "dyr get 1",
		Args:    cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			id, err := strconv.Atoi(args[0])
			if err != nil {
				panic(err)
			}
			note, err := storage.GetNoteById(uint64(id))
			if err != nil {
				panic(err)
			}
			if options.Banner {
				banner := config.Get().Banner
				if len(banner) > 0 {
					fmt.Println(config.Get().Banner)
				}
			}
			if options.TextOnly {
				fmt.Println(note.Text)
			} else {
				fmt.Printf("ID: %d\nTAGS: %s\nTEXT: %s\n", note.Id, note.Tags, note.Text)
			}
		},
	}

	cmd.Flags().BoolVarP(&options.TextOnly, "text-only", "t", options.TextOnly, "Whether to print only the text")
	cmd.Flags().BoolVarP(&options.Banner, "banner", "b", options.Banner, "Whether to show the banner configured with... TODO: make a command for creating a banner")

	return cmd
}
