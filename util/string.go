package util

import (
	"fmt"
	"strings"
)

func stringTest() {
	controllerName := "AdminController"
	name := controllerName[0 : len(controllerName)-10]
	fmt.Println(name) //Admin
	lowerName := strings.ToLower(name)
	fmt.Println(lowerName) //admin
}

func initstring() {
	stringTest()
}
