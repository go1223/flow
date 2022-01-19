package main

import (
	"fmt"
	"os"
)

func main() {
	var port = "8890"
	if len(os.Args) > 2 {
		port = os.Args[1]
	}
	fmt.Println("single flow server")
	RunServer(port)
}
