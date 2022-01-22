package entyties

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
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
