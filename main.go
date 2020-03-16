package main

import (
	"fmt"
	"github.com/travelist/aoj-cli/cmd"
	"os"
)

func main() {
	if e := cmd.Run(); e != nil {
		fmt.Printf("%+v\n", e)
		os.Exit(1)
	}
}
