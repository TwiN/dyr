# dyr

Do you remember? is a tool meant to help you remember these nifty tricks you once learned.

As a programmer, I keep learning new useful tricks, but unfortunately, they escape my memory every now and then -- sometimes to be rediscovered at a later date, others to be simply forgotten in the river of time.

This very simple command line tool's purpose is very simple: help you save and remind you of these tricks.

Obviously, it doesn't have to be useful tricks, it can be anything you want it to be:
- Good memories
- Accomplishments 
- Jokes
- Quotes
- Other stuffs

I personally add `dyr get --random` at the bottom of my `.bashrc` so that every time I open a new terminal, I get reminded.


## Usage

### Creating a note

```
dyr create NOTE_TO_ADD [flags]
```

Flags:
```
  -h, --help           help for create
  -t, --tags strings   Comma-separated list of tags
```

Example:
```
dyr create :term allows you to open a terminal in vim --tags programming,vim
```

### Reading existing note(s)

```
dyr get [NOTE_ID] [flags]
```

Flags:
```
  -b, --banner       Whether to show the banner configured with... TODO: make a command for creating a banner (default true)
  -h, --help         help for get
  -r, --random       Whether to print a random note. Ignored if an id is specified
  -t, --tag string   List all nodes with the provided tag. Ignored if an id is specified
  -v, --verbose      Whether to print only the text
```

List all notes:
```
dyr get
```

List a single note with the specified id:
```
dyr get 1
```

List a single random note:
```
dyr get --random
```

List all notes tagged with `test`:
```
dyr get --tag test
```


### Deleting a note

```
dyr remove NOTE_ID
```


## Useful reminders

Here's a list of tips I kept forgetting, which is what prompted 
the creation of this project:
- You can quickly move the current line down with vim using "ddp"
- You can open a terminal in vim using ":term"

