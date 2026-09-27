package conversor

import "fmt"

func ProgramaConv() {
	var dolares float64
	var opcion int

	fmt.Print("Ingresa el valor en dólares: $")
	fmt.Scan(&dolares)

	fmt.Println("Selecciona la moneda:")
	fmt.Println("1. Euros")
	fmt.Println("2. Libras Esterlinas")
	fmt.Println("3. Won Surcoreano")
	fmt.Println("4. BTC")
	fmt.Print("Opción: ")
	fmt.Scan(&opcion)

	if opcion == 1 {
		euros := dolares * 0.95
		fmt.Println("Equivale a:", euros, "Euros")
	} else if opcion == 2 {
		libras := dolares * 0.79
		fmt.Println("Equivale a:", libras, "Libras")
	} else if opcion == 3 {
		won := dolares * 1340.50
		fmt.Println("Equivale a:", won, "Won")
	} else if opcion == 4 {
		btc := dolares * 0.000015
		fmt.Println("Equivale a:", btc, "BTC")
	} else {
		fmt.Println("Opción incorrecta")
	}
}
