package main //Para que funcione hay que cambiar a main

import "fmt"

func main() { //Para que funcione hay que cambiar a main
	// Arrais y listas
	var personas = [3]string{"Pablo", "Pedro", "Juan"}

	//fmt.Println(personas[2])

	// Bucles for
	for contador := 0; contador < len(personas); contador++ {
		// Bloque de codigo
		fmt.Println("-", personas[contador])
	}
}
