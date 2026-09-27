package contador

import (
	"bufio"
	"fmt"
	"os"
)

func ProgramaVoc() {
	var a, e, i, o, u int

	fmt.Print("Ingresa una frase: ")

	// 1. Preparamos el escáner para leer desde el teclado
	scanner := bufio.NewScanner(os.Stdin)

	// 2. Le decimos que capture todo hasta presionar Enter
	scanner.Scan()

	// 3. Guardamos lo capturado en la variable
	frase := scanner.Text()

	// Si venimos del menú principal, a veces queda un 'Enter' fantasma en la memoria.
	// Esta validación asegura que si la frase está vacía, lea de nuevo correctamente.
	if frase == "" {
		scanner.Scan()
		frase = scanner.Text()
	}

	// Evaluamos letra por letra incluyendo mayúsculas, minúsculas, tildes y diéresis
	for _, letra := range frase {
		if letra == 'a' || letra == 'A' || letra == 'á' || letra == 'Á' {
			a = a + 1
		} else if letra == 'e' || letra == 'E' || letra == 'é' || letra == 'É' {
			e = e + 1
		} else if letra == 'i' || letra == 'I' || letra == 'í' || letra == 'Í' {
			i = i + 1
		} else if letra == 'o' || letra == 'O' || letra == 'ó' || letra == 'Ó' {
			o = o + 1
		} else if letra == 'u' || letra == 'U' || letra == 'ú' || letra == 'Ú' || letra == 'ü' || letra == 'Ü' {
			u = u + 1
		}
	}

	fmt.Println("Conteo de vocales:")
	fmt.Println("A:", a)
	fmt.Println("E:", e)
	fmt.Println("I:", i)
	fmt.Println("O:", o)
	fmt.Println("U:", u)
}
