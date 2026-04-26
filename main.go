package main

import (
	"fmt"
	"os"
)

func main() {
	for _, arg := range(os.Args) {
		fmt.Println(arg)
	}
	// if len(os.Args) < 2 {
	// 	fmt.Println("no lol")
	// 	return
	// }
	// fmt.Println("yes")
}
