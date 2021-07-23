package lib

import (
	"bufio"
	"fmt"
	"os/exec"
)

func Pull(imageName string, tag string) {
	runCmd(exec.Command("docker", "pull", imageName + ":" + tag))
}

func Delete(imageName string, tag string) {
	runCmd(exec.Command("docker", "image", "rm", imageName + ":" + tag))
}

func runCmd(cmd *exec.Cmd) {
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		panic(err)
	}
	stderr, err := cmd.StderrPipe()
	if err != nil {
		panic(err)
	}

	outScanner := bufio.NewScanner(stdout)
	errScanner := bufio.NewScanner(stderr)
	go func() {
		for outScanner.Scan() {
			fmt.Printf("  %s\n", outScanner.Text())
		}
	}()
	go func() {
		for errScanner.Scan() {
			fmt.Printf("  %s\n", outScanner.Text())
		}
	}()

	cmd.Start()
	cmd.Wait()
}
