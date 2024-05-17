package main

import (
	"flag"
	"log"
	"os"
	"os/exec"
	"strings"
	"sync"
)

func main() {
	flag.Parse()
	cmds := parseArgs(flag.Args())
	runCmd(cmds)
}

type cmdArgs struct {
	name string
	args []string
}

const cmdDelimiter = "::"

func parseArgs(args []string) []cmdArgs {
	if len(args) == 0 {
		return nil
	}
	var cmds []cmdArgs
	var cmd cmdArgs
	for _, arg := range args {
		switch {
		case arg == cmdDelimiter:
			cmds = append(cmds, cmd)
			cmd = cmdArgs{}
		case cmd.name == "":
			cmd.name = arg
		default:
			cmd.args = append(cmd.args, arg)
		}
	}
	cmds = append(cmds, cmd)
	return cmds
}

func runCmd(cmds []cmdArgs) {
	wg := &sync.WaitGroup{}
	for i, args := range cmds {
		wg.Add(1)
		go func(nr int, args cmdArgs) {
			log.Printf("Starting comd %d: %s %s\n",
				nr,
				args.name,
				strings.Join(args.args, " "),
			)
			cmd := exec.Command(args.name, args.args...)
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			err := cmd.Run()
			log.Printf("cmd %d is ready\n", nr)
			if err != nil {
				log.Printf("cmd %d with error = %s\n", nr, err)
			}
			wg.Done()
		}(i, args)
	}
	wg.Wait()
}
