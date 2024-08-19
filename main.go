package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/fishingboy/GoTest/commend"
	"os"
)

const VERSION = "1.0.0"

func main() {
	var showVersion bool
	var cmd string
	flag.BoolVar(&showVersion, "version", false, "Print version information.")
	flag.BoolVar(&showVersion, "v", false, "Print version information.")
	flag.StringVar(&cmd, "commend", "commend", "commend")
	flag.Parse()

	if showVersion {
		fmt.Println(VERSION)
		os.Exit(0)
	}

	switch cmd {
	case "commend":
		commend.Index(context.Background())
	}
}
