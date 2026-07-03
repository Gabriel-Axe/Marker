package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"os/exec"
	// "path"
	"path/filepath"
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

type Path struct {
	URI string
}

func main() {
	cwd, err := os.Getwd()
	if err != nil {
		fmt.Printf("Could not get current directory: %v", err)
	}

	mflags := parseFlags()
	notePath := cwd + "/" + mflags.NoteName
	note := createNoteStruct(mflags.NoteName, notePath)

	defer executeEditor(mflags, note)
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

func executeEditor(mflags Flags, note *Note) error {
	if mflags.OpenEditor {
		editor := os.Getenv("VISUAL")
		if editor == "" {
			editor = os.Getenv("EDITOR")
		}
		if editor == "" {
			editor = "vi"
		}

		cmd := exec.Command(editor, note.Path.URI)
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		err := cmd.Run()
		if err != nil {
			return fmt.Errorf("Could not open editor: %v", err)
		}
	}
	return nil
}

func printDefaultCommandCall() {
	fmt.Println("usage: k -template <path> <path> -name <name> [-open]")
}

// Returns a struct rerpresenting a template and the name of the new note
func parseFlags() Flags {
	templatePath := flag.String("template", "", "path to template file")
	noteName := flag.String("name", "", "name of the new note")
	open := flag.Bool("open", false, "open the note after creation")

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
}

func createNoteStruct(name string, absolutePath string) *Note {

	path := Path { URI: absolutePath }
	path = sanitizePath(path)

	return &Note {
		Name: name,
		Path: path,
	}
}
func sanitizePath(path Path) Path {
	if strings.HasSuffix(path.URI, ".md") {
		path.URI = strings.TrimSuffix(path.URI, ".md")
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
