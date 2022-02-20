package main

import (
	"bufio"
	"fmt"
	"plugin"
	"strings"
	"os"
)

func main() {
	stdin := bufio.NewReader(os.Stdin)
	for {
		fmt.Printf("> ")
		rawCmd, err := stdin.ReadString('\n')
		if err != nil {
			fmt.Println(err)
			continue
		}

		cmd := strings.Split(rawCmd, " ")
		for k, v := range cmd {
			cmd[k] = strings.TrimSpace(v)
		}

		if cmd[0] == "exit" {
			break
		}
		p, err := plugin.Open(fmt.Sprintf("./dist/%s.so", cmd[0]))
		if err != nil {
			fmt.Println(err)
			continue
		}

		f, err := p.Lookup("Exec")
		if err != nil {
			fmt.Println(err)
			continue
		}

		f.(func([]string) error)(cmd)
	}

}
