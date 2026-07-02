package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
)

func main() {

	cwd, err := os.Getwd()
	if err != nil {
		fmt.Printf("Could not get current directory: %v", err)
	}

	flag.Parse()
	args := flag.Args()

	if len(args) < 2 || len(args) > 3 {
		fmt.Printf("usage: k <path/to/template> <name> [-open], number of args: %d\n", len(args))
		os.Exit(1)
	}

	templatePath := os.Args[1]
	newNoteName := os.Args[2]
	var open bool = true

	newNotePath := cwd + "/" + newNoteName + ".md"
	fmt.Println(newNotePath)

	err = createFileFromTemplate(templatePath, newNotePath)
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

func createFileFromTemplate(templatePath, newPath string) error {
	data, err := os.ReadFile(templatePath)
	if err != nil {
		return err
	}

	return os.WriteFile(newPath, data, 0644)
}
