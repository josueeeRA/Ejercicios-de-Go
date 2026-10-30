package main

import "fmt"

type Producto struct {
	id       int
	nombre   string
	precio   float64
	cantidad int
}

func NuevoProducto(id int, nombre string, precio float64, cantidad int) Producto {
	return Producto{
		id:       id,
		nombre:   nombre,
		precio:   precio,
		cantidad: cantidad,
	}
}

func (p Producto) GetID() int {
	return p.id
}

func (p Producto) GetNombre() string {
	return p.nombre
}

func (p Producto) GetPrecio() float64 {
	return p.precio
}

func (p Producto) GetCantidad() int {
	return p.cantidad
}

func (p *Producto) SetPrecio(nuevoPrecio float64) {
	p.precio = nuevoPrecio
}

func (p *Producto) SetCantidad(nuevaCantidad int) {
	p.cantidad = nuevaCantidad
}

func (p Producto) MostrarInfo() {
	fmt.Println("ID:", p.GetID(), "| Nombre:", p.GetNombre(), "| Precio:", p.GetPrecio(), "| Cantidad:", p.GetCantidad())
}

type Inventario struct {
	productos map[int]Producto
}

func NuevoInventario() Inventario {
	return Inventario{
		productos: make(map[int]Producto),
	}
}

func (inv *Inventario) AgregarProducto(p Producto) {
	inv.productos[p.GetID()] = p
}

func (inv *Inventario) EliminarProducto(id int) {
	delete(inv.productos, id)
}

func (inv Inventario) BuscarProducto(id int) (Producto, bool) {
	p, existe := inv.productos[id]
	return p, existe
}

func (inv Inventario) ListarProductos() {
	for _, p := range inv.productos {
		p.MostrarInfo()
	}
}

func main() {

	tienda := NuevoInventario()

	p1 := NuevoProducto(1, "Teclado", 25.50, 10)
	p2 := NuevoProducto(2, "Mouse", 15.00, 20)

	p1.SetPrecio(22.00)
	p1.SetCantidad(15)

	tienda.AgregarProducto(p1)
	tienda.AgregarProducto(p2)

	fmt.Println("Inventario inicial:")
	tienda.ListarProductos()

	fmt.Println("\nBuscando producto con ID 1:")
	prodEncontrado, existe := tienda.BuscarProducto(1)
	if existe {
		prodEncontrado.MostrarInfo()
	} else {
		fmt.Println("Producto no encontrado")
	}

	tienda.EliminarProducto(2)
	fmt.Println("\nInventario después de eliminar el ID 2:")
	tienda.ListarProductos()
}