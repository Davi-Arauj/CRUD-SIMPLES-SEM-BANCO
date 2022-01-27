package main

import (
	"CRUD-Simples/entyties"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func rotaPrincipal(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Bem vindo")
}

func configurandoRotas(roteador *mux.Router) {

	roteador.HandleFunc("/", rotaPrincipal)
	roteador.HandleFunc("/clientes", entyties.ListarClientes).Methods("GET")
	roteador.HandleFunc("/clientes", entyties.CadastrarCliente).Methods("POST")
	roteador.HandleFunc("/clientes/{Id}", entyties.BuscarCliente).Methods("GET")
	roteador.HandleFunc("/clientes/{Id}", entyties.DeletarCliente).Methods("DELETE")
	roteador.HandleFunc("/clientes/{Id}", entyties.AtualizarCliente).Methods("PUT")
}

func configurandoServidor() {

	roteador := mux.NewRouter().StrictSlash(true)
	configurandoRotas(roteador)

	fmt.Println("O servidor está Rodando na porta 1337")

	log.Fatal(http.ListenAndServe(":1337", roteador))
}

func main() {
	configurandoServidor()

}
