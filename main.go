package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

type Template struct {
	Path string
}

type Note struct {
	Path string
	Name string
}

func main() {

	cwd, err := os.Getwd()
	if err != nil {
		fmt.Printf("Could not get current directory: %v", err)
	}

	template, name := parseFlags()

	newNotePath := cwd + "/" + newNoteName + ".md"
	fmt.Println(newNotePath)

	if strings.HasSuffix(name, ".md") {
		name = strings.TrimSuffix(name, ".md")
		print(name)
	}

	err = createFileFromTemplate(template, newNotePath)
	if err != nil {
		fmt.Printf("Could not create file: %v", err)
	}

	print(open)
	if open == true {
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
	args := flag.Args()

	if len(args) < 2 || len(args) > 3 {
		fmt.Printf("usage: k <path/to/template> <name> [-open], number of args: %d\n", len(args))
		os.Exit(1)
	}

	template := Template { Path: os.Args[1] }
	newNoteName := os.Args[2]

	return template, newNoteName
}

func createFileFromTemplate(template Template, newPath string) error {
	data, err := os.ReadFile(template.Path)
	if err != nil {
		return err
	}

	return os.WriteFile(newPath, data, 0644)
}
