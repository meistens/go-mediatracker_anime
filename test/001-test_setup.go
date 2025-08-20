package main

import (
	"fmt"

	"gioui.org/app"
	"github.com/sirupsen/logrus"
)

func main() {
	// test installed deps and Gio
	fmt.Println("Deps installs correct!")
	fmt.Printf("Gio app package: %T\n", &app.Window{})

	logger := logrus.New()
	logger.Info("logrus works...")

	fmt.Println("all works as it should")
}
