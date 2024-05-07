package main

import (
	"fmt"
	"runtime"

	"github.com/MAS234/e-commerce-GO/variables"
)

func main() {
	estado, texto := variables.ConviertoTexto(10)
	fmt.Println(estado, texto)

	os := runtime.GOOS

	if os == "windows" {
		fmt.Println("Es un equipo Windows, confirmado")
	}
}
