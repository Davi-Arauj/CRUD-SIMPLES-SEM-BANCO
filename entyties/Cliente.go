package entyties

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Cliente struct {
	Id    int    `json:"id"`
	Name  string `json:"Nome"`
	Email string `json:"e-mail"`
	Fone  int    `json:"fone"`
}

var Clientes = []Cliente{
	Cliente{
		1,
		"José",
		"j@j.com",
		8888,
	},
	Cliente{
		2,
		"Davi",
		"d@d.com",
		8888,
	},
}

func ListarClientes(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Clientes)
}

func CadastrarCliente(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	body, error := ioutil.ReadAll(r.Body)

	if error != nil {
		log.Fatal(error.Error())
	}

	var novoCliente Cliente
	json.Unmarshal(body, &novoCliente)
	novoCliente.Id = len(Clientes) + 1
	Clientes = append(Clientes, novoCliente)
	json.NewEncoder(w).Encode(novoCliente)
	w.WriteHeader(http.StatusCreated)
}

func BuscarCliente(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	vars := mux.Vars(r)

	id, error := strconv.Atoi(vars["Id"])

	if error != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	for _, cli := range Clientes {
		if cli.Id == id {
			json.NewEncoder(w).Encode(cli)
			return

		}
	}

	w.WriteHeader(http.StatusNotFound)

}

func DeletarCliente(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	vars := mux.Vars(r)

	id, error := strconv.Atoi(vars["Id"])

	if error != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	var indiceCliente int = -1

	for indice, cli := range Clientes {
		if cli.Id == id {
			indiceCliente = indice
			break
		}
	}

	if indiceCliente < 0 {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	ladoEsquerdo := Clientes[0:indiceCliente]
	ladoDireito := Clientes[indiceCliente+1:]
	Clientes = append(ladoEsquerdo, ladoDireito...)

}

func AtualizarCliente(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	vars := mux.Vars(r)

	id, error := strconv.Atoi(vars["Id"])

	if error != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	var indiceCliente int = -1

	for indice, cli := range Clientes {
		if cli.Id == id {
			indiceCliente = indice
			break
		}
	}

	if indiceCliente < 0 {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	body, error := ioutil.ReadAll(r.Body)
	if error != nil {
		fmt.Println(error.Error())
	}

	var alteracaoCliente Cliente
	json.Unmarshal(body, &alteracaoCliente)
	Clientes[indiceCliente] = alteracaoCliente
	json.NewEncoder(w).Encode(alteracaoCliente)
}
