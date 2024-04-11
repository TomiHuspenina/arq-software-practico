package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type Result struct {
	Title             string  `json:"title"`
	Price             float64 `json:"price"`
	AvailableQuantity int     `json:"avaiable_quantity"`
}

type Response struct {
	SiteID  string   `json:"site_id"`
	Results []Result `json:"results"` // por eso creo la strucr Result, por que no puede retornar un tipo de dato especifico
}

func main() {

	fmt.Println("Busca un producto: ")
	input := readInput()
	fmt.Println(input)

	results := search(input)
	for i, result := range results {
		fmt.Println(fmt.Sprintf(" [%d] %s", i, result.Title)) // en comillas dobles esta la variable donde voy a generar, %d es int, %s string
		//[0] Celular Samsung S23
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

// 1. Input de busqueda
// 2. Ejecutar llamada HTTP API mercado libre
// 3. Leer respuesta API mercado libre
// 4. Validar el status code de la response
// 5. Parsear y mostrar resultados por la terminal
