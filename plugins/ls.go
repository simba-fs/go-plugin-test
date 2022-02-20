package main

import "os/exec"
import "os"

func Exec(args []string) error {
	ls := exec.Command("ls", args[1:]...)
	ls.Stdout = os.Stdout
	ls.Stderr = os.Stderr
	return ls.Run()
}
