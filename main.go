package main

import (
	"flag"
	"fmt"
	"os"
)

const VERSION = "1.0.0"

func main() {
	var showVersion bool
	var commend string
	flag.BoolVar(&showVersion, "version", false, "Print version information.")
	flag.BoolVar(&showVersion, "v", false, "Print version information.")
	flag.StringVar(&commend, "commend", "grep", "commend")
	flag.Parse()
	if showVersion {
		fmt.Println(VERSION)
		os.Exit(0)
	}
}
