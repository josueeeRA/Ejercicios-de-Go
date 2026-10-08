package main

import "fmt"

func main() {
	notas := [6][4]float64{
		{8.2, 9.4, 7.3, 8.8},
		{6.1, 7.7, 6.9, 5.3},
		{8.8, 10.0, 9.9, 76.},
		{7.9, 8.5, 7.2, 8.3},
		{5.0, 6.6, 5.5, 8.5},
		{8.0, 9.5, 7.0, 8.8},
	}

	var sumaTotal float64
	totalNotas := 0

	for i := 0; i < len(notas); i++ {
		notasEstudiante := notas[i][:]

		sumaEstudiante := 0.0
		notaAlta := notasEstudiante[0]
		notaBaja := notasEstudiante[0]

		for j := 0; j < len(notasEstudiante); j++ {
			nota := notasEstudiante[j]
			sumaEstudiante += nota
			sumaTotal += nota
			totalNotas++

			if nota > notaAlta {
				notaAlta = nota
			}
			if nota < notaBaja {
				notaBaja = nota
			}
		}

		promedioEstudiante := sumaEstudiante / float64(len(notasEstudiante))

		fmt.Println("Estudiante", i+1)
		fmt.Println("Promedio:", promedioEstudiante)
		fmt.Println("Nota mas alta:", notaAlta)
		fmt.Println("Nota mas baja:", notaBaja)
	}

	promedioGeneral := sumaTotal / float64(totalNotas)
	fmt.Println("Promedio general de la clase:", promedioGeneral)
}