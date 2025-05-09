package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
	"time"
)

func guppa(times int, names ...string) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	total := len(names)

	abs := func(x int) int {
		if x < 0 {
			return -x
		}
		return x
	}

	fmt.Printf("Good party, ")

	for i := 0; i < times; i++ {
		goo := make([]string, 0, total)
		par := make([]string, 0, total)

		for _, name := range names {
			if r.Intn(2) == 0 {
				goo = append(goo, name)
			} else {
				par = append(par, name)
			}
		}

		fmt.Printf("work out, let mash up!\n")
		time.Sleep(time.Millisecond * 500)

		fmt.Printf("✊ %d\n  %v\n", len(goo), goo)
		fmt.Printf("✋ %d\n  %v\n\n", len(par), par)

		if abs(len(goo)-len(par)) <= 1 {
			break
		}

		time.Sleep(time.Millisecond * 500)
	}
}

var max int

func usage() {
	fmt.Fprintf(flag.CommandLine.Output(), "Usage: goopar [opts...] <args...>\n")
	flag.PrintDefaults()
	os.Exit(0)
}

func init() {
	flag.IntVar(&max, "m", 10, "max times")
	flag.Usage = usage
}

func main() {
	flag.Parse()

	if len(flag.Args()) < 1 {
		fmt.Fprintf(flag.CommandLine.Output(), "least one arg is required")
		flag.Usage()
		os.Exit(1)
	}

	guppa(max, flag.Args()...)
}
