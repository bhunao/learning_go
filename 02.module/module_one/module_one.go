package module_one

import "fmt"

func Hello(name string) string {
	message:= fmt.Sprintf("Salve, %v. bem vindo", name)
	return message
}
