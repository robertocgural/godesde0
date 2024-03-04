package main

import (
	"fmt"

	"github.com/robertocgural/godesde0/variables"
)

func main() {
	//variables.MostroEnteros()
	//primerospasos.PracticasEnGo()
	//variables.RestoVariables()
	estado, texto := variables.ConviertoaTexto(1588)
	fmt.Println(estado)
	fmt.Println(texto)
}
