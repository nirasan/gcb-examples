package main

import (
	"fmt"
	"gopkg.in/alecthomas/kingpin.v2"
)
var (
	verbose = kingpin.Flag("verbose", "Verbose mode.").Short('v').Bool()
	name    = kingpin.Arg("name", "Name of user.").Required().String()
)

func main() {
	kingpin.Parse()
	fmt.Printf(Sprint(*verbose, *name))
}

func Sprint(v bool, n string) string {
	return fmt.Sprintf("%v, %s\n", v, n)
}
