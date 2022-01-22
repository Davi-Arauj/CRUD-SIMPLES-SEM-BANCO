package main

import (
	"CRUD-Simples/entyties"
	"fmt"
	"log"
	"net/http"
)

func rotaPrincipal(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Bem vindo")
}

func rotearCliente(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		entyties.ListarClientes(w, r)
	} else if r.Method == "POST" {
		entyties.CadastrarLivros(w, r)
	}
}

func configurandoRotas() {
	http.HandleFunc("/", rotaPrincipal)
	http.HandleFunc("/clientes", rotearCliente)
	//http.HandleFunc("/clientes", entyties.CadastrarLivros)
}

func configurandoServidor() {
	configurandoRotas()

	fmt.Println("O servidor está Rodando na porta 1337")

	log.Fatal(http.ListenAndServe(":1337", nil))
}

func main() {
	configurandoServidor()

}
