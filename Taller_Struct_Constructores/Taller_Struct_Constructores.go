package main

import "fmt"

type Producto struct {
	ID       int
	Nombre   string
	Precio   float64
	Cantidad int
}

func NuevoProducto(idProducto int, nombreProducto string, precioProducto float64, cantidadProducto int) Producto {
	return Producto{
		ID:       idProducto,
		Nombre:   nombreProducto,
		Precio:   precioProducto,
		Cantidad: cantidadProducto,
	}
}

func (producto *Producto) ActualizarPrecio(nuevoPrecio float64) {
	producto.Precio = nuevoPrecio
}

func (producto *Producto) ActualizarCantidad(nuevaCantidad int) {
	producto.Cantidad = nuevaCantidad
}

func (producto Producto) MostrarInfo() {
	fmt.Println("ID:", producto.ID, "- Nombre:", producto.Nombre, "- Precio:", producto.Precio, "- Cantidad:", producto.Cantidad)
}

type Inventario struct {
	Productos map[int]Producto
}

func NuevoInventario() Inventario {
	return Inventario{
		Productos: make(map[int]Producto),
	}
}

func (miInventario *Inventario) AgregarProducto(productoNuevo Producto) {
	miInventario.Productos[productoNuevo.ID] = productoNuevo
}

func (miInventario *Inventario) EliminarProducto(idABuscar int) {
	delete(miInventario.Productos, idABuscar)
}

func (miInventario Inventario) BuscarProducto(idABuscar int) (Producto, bool) {
	productoEncontrado, existe := miInventario.Productos[idABuscar]
	return productoEncontrado, existe
}

func (miInventario Inventario) ListarProductos() {
	for _, productoActual := range miInventario.Productos {
		productoActual.MostrarInfo()
	}
}

func main() {
	tienda := NuevoInventario()

	prod1 := NuevoProducto(1, "Cuaderno", 2.50, 20)
	prod2 := NuevoProducto(2, "Mochila", 35.00, 5)

	tienda.AgregarProducto(prod1)
	tienda.AgregarProducto(prod2)

	fmt.Println("Lista de productos:")
	tienda.ListarProductos()

	prod1.ActualizarPrecio(2.25)
	prod1.ActualizarCantidad(15)
	tienda.AgregarProducto(prod1)

	tienda.EliminarProducto(2)

	fmt.Println("Inventario modificado:")
	tienda.ListarProductos()
}