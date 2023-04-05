package utilities

import (
	"fmt"
	"os"
)

func Confirm(message string) {
	fmt.Printf("Are you sure you would like to %s? (Y/n)\n", message)

	var answer string

	fmt.Scanln(&answer)

	if answer != "Y" && answer != "y" {
		os.Exit(0)
	}
}
