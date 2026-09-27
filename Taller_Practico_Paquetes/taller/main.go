package main

import (
	"fmt"
	"taller/contador"
	"taller/conversor"
)

func main() {
	var eleccion int

	fmt.Println("--- TALLER ---")
	fmt.Println("1. Conversor de Monedas")
	fmt.Println("2. Contador de Vocales")
	fmt.Print("Elige (1 o 2): ")

	fmt.Scan(&eleccion)

	if eleccion == 1 {
		conversor.ProgramaConv()
	} else if eleccion == 2 {
		contador.ProgramaVoc()
	} else {
		fmt.Println("Opción no válida")
	}
}
