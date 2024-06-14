package main

import (
	"fmt"
	"os"

	"github.com/Aries-Financial-inc/golang-dev-logic-challenge-raadhshenshahhaseeb/controllers"
)

func main() {
	if err := controllers.Start(); err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		os.Exit(1)
	}
}
