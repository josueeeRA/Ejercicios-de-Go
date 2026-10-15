package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Print("Ingrese un texto: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	texto := scanner.Text()

	palabras := strings.Fields(texto)

	contador := make(map[string]int)
	for i := 0; i < len(palabras); i++ {
		palabra := palabras[i]
		contador[palabra]++
	}

	palabraMasRepetida := ""
	mayorCantidad := 0

	fmt.Println("Conteo de palabras:")
	for palabra, cantidad := range contador {
		fmt.Println(palabra, ":", cantidad)

		if cantidad > mayorCantidad {
			mayorCantidad = cantidad
			palabraMasRepetida = palabra
		}
	}

	if mayorCantidad > 0 {
		fmt.Println("La palabra mas repetida fue:", palabraMasRepetida)
	} else {
		fmt.Println("No se ingresaron palabras.")
	}
}