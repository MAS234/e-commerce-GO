package ejercicios

import (
	"strconv"
)

func Ejercicio01(texto string) (string, int) {

	numero, err := strconv.Atoi(texto)
	if err != nil {
		return "Error al parsear el número", 0
	}

	if numero >= 100 {
		respuesta := "El número es mayor o igual a 100"
		return respuesta, numero
	} else {
		respuesta := "El número es menor a 100"
		return respuesta, numero
	}

}
