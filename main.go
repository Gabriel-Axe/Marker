package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {

	cwd, err := os.Getwd()
	if err != nil {
		fmt.Printf("Could not get current directory: %v", err)
	}

	if len(os.Args) != 3 {
		fmt.Println("usage: k <path/to/template> <name>")
		return
	}

	// templatePath := os.Args[1]
	newNoteName := os.Args[2]

	newNotePath := cwd + "/" + newNoteName + ".md"
	fmt.Println(newNotePath)

	// err = createFileFromTemplate(templatePath, newNoteName)
	// if err != nil {
	// 	fmt.Printf("Could not create file: %v", err)
	// }

	editor := os.Getenv("VISUAL")
	if editor == "" {
		editor = os.Getenv("EDITOR")
	}
	if editor == "" {
		editor = "vi"
	}

	cmd := exec.Command(editor, )
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	// cmd.Run()
}

func createFileFromTemplate(templatePath, newPath string) error {
	data, err := os.ReadFile(templatePath)
	if err != nil {
		return err
	}

	return os.WriteFile(newPath, data, 0644)
}
