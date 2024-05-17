package main

//try
//.\goconc.exe -t 3s  waiter -d 0.5s 1 2 3 4 5 6 :: waiter -d 2s a b c
// or Ctrl+C to cancel

import (
	"context"
	"flag"
	"log"
	"os"
	"os/exec"
	"os/signal"
	"strings"
	"sync"
)

var (
	flagTimeout = flag.Duration("t", 0, "set a timeout")
)

func main() {
	flag.Parse()
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	c := make(chan os.Signal, 1)
	go func() {
		signal.Notify(c, os.Interrupt)
		<-c
		log.Println("Cancel with Ctrl+C")
		cancel()
	}()

	if *flagTimeout != 0 {
		ctx, cancel = context.WithTimeout(ctx, *flagTimeout)
	}

	cmds := parseArgs(flag.Args())
	runCmd(ctx, cmds)
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

func runCmd(ctx context.Context, cmds []cmdArgs) {
	wg := &sync.WaitGroup{}
	for i, args := range cmds {
		wg.Add(1)
		go func(nr int, args cmdArgs) {
			log.Printf("Starting comd %d: %s %s\n",
				nr,
				args.name,
				strings.Join(args.args, " "),
			)
			cmd := exec.CommandContext(ctx, args.name, args.args...)
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			err := cmd.Run()
			log.Printf("cmd %d is ready\n", nr)
			if err != nil {
				log.Printf("cmd %d with error = %s\n", nr, err)
			}
			wg.Done()
		}(i+1, args)
	}
	wg.Wait()
}
