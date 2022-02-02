package entyties

import (
	"CRUD-Simples/db"
	"encoding/json"
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

//Service do Cliente
func ListarClientes(w http.ResponseWriter, r *http.Request) {
	var clientes []Cliente

	if result := db.DB.Find(&clientes); result.Error != nil {
		w.WriteHeader(400)
		panic(result.Error)
	} else if result.RowsAffected == 0 {
		w.WriteHeader(http.StatusNotFound)
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&clientes)
}

func CadastrarCliente(w http.ResponseWriter, r *http.Request) {
	var novoCliente Cliente
	json.NewDecoder(r.Body).Decode(&novoCliente)

	createdCliente := db.DB.Create(&novoCliente)

	if createdCliente.Error != nil {
		w.WriteHeader(400)
		panic(createdCliente.Error)
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(&createdCliente)

}

func BuscarCliente(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	var cliente Cliente
	if err := db.DB.First(&cliente, params["Id"]); err.Error != nil {
		w.WriteHeader(400)
	}

	json.NewEncoder(w).Encode(&cliente)

}

func DeletarCliente(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	var cliente Cliente

	if err := db.DB.First(&cliente, vars["Id"]); err.Error != nil {
		w.WriteHeader(400)
		panic(err.Error)
	}
	db.DB.Delete(&cliente)

	json.NewEncoder(w).Encode(&cliente)

}

func AtualizarCliente(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var clienteUpdate Cliente
	json.NewDecoder(r.Body).Decode(&clienteUpdate)
	var cliente Cliente
	if err := db.DB.First(&cliente, vars["Id"]); err.Error != nil {
		w.WriteHeader(400)
		panic(err.Error)
	}
	db.DB.Model(&cliente).Update(&clienteUpdate)

	json.NewEncoder(w).Encode(&cliente)
}

func JsonMiddleWare(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")

		next.ServeHTTP(w, r)
	})
}
