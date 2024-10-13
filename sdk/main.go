package main

import (
	"flag"
	"fmt"
)

func main() {
	helpFlag := flag.Bool("help", false, "Returns configuration")
	runFlag := flag.Bool("run", false, "Runs the asset")
	defaultConfigFlag := flag.Bool("defaultConfig", false, "Returns default configuration")
	schemaFlag := flag.Bool("schema", false, "Returns the schema")

	flag.Parse()

	if *helpFlag {
		fmt.Println("The asset help was requested")
	}
	if *runFlag {
		fmt.Println("The asset run was requested")
	}
	if *defaultConfigFlag {
		fmt.Println("The default configuration was requested")
	}
	if *schemaFlag {
		fmt.Println("The schema was requested")
	}
}
