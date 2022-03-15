package main

import (
	"fmt"
	"github.com/eolinker/apinto-ingress-controller/cmd"
	"os"
)

func main() {
	root := cmd.NewIngressControllerCommand()
	if err := root.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
}
