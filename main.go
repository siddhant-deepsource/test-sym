package main

import (
	"fmt"

	gr "github.com/hashicorp/go-rootcerts"
)

func main() {
	fmt.Println(" hello")
	// TODO : hello
	_, _ = gr.LoadCAPath("demo")
}
