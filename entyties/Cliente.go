package entyties

import (
	"encoding/json"
	"net/http"
)

type Cliente struct {
	Id    int
	Name  string
	Email string
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

func  ListarClientes(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(Clientes)
}
