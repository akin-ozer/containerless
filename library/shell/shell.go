package shell

import (
	"bufio"
	"fmt"
	"log"
	"os/exec"
	"strings"
)

func Execute(command string) {
	cmd := exec.Command("sh", command)
	stdoutStderr, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", stdoutStderr)
}

func Piped(command string) {
	cmd := exec.Command("sh", "-c", command)
	pipe, _ := cmd.StdoutPipe()
	cmd.Stderr = cmd.Stdout

	if err := cmd.Start(); err != nil {
		fmt.Println(err.Error())
	}

	reader := bufio.NewReader(pipe)
	line, err := reader.ReadString('\n')
	for err == nil {
		fmt.Println(line)
		line, err = reader.ReadString('\n')
	}

}

func PipedStdin(command string, stdin string) {
	cmd := exec.Command("sh", "-c", command)
	cmd.Stdin = strings.NewReader(stdin)
	pipe, _ := cmd.StdoutPipe()
	cmd.Stderr = cmd.Stdout

	scanner := bufio.NewScanner(pipe)
	go func() {
		for scanner.Scan() {
			fmt.Println(scanner.Text())
		}
	}()
	if err := cmd.Start(); err != nil {
		fmt.Println(err.Error())
	}
	cmd.Wait()
	//reader := bufio.NewReader(pipe)
	//line, err := reader.ReadString('\n')
	//for err == nil {
	//	fmt.Println(line)
	//	line, err = reader.ReadString('\n')
	//}

}
