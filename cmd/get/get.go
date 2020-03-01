package get

import (
	"fmt"
	"github.com/TwinProduction/dyr/config"
	"github.com/TwinProduction/dyr/core"
	"github.com/TwinProduction/dyr/storage"
	"github.com/spf13/cobra"
	"strconv"
	"strings"
)

type Options struct {
	Verbose bool
	Banner  bool
	Random  bool
}

func NewGetCmd() *cobra.Command {
	options := Options{
		Banner:  true,
		Verbose: false,
		Random:  false,
	}

	var cmd = &cobra.Command{
		Use:     "get [NOTE_ID]",
		Short:   "Get an entry by id",
		Long:    "Get an entry by id",
		Example: "dyr get 1 # Print one note\ndyr get # Print all notes\ndyr get --random # Print one random note",
		Args:    cobra.RangeArgs(0, 1),
		Run: func(cmd *cobra.Command, args []string) {
			var notes []*core.Note
			if len(args) == 0 {
				if options.Random {
					note, err := storage.GetRandomNote()
					if err != nil {
						panic(err)
					}
					notes = append(notes, note)
				} else {
					var err error
					notes, err = storage.GetAllNotes()
					if err != nil {
						panic(err)
					}
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
			isPrintingAllNotes := len(args) == 0 && !options.Random
			for _, note := range notes {
				if options.Verbose {
					fmt.Printf("[%d]\nTAGS: %s\nTEXT: %s\n", note.Id, strings.Join(note.Tags, ","), note.Text)
				} else {
					// Print the ID only if we print more than one note
					if isPrintingAllNotes {
						fmt.Printf("%d: %s\n", note.Id, note.Text)
					} else {
						fmt.Println(note.Text)
					}
				}
			}
		},
	}

	cmd.Flags().BoolVarP(&options.Verbose, "verbose", "v", options.Verbose, "Whether to print only the text")
	cmd.Flags().BoolVarP(&options.Random, "random", "r", options.Random, "Whether to print a random note. Ignored if an id is specified")
	cmd.Flags().BoolVarP(&options.Banner, "banner", "b", options.Banner, "Whether to show the banner configured with... TODO: make a command for creating a banner")

	return cmd
}
