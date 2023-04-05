package utilities

import (
	"fmt"
	"log"
	"os/exec"
)

func GitCommit(message string) {
	out, err := exec.Command("git", "commit", "-m", message).CombinedOutput()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(out)
}
