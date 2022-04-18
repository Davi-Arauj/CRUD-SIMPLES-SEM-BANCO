package main

import (
	"CRUD-Simples/db"
	"CRUD-Simples/entyties"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func rotaPrincipal(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Bem vindo ao MUX")
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
	roteador.Use(entyties.JsonMiddleWare)
	configurandoRotas(roteador)
	fmt.Println("O servidor está Rodando na porta 1339")
	log.Fatal(http.ListenAndServe(":1339", roteador))
}

func main() {
	db.ConnectDB()

	//Criando tabela Cliente de acordo com a struct
	db.DB.AutoMigrate(&entyties.Cliente{})

	//Iniciando o Servidor
	configurandoServidor()

	// Close the databse connection when the main function closes
	defer db.DB.Close()
}
