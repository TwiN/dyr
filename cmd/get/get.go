package get

import (
	"fmt"
	"github.com/TwinProduction/dyr/config"
	"github.com/TwinProduction/dyr/core"
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
		Banner:   true,
		TextOnly: true,
	}

	var cmd = &cobra.Command{
		Use:     "get [NOTE_ID]",
		Short:   "Get an entry by id",
		Long:    "Get an entry by id",
		Example: "dyr get 1",
		Args:    cobra.RangeArgs(0, 1),
		Run: func(cmd *cobra.Command, args []string) {
			var notes []*core.Note
			if len(args) == 0 {
				var err error
				notes, err = storage.GetAllNotes()
				if err != nil {
					panic(err)
				}
			} else {
				id, err := strconv.Atoi(args[0])
				if err != nil {
					panic(err)
				}
				note, err := storage.GetNoteById(uint64(id))
				if err != nil {
					panic(err)
				}
				notes = append(notes, note)
			}
			if options.Banner {
				banner := config.Get().Banner
				if len(banner) > 0 {
					fmt.Println(config.Get().Banner)
				}
			}
			isPrintingAll := len(args) == 0
			for _, note := range notes {
				if options.TextOnly {
					if isPrintingAll {
						fmt.Printf("%d: %s\n", note.Id, note.Text)
					} else {
						fmt.Println(note.Text)
					}
				} else {
					fmt.Printf("ID: %d\nTAGS: %s\nTEXT: %s\n", note.Id, note.Tags, note.Text)
				}
			}
		},
	}

	cmd.Flags().BoolVarP(&options.TextOnly, "text-only", "t", options.TextOnly, "Whether to print only the text")
	cmd.Flags().BoolVarP(&options.Banner, "banner", "b", options.Banner, "Whether to show the banner configured with... TODO: make a command for creating a banner")

	return cmd
}
