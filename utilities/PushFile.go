package utilities

import (
	"bufio"
	"fmt"
	"log"
	"os/exec"
)

func PushFile() {
	fmt.Println("Pushing")
	RunLiveCommand("git", "push")
}

func RunLiveCommand(name string, arg ...string) {
	cmd := exec.Command(name, arg...)
	pipe, _ := cmd.StdoutPipe()
	if err := cmd.Start(); err != nil {
		log.Fatal(err)
	}
	reader := bufio.NewReader(pipe)
	line, err := reader.ReadString('\n')
	for err == nil {
		fmt.Println(line)
		line, err = reader.ReadString('\n')
	}
}
