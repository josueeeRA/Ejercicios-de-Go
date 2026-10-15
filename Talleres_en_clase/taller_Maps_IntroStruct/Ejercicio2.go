package main

import "fmt"

type Producto struct {
	nombre   string
	precio   float64
	cantidad int
}

func mostrarProducto(p Producto) {
	fmt.Println("Nombre:", p.nombre)
	fmt.Println("Precio:", p.precio)
	fmt.Println("Cantidad:", p.cantidad)
	fmt.Println("Valor total:", valorTotal(p))
	fmt.Println("Stock:", estadoStock(p))
	fmt.Println("")
}

func valorTotal(p Producto) float64 {
	return p.precio * float64(p.cantidad)
}

func estadoStock(p Producto) string {
	if p.cantidad > 0 {
		return "Disponible"
	}
	return "Sin stock"
}

func main() {
	producto1 := Producto{"Encebollado", 2.00, 5}
	producto2 := Producto{"Tigrillo", 2.50, 0}
	producto3 := Producto{"Bolon", 3.50, 3}

	mostrarProducto(producto1)
	mostrarProducto(producto2)
	mostrarProducto(producto3)
}