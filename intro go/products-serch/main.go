package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

type Result struct {
	Title             string  `json:"title"`
	Price             float64 `json:"price"`
	AvailableQuantity int     `json:"available_quantity"`
}

type Response struct {
	SiteID  string   `json:"site_id"`
	Results []Result `json:"results"` // por eso creo la strucr Result, por que no puede retornar un tipo de dato especifico
}

func main() {
	name := "Productos"
	CreateFile(name)
	fmt.Print("Busca un Producto: ")
	input := readInput()
	fmt.Println(input)

	file, err := os.OpenFile(name, os.O_APPEND|os.O_WRONLY, 0644) //os.O_APPEND indica que se debe agregar al final del archivo.
	if err != nil {                                               //os.O_WRONLY indica que el archivo se abrirá en modo de solo escritura.
		fmt.Println("Error abriendo archivo:", err) //argumento 0644 especifica los permisos del archivo.
		return
	}

	results := search(input)
	for i, result := range results {
		line := fmt.Sprintf(" [%d] %s \n", i, result.Title)
		fmt.Println(line)
		file.WriteString(line)
		time.Sleep(200 * time.Millisecond)
	}
}

func readInput() string {

	var input string
	_, err := fmt.Scan(&input) // se pasa & para cambiar el contenido
	if err != nil {            // si es distinto de nil significa que tiene un error
		fmt.Println("Error de reading input")
	}

	return input
}

func search(query string) []Result {

	url := fmt.Sprintf("http://api.mercadolibre.com/sites/MLA/search?q=%s", query)
	response, err := http.Get(url)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}
	bytes, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}
	var searchResponse Response
	err = json.Unmarshal(bytes, &searchResponse)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}

	return searchResponse.Results

}

func CreateFile(path string) error {
	file, err := os.Create(path)
	defer file.Close()
	if err != nil {
		return err
	}
	return nil
}

// 1. Input de busqueda
// 2. Ejecutar llamada HTTP API mercado libre
// 3. Leer respuesta API mercado libre
// 4. Validar el status code de la response
// 5. Parsear y mostrar resultados por la terminal
