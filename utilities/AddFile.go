package utilities

import (
	"fmt"
	"os"
	"os/exec"
)

func AddFile(file string) {
	if file == "" {
		fmt.Println("That File was not Found! What is the name of the file you would like to add?")

		var file string
		fmt.Scanln(&file)

		GitAddFile(file)
	} else {
		if _, err := os.Stat(file); os.IsNotExist(err) {
			AddFile("")
		} else {
			GitAddFile(file)
		}
	}
}

func GitAddFile(file string) {
	out, err := exec.Command("git", "add", file).CombinedOutput()
	if err != nil {
		fmt.Printf("The File (%s) was not found! Try Again\n", file)
		AddFile("")
	}
	fmt.Println(string(out))
	fmt.Printf("%s Added\n", file)
}
