// definicion del paquete
package main

//dependencias
import (
	"encoding/json"
	"fmt"
)

// estructuras
type Producto struct {
	Titulo      string  `json:"titulo"` //primero se pone el nombre de la variable y seguido el tipo de variable
	Descripcion string  `json:"descripcion"`
	Precio      float64 `json:"precio"`
	Disponible  bool    `json:"disponible"`
	//json no acepta mayusculas
}

// funciones
func main() {

	producto1 := Producto{
		Titulo:      "TV 55 Pulgadas",
		Descripcion: "TV samsung oferta",
		Precio:      155000,
		Disponible:  true,
	}

	productoConDescuento := aplicarDescuento(producto1)

	//muestra producto
	bytes, err := json.Marshal(producto1) // cera variable, ademas json.Marshal(producto1) donde producto1 es un parametro que es el objeto
	fmt.Println(string(bytes))
	fmt.Println(err)

	//muestra producto con descuento
	bytes, err = json.Marshal(productoConDescuento) //la variable ya esta creada
	fmt.Println(string(bytes))
	fmt.Println(err)
}

//Ejercicio: implementar funcion aplicarDescuento
//que recibe un producto y rebaja su precio 50%, mostrar el producto por la terminal

func aplicarDescuento(p Producto) Producto { //donde Producto sin parentesis es el tipo de retorno

	return Producto{
		Titulo:      p.Titulo,
		Descripcion: p.Descripcion,
		Precio:      p.Precio / 2,
		Disponible:  p.Disponible,
	}
}

//Usando punteros

/*
func aplicarDescuento(p *Producto) {

	p.Precio = p.Precio / 2

}

llamo al puntero --> aplicarDescuento(&producto1)
bytes, err := json.Marshal(producto1)
	fmt.Println(string(bytes))
	fmt.Println(err)
*/
