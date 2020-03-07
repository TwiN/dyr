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
	Tag     string
}

func NewGetCmd() *cobra.Command {
	options := Options{
		Banner:  true,
		Verbose: false,
		Random:  false,
	}

	var cmd = &cobra.Command{
		Use:   "get [NOTE_ID]",
		Short: "Get an entry by id",
		Long:  "Get an entry by id",
		Example: `  # List all notes
  dyr get

  # List a single note with the specified id
  dyr get 1

  # List a single random note
  dyr get --random

  # List all notes tagged with "test"
  dyr get --tag test`,
		Args: cobra.RangeArgs(0, 1),
		Run: func(cmd *cobra.Command, args []string) {
			var notes []*core.Note
			if len(args) == 0 {
				var err error
				if options.Tag != "" {
					notes, err = storage.GetNotesByTag(options.Tag)
					if err != nil {
						panic(err)
					}
				} else if options.Random {
					note, err := storage.GetRandomNote()
					if err != nil {
						panic(err)
					}
					notes = append(notes, note)
				} else {
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
	cmd.Flags().StringVarP(&options.Tag, "tag", "t", options.Tag, "List all nodes with the provided tag. Ignored if an id is specified")
	// TODO: Allow --tag and --random to be used together

	return cmd
}
