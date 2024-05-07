package main

import (
	"fmt"

	"github.com/MAS234/e-commerce-GO/ejercicios"
)

func main() {

	ejercicioResuelto, numero := ejercicios.Ejercicio01("50")
	fmt.Println(ejercicioResuelto)
	fmt.Println("numero", numero)
}
