package main

import (
	"fmt" // Nos trae el ambiente de la pc actual.

	"github.com/robertocgural/godesde0/ejercicios"
)

func main() {
	/* variables.RestoVariables()
	primerospasos.PracticasEnGo()
	variables.RestoVariables()

	estado, texto := variables.ConviertoaTexto(1588)
	fmt.Println(estado)
	fmt.Println(texto)

	if os := runtime.GOOS; os == "linux" || os == "OS X" {
		fmt.Println("Esto no es Windows es:", os)
	} else {
		fmt.Println("Esto es Windows")
	}

	switch os := runtime.GOOS; os {
	case "windows":
		fmt.Printf("Es Windows")
	case "darwin":
		fmt.Printf("Es Darwin")
	default:
		fmt.Printf("%s \n", os)
	} */

	numero, texto := ejercicios.ConvNumerico("7")
	fmt.Println(numero)
	fmt.Println(texto)
}
