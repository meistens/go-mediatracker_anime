package main

import (
	"fmt"

	"gioui.org/app"
)

func main() {
	// test installed deps and Gio
	fmt.Println("Deps installs correct!")
	fmt.Printf("Gio app package: %T\n", &app.Window{})
}
