package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

type Config struct {
	TemplatesDir string
	NotesDir string
	// Editor string
}
type Flags struct {
	TemplatePath string
	NoteName string
	OpenEditor bool
}

type Template struct {
	Path Path
}

type Note struct {
	Path Path
	Name string
}

func main() {

	cwd, err := os.Getwd()
	if err != nil {
		fmt.Printf("Could not get current directory: %v", err)
	}

	mflags := parseFlags()
	notePath := cwd + "/" + mflags.NoteName
	note := createNoteStruct(mflags.NoteName, notePath)

	templatePath := Path{ mflags.TemplatePath }
	template := Template{ Path: templatePath }

	err = createFileFromTemplate(template, *note)
	if err != nil {
		fmt.Printf("Could not create file: %v", err)
	}
}

func createFlagStruct(templatePath string, noteName string, openEditor bool) *Flags {

	return &Flags {
		TemplatePath : templatePath,
		NoteName : noteName,
		OpenEditor : openEditor,
	}
}
		editor := os.Getenv("VISUAL")
		if editor == "" {
			editor = os.Getenv("EDITOR")
		}
		if editor == "" {
			editor = "vi"
		}

		cmd := exec.Command(editor, newNotePath)
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		err = cmd.Run()
		if err != nil {
			fmt.Printf("Could not open editor: %v", err)
		}
	}
}

// Returns a struct rerpresenting a template and the name of the new note
func parseFlags() (Template, string) {
	flag.Parse()

	if *templatePath == "" || *noteName == "" {
		printDefaultCommandCall()
		os.Exit(1)
	}

	return Flags {
		TemplatePath: *templatePath,
		NoteName: *noteName,
		OpenEditor: *open,
	}

	return path
}

func createFileFromTemplate(template Template, note Note) error {
	data, err := os.ReadFile(template.Path.URI)
	if err != nil {
		return err
	}

	path := note.Path.URI

	f, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_EXCL, 0644)
	if err != nil {
		if os.IsExist(err) {
			return fmt.Errorf("file %s already exists", path)
		}
	}

	defer f.Close()
	_, err = f.Write(data)
	return err
}
}
