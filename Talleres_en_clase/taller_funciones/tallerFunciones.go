package main

import (
	"fmt"
)

func promedio(notas []int) float64 {
	var suma int = 0
	for i := 0; i < len(notas); i++ {
		suma = suma + notas[i]
	}
	return float64(suma) / float64(len(notas))
}

func opcion1() {
	var cantidad int
	fmt.Print("Ingresa la cantidad de estudiantes en el curso: ")
	fmt.Scan(&cantidad)

	if cantidad <= 0 {
		fmt.Println("La cantidad debe ser mayor a 0.")
		return
	}

	var notas []int
	var nota int

	for i := 1; i <= cantidad; i++ {
		fmt.Print("Ingresa la nota del estudiante ", i, " (0 a 100): ")
		fmt.Scan(&nota)
		notas = append(notas, nota)
	}

	promedio := promedio(notas)
	fmt.Println("El promedio del curso es:", promedio)

	if promedio >= 70 {
		fmt.Println("Estado del curso: aprobado")
	} else {
		fmt.Println("Estado del curso: reprobado")
	}

	switch {
	case promedio >= 90 && promedio <= 100:
		fmt.Println("Rendimiento: Excellent performance")
	case promedio >= 80 && promedio < 90:
		fmt.Println("Rendimiento: Good performance")
	case promedio >= 70 && promedio < 80:
		fmt.Println("Rendimiento: Satisfactory performance")
	case promedio < 70:
		fmt.Println("Rendimiento: Needs improvement")
	}
}

func opcion2() {
	var n int
	fmt.Print("Ingresa un número n: ")
	fmt.Scan(&n)

	var suma int = 0
	for i := 1; i <= n; i++ {
		suma = suma + i
	}
	fmt.Println("La suma de los números del 1 al", n, "es:", suma)
}

func opcion3() {
	var celsius float64
	fmt.Print("Ingresa la temperatura en grados Celsius: ")
	fmt.Scan(&celsius)

	fahrenheit := (celsius * 9 / 5) + 32
	fmt.Println(celsius, "grados Celsius equivalen a", fahrenheit, "grados Fahrenheit")
}

func opcion4() {
	var fahrenheit float64
	fmt.Print("Ingresa la temperatura en grados Fahrenheit: ")
	fmt.Scan(&fahrenheit)

	celsius := (fahrenheit - 32) * 5 / 9
	fmt.Println(fahrenheit, "grados Fahrenheit equivalen a", celsius, "grados Celsius")
}

func main() {
	var opcion string

	for {
		fmt.Println("--- MENÚ PRINCIPAL ---")
		fmt.Println("1. Analizar notas de estudiantes")
		fmt.Println("2. Sumar números del 1 al n")
		fmt.Println("3. Convertir Celsius a Fahrenheit")
		fmt.Println("4. Convertir Fahrenheit a Celsius")
		fmt.Println("0. Salir (o escribe 'salir')")
		fmt.Print("Elige una opción: ")
		fmt.Scan(&opcion)

		if opcion == "0" || opcion == "salir" {
			fmt.Println("Saliendo del programa...")
			break
		}

		switch opcion {
		case "1":
			opcion1()
		case "2":
			opcion2()
		case "3":
			opcion3()
		case "4":
			opcion4()
		default:
			fmt.Println("Opción no válida. Inténtalo de nuevo.")
		}
	}
}
