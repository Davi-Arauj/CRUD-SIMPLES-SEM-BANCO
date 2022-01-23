package entyties

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"strings"
)

type Cliente struct {
	Id    int    `json:"id"`
	Name  string `json:"Nome"`
	Email string `json:"e-mail"`
}

var Clientes = []Cliente{
	Cliente{
		1,
		"José",
		"j@j.com",
	},
	Cliente{
		2,
		"Davi",
		"d@d.com",
	},
}

func ListarClientes(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(Clientes)
}

func CadastrarLivros(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	body, error := ioutil.ReadAll(r.Body)

	if error != nil {
		fmt.Println(error.Error())
	}

	var novoCliente Cliente
	json.Unmarshal(body, &novoCliente)
	novoCliente.Id = len(Clientes) + 1
	Clientes = append(Clientes, novoCliente)
	json.NewEncoder(w).Encode(novoCliente)

}

func BuscarCliente(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	partes := strings.Split(r.URL.Path, "/")

	id, error := strconv.Atoi(partes[2])

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
	partes := strings.Split(r.URL.Path, "/")

	id, error := strconv.Atoi(partes[2])

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
