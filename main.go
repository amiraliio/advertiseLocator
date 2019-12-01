package main

import (
	app "github.com/amiraliio/advertiselocator/providers"
	"os"
)

func main() {
	root, err := os.Getwd()
	if err != nil {
		panic("cannot get root path")
	}
	app.Start(root)
}
