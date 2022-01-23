package main

import (
	"CRUD-Simples/entyties"
	"fmt"
	"log"
	"net/http"
	"strings"
)

func rotaPrincipal(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Bem vindo")
}

func rotearCliente(w http.ResponseWriter, r *http.Request) {

	partes := strings.Split(r.URL.Path, "/")

	if len(partes) == 2 || len(partes) == 3 && partes[2] == "" {
		if r.Method == "GET" {
			entyties.ListarClientes(w, r)
		} else if r.Method == "POST" {
			entyties.CadastrarLivros(w, r)
		}
	} else if len(partes) == 3 || len(partes) == 4 && partes[3] == "" {
		if r.Method == "GET" {
			entyties.BuscarCliente(w, r)
		} else if r.Method == "DELETE" {
			entyties.DeletarCliente(w, r)
		} else {
			w.WriteHeader(http.StatusNotFound)
		}

	} else {
		w.WriteHeader(http.StatusNotFound)
	}

}

func configurandoRotas() {
	http.HandleFunc("/", rotaPrincipal)
	http.HandleFunc("/clientes", rotearCliente)
	http.HandleFunc("/clientes/", rotearCliente)
}

func configurandoServidor() {
	configurandoRotas()

	fmt.Println("O servidor está Rodando na porta 1337")

	log.Fatal(http.ListenAndServe(":1337", nil))
}

func main() {
	configurandoServidor()

}
