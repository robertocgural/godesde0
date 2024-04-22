package estructuras //Para que funcione hay que cambiar a main

import "fmt"

func mostrarAltura(altura int) string {
	var resultado string

	if altura >= 185 {
		resultado = "Eres una persona alta"

	} else {
		resultado = "Eres una persona bajita"

	}

	return resultado
}

func estructuras() { //Para que funcione hay que cambiar a main
	// Condiciones
	var miAltura int
	fmt.Println("¿Cuanto mides (centimetros)?")
	fmt.Scan(&miAltura)

	fmt.Println(mostrarAltura(miAltura))
	fmt.Println("Mides", miAltura, "cm")

}
