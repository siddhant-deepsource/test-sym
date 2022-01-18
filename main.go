package main

import (
	"fmt"

	gr "github.com/hashicorp/go-rootcerts"
)

func main() {
	fmt.Println("hello")
	_, _ = gr.LoadCAPath("demo")
	fmt.Println("test")
}
