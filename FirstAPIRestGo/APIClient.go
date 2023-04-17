package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Cliente é uma estrutura para representar um cliente
type Cliente struct {
	ID    int    `json:"id"`
	Nome  string `json:"nome"`
	Idade int    `json:"idade"`
}

// simulando uma base de dados de clientes
var clientes []Cliente

func main() {
	r := mux.NewRouter()

	// endpoint para obter todos os clientes
	r.HandleFunc("/clientes", obterClientes).Methods("GET")

	// endpoint para obter um cliente pelo ID
	r.HandleFunc("/clientes/{id}", obterCliente).Methods("GET")

	// endpoint para criar um novo cliente
	r.HandleFunc("/clientes", criarCliente).Methods("POST")

	// endpoint para atualizar um cliente existente
	r.HandleFunc("/clientes/{id}", atualizarCliente).Methods("PUT")

	// endpoint para excluir um cliente existente
	r.HandleFunc("/clientes/{id}", excluirCliente).Methods("DELETE")

	// inicia o servidor
	log.Fatal(http.ListenAndServe(":8000", r))
}

// obterClientes retorna todos os clientes
func obterClientes(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(clientes)
}

// obterCliente retorna um cliente pelo ID
func obterCliente(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	for _, cliente := range clientes {
		if cliente.ID == id {
			json.NewEncoder(w).Encode(cliente)
			return
		}
	}
	json.NewEncoder(w).Encode(&Cliente{})
}

// criarCliente cria um novo cliente
func criarCliente(w http.ResponseWriter, r *http.Request) {
	var cliente Cliente
	_ = json.NewDecoder(r.Body).Decode(&cliente)
	cliente.ID = len(clientes) + 1
	clientes = append(clientes, cliente)
	json.NewEncoder(w).Encode(cliente)
}

// atualizarCliente atualiza um cliente existente
func atualizarCliente(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	for index, cliente := range clientes {
		if cliente.ID == id {
			var novoCliente Cliente
			_ = json.NewDecoder(r.Body).Decode(&novoCliente)
			clientes[index] = Cliente{ID: id, Nome: novoCliente.Nome, Idade: novoCliente.Idade}
			json.NewEncoder(w).Encode(clientes[index])
			return
		}
	}
	json.NewEncoder(w).Encode(clientes)
}

// excluirCliente exclui um cliente existente
func excluirCliente(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	for index, cliente := range clientes {
		if cliente.ID == id {
			clientes = append(clientes[:index], clientes[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(clientes)
}
