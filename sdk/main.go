package main

import (
	"flag"
	"fmt"
)

func main() {
	helpFlag := flag.Bool("help", false, "Returns configuration")
	// runFlag := flag.Bool("run", false, "Runs the asset")

	flag.Parse()

	if *helpFlag {
		fmt.Println("The asset help was requested")
	}
}
