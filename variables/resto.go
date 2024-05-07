package variables

import (
	"fmt"
	"strconv"
	"time"
)

var Nombre string //esta variable puede ser importada en cualquier otro archivo
var Estado bool
var Sueldo float32
var Fecha time.Time

func RestoVariables() {

	Nombre = "Juan"
	Estado = true
	Sueldo = 10.5
	Tiempo := time.Now()

	fmt.Println(Nombre)
	fmt.Println(Estado)
	fmt.Println(Sueldo)
	fmt.Println(Tiempo)
}

func ConviertoTexto(numero int) (bool, string) {
	texto := strconv.Itoa(numero)
	return true, texto
}
