package main

import (
	"fmt"
	"io"
	"lib/fibonacci"
	"log"
	"net/http"
	"strconv"

)

func main() {

	http.HandleFunc("/hola", func(w http.ResponseWriter, peticion *http.Request) {
		n := fibonacci.Fibonacci(4)
		io.WriteString(w, strconv.Itoa(n))
	})
	direccion := ":8080"
	fmt.Println("Servidor listo escuchando en " + direccion)
	log.Fatal(http.ListenAndServe(direccion, nil))
}
