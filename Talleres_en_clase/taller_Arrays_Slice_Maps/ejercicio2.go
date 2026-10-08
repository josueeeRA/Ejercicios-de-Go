package main

import "fmt"

// 5. Funcion para determinar la actividad ganadora
func obtenerGanador(votos map[string]int) string {
	actividadGanadora := ""
	mayorVotos := -1

	for actividad, cantidad := range votos {
		if cantidad > mayorVotos {
			mayorVotos = cantidad
			actividadGanadora = actividad
		}
	}
	return actividadGanadora
}

func main() {
	// 1. Crear el map con las actividades inicializadas en 0
	votos := map[string]int{
		"deportes":    0,
		"videojuegos": 0,
		"cine":        0,
		"musica":      0,
	}

	// 2 y 3. Ingreso de 5 votos por teclado y actualizacion
	fmt.Println("Opciones: deportes, videojuegos, cine, musica")
	for i := 1; i <= 5; i++ {
		var eleccion string
		fmt.Println("Ingrese el voto", i, ":")
		fmt.Scanln(&eleccion)

		// Verificamos si existe en el map
		_, existe := votos[eleccion]
		if existe {
			votos[eleccion]++
		} else {
			fmt.Println("Actividad no valida. Intente de nuevo.")
			i-- // Repetir el intento para completar los 5 votos validos
		}
	}

	// 4. Mostrar los resultados recorriendo el map
	fmt.Println("Resultados de la votacion:")
	for actividad, cantidad := range votos {
		fmt.Println(actividad, ":", cantidad, "votos")
	}

	// Determinar el ganador usando la funcion
	ganador := obtenerGanador(votos)
	fmt.Println("La actividad con mayor aceptacion es:", ganador)
}