package main

import (
	"fmt"
	"io/ioutil"
)

func Exec(args []string) error {
	file, err := ioutil.ReadFile(args[1])
	if err != nil {
		return err
	}
	fmt.Println(string(file))
	return nil
}
