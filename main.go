package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"plugin"
	"strings"
)

func src(cmd string) string {
	return fmt.Sprintf("./plugins/%s.go", cmd)
}

func dist(cmd string) string {
	return fmt.Sprintf("./dist/%s.so", cmd)
}

// build command, you need to make sure that src file exist, dictionary dist exist
func build(src, dist string) error {
	cmd := exec.Command("go", "build", "-buildmode=plugin", "-o", dist, src)
	return cmd.Run()
}

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

		// buildin commands
		if cmd[0] == "exit" {
			break
		}

		if cmd[0] == "build" {
			if len(cmd) < 2 {
				fmt.Println("need command name")
				continue
			}
			src := src(cmd[1])
			dist := dist(cmd[1])
			if _, err := os.Stat(src); err != nil {
				fmt.Println("plugin not found")
				continue
			}
			if err := build(src, dist); err != nil {
				fmt.Println(err)
				continue
			}
			fmt.Println("build success")
			continue
		}

		// check if the plugin is compiled
		if _, err := os.Stat(dist(cmd[0])); os.IsNotExist(err) {
			// check if src exists
			if _, err := os.Stat(src(cmd[0])); os.IsNotExist(err) {
				fmt.Println("plugin not found")
				continue
			}else{
				fmt.Printf("found source code but not compiled, use `build %s` to compile\n", cmd[0])
			}
			continue
		}

		p, err := plugin.Open(dist(cmd[0]))
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
