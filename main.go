package main

import (
	"fmt"
	"my-project-module/subfolder"
	"my-project-module/subfolder/subsubfolder"
)

func main() {
	fmt.Println("Main file in my-project-module")
	subfolder.PrintHello()
	subfolder.PrintWelcome()
	subsubfolder.PrintThanks()
}
