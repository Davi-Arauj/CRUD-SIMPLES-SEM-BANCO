package main

import (
	"CRUD-Simples/entyties"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

func TestBuscarCliente(t *testing.T) {
	tt := []struct {
		rVar     entyties.Cliente
		expected entyties.Cliente
	}{
		{
			rVar: entyties.Cliente{
				Name:  "aaa",
				Email: "v@v",
				Fone:  9999,
			},
			expected: entyties.Cliente{
				Name:  "aaa",
				Email: "v@v",
				Fone:  9999,
			},
		}, {
			rVar: entyties.Cliente{
				Name:  "bbb",
				Email: "v@v",
				Fone:  9999,
			},
			expected: entyties.Cliente{
				Name:  "bbb",
				Email: "v@v",
				Fone:  9999,
			}},
	}

	req, err := http.NewRequest("GET", "localhost:1339/clientes", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc("/clientes", entyties.ListarClientes).Methods("GET")
	router.ServeHTTP(rr, req)

	assert.Equal(t, 200, rr.Code)

	corpo, err := ioutil.ReadAll(rr.Body)
	if err != nil {
		t.Errorf("Não foi possivel analisar o corpo da resposta: %v\n%v", rr.Body.String(), err)
		t.FailNow()
	}

	dadosRes := make(map[string]interface{})
	if err = json.Unmarshal(corpo, &dadosRes); err != nil {
		t.Errorf("Não foi possivel ler o corpo da resposta: %v\n%v", string(corpo), err)
		t.FailNow()
	}
	for chaveEntrada, valorEntrada := range tt {
		if valorSaida, ok := dadosRes[chaveEntrada]; ok {
			if fmt.Sprintf("%v", valorEntrada) != fmt.Sprintf("%v", valorSaida) {
				t.Errorf("esperava %v, obteve %v", valorEntrada, valorSaida)
				t.FailNow()
			}
		}
	}
}

func TestAdicionarCliente(t *testing.T) {

	// {
	// 	rVar: entyties.Cliente{
	// 		Name:  "Davi",
	// 		Email: "davi@email",
	// 		Fone:  88995588,
	// 	},
	// 	expectd: "Nome: Davi - Email: davi@email - Fone: 88995588",
	// }, {
	// 	rVar: entyties.Cliente{
	// 		Name:  "Augusto",
	// 		Email: "augusto@email",
	// 		Fone:  88995588,
	// 	},
	// 	expectd: "Nome: Augusto - Email: augusto@email - Fone: 88995588",
	// },
}
