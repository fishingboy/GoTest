package main

import (
	"flag"
	"github.com/fishingboy/GoTest/commend"
)

func main() {
	var showVersion bool
	var cmd string
	flag.BoolVar(&showVersion, "version", false, "Print version information.")
	flag.BoolVar(&showVersion, "v", false, "Print version information.")
	flag.StringVar(&cmd, "commend", "commend", "commend")
	flag.Parse()

	if showVersion {
		cmd = "version"
	}

	switch cmd {
	case "version":
		commend.Version()
	case "grep":
		commend.Grep()
	case "scope":
		commend.Scope()
	}
}
