package main

import (
	"fmt"
	"os/exec"
)

func main() {
	var (
		cmd *exec.Cmd
		err error
	)
	cmd = exec.Command("C:\\cygwin64\\bin\\bash.exe", "-c", "echo asdhhsdakdad")
	err = cmd.Run()
	fmt.Println(err)
}
