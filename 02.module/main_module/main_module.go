package main

import (
	"fmt"

	"example.com/module_one"
)

func main() {
	message := module_one.Hello("SorrisoRonaldo")
	fmt.Println(message)
}
