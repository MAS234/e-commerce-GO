package main

import (
	"fmt"

	"github.com/MAS234/e-commerce-GO/variables"
)

func main() {
	estado, texto := variables.ConviertoTexto(10)
	fmt.Println(estado, texto)
}
