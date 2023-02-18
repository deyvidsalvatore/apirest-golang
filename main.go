package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Pessoa struct {
	ID        string    `json: "id,omitempty"`
	Nome      string    `json:"nome,omitempty"`
	Sobrenome string    `json:"sobrenome,omitempty"`
	Endereco  *Endereco `json:"endereco,omitempty"`
}

type Endereco struct {
	Cidade string `json:"cidade,omitempty"`
	Estado string `json:"estado,omitempty"`
}

var pessoas []Pessoa

func getPessoas(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(pessoas)
}

func getPessoa(w http.ResponseWriter, r *http.Request){
	params := mux.Vars(r)
	for _, item := range pessoas {
		if item.ID == params["id"]{
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Pessoa{})
}

func createPessoa(w http.ResponseWriter, r *http.Request)  {
	params := mux.Vars(r)
	var pessoa Pessoa
	_ = json.NewDecoder(r.Body).Decode(&pessoa)
	pessoa.ID = params["id"]
	pessoas = append(pessoas, pessoa)
	json.NewEncoder(w).Encode(pessoas)
}

func deletePessoa(w http.ResponseWriter, r *http.Request){
	params := mux.Vars(r)
	
	for index, item := range pessoas {
		if item.ID == params["id"]{
			pessoas = append(pessoas[:index], pessoas[index+1:]...)
			break
		}
		json.NewEncoder(w).Encode(pessoas)
	}
}

func main() {
	router := mux.NewRouter()
	pessoas = append(pessoas, Pessoa{ID: "1", Nome: "John", Sobrenome: 
	"Doe", Endereco: &Endereco{Cidade: "City X", Estado: "State X"}})
	pessoas = append(pessoas, Pessoa{ID: "2", Nome: "Deyvid", Sobrenome: 
	"Santos", Endereco: &Endereco{Cidade: "Recife", Estado: "Pernambuco"}})
	router.HandleFunc("/contato", getPessoas).Methods("GET")
	router.HandleFunc("/contato/{id}", getPessoa).Methods("GET")
	router.HandleFunc("/contato/{id}", createPessoa).Methods("POST")
	router.HandleFunc("/contato/{id}", deletePessoa).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8000", router))
}	