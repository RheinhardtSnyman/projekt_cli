package main

import (
	"flag"
	"fmt"
	"time"
)

var (
	flagDuration = flag.Duration("d", time.Second, "waits d seconds between each arg output")
)

func main() {
	flag.Parse()
	args := flag.Args()
	for i, arg := range args {
		time.Sleep(*flagDuration)
		fmt.Println(i, " : ", arg)
	}
}
