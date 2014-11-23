package main

import (
	"flag"
	"fmt"
	"os"
	"time"
)

func main() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "usage: %s <duration>\n", os.Args[0])
		flag.PrintDefaults()
		os.Exit(1)
	}
	flag.Parse()

	// display usage if no args provided
	if len(flag.Args()) < 1 {
		flag.Usage()
	}

	// parse & handle errors
	d, err := time.ParseDuration(flag.Arg(0))
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s: invalid duration %s\n", os.Args[0], flag.Arg(0))
		os.Exit(1)
	}

	// output duration in seconds
	fmt.Println(d.Seconds())
}
