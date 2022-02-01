package entyties

import (
	"CRUD-Simples/db"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

type Cliente struct {
	gorm.Model
	Name  string `json:"Nome"`
	Email string `json:"e-mail"`
	Fone  int    `json:"fone"`
}

func ListarClientes(w http.ResponseWriter, r *http.Request) {
	var clientes []Cliente
	db.DB.Find(&clientes)

	json.NewEncoder(w).Encode(&clientes)
}

func CadastrarCliente(w http.ResponseWriter, r *http.Request) {
	var novoCliente Cliente
	json.NewDecoder(r.Body).Decode(&novoCliente)

	createdCliente := db.DB.Create(&novoCliente)

	err := createdCliente.Error

	if err != nil {
		fmt.Println(err)
	}

	json.NewEncoder(w).Encode(&createdCliente)

}

func BuscarCliente(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	var cliente Cliente
	db.DB.First(&cliente, params["Id"])

	json.NewEncoder(w).Encode(&cliente)

}

func DeletarCliente(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	var cliente Cliente

	db.DB.First(&cliente, vars["Id"])
	db.DB.Delete(&cliente)

	json.NewEncoder(w).Encode(&cliente)

}

func AtualizarCliente(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var clienteUpdate Cliente
	json.NewDecoder(r.Body).Decode(&clienteUpdate)
	var cliente Cliente
	db.DB.First(&cliente, vars["Id"])
	db.DB.Model(&cliente).Update(&clienteUpdate)

	json.NewEncoder(w).Encode(&cliente)
}

func JsonMiddleWare(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")

		next.ServeHTTP(w, r)
	})
}
