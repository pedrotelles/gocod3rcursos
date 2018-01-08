package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func horaCerta(w http.ResponseWriter, r *http.Request) {
	s := time.Now().Format("02/01/2006 03:04:05")
	// 02 -> Dia
	// 01 -> Mês
	// 2006 -> Ano
	// 03 -> Hora
	// 04 -> Minuto
	// 05 -> Segundos
	fmt.Fprintf(w, "<h1>Hora Certa: %s</h1>", s)
}

func main() {
	http.HandleFunc("/horaCerta", horaCerta)
	log.Println("Executando...")
	log.Fatal(http.ListenAndServe(":3000", nil))
}
