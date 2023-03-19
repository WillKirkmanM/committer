package utilities

import (
	"fmt"
	"log"
)

func Commit() {
	message, scope, err := ParseCliArgs()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Scope & Message")
	fmt.Println(scope, message)
}
