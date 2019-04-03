package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Person struct {
	ID        string   `json:"id,omitempty"`
	Firstname string   `json:"firstname,omitempty"`
	Lastname  string   `json:"lastname,omitempty"`
	Address   *Address `json:"address,omitempty"`
}

type Address struct {
	City  string `json:"city,omitempty"`
	State string `json:"state,omitempty"`
}

type PersonReturn struct {
	People  []Person `json:"people,omitempty"`
	Person  *Person  `json:"person,omitempty"`
	Message string   `json:"message,omitempty"`
}

var people []Person

func GetPeopleEndpoint(w http.ResponseWriter, req *http.Request) {
	json.NewEncoder(w).Encode(PersonReturn{
		People: people,
	})
}

func GetPersonEndpoint(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	for _, item := range people {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(PersonReturn{})
			return
		}
	}

	json.NewEncoder(w).Encode(&Person{})
}

func CreatePersonEndpoint(w http.ResponseWriter, req *http.Request) {
	var person Person
	var id = len(people) + 1
	_ = json.NewDecoder(req.Body).Decode(&person)
	person.ID = strconv.Itoa(id)
	people = append(people, person)
	json.NewEncoder(w).Encode(PersonReturn{
		Message: "New Person Added with Success!",
		People:  people,
	})
}

func DeletePersonEndpoint(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)

	for index, item := range people {
		if item.ID == params["id"] {
			people = append(people[:index], people[index+1:]...)
			break
		} else {
			json.NewEncoder(w).Encode(PersonReturn{
				Message: "This person does not exists",
			})
			return
		}
	}

	json.NewEncoder(w).Encode(PersonReturn{
		Message: "Person Deleted with Success!",
		People:  people,
	})
}

func main() {
	people = append(people, Person{
		ID:        "1",
		Firstname: "Luiz",
		Lastname:  "Fernando",
		Address: &Address{
			City:  "San Francisco",
			State: "California",
		},
	})

	people = append(people, Person{
		ID:        "2",
		Firstname: "Chris",
		Lastname:  "Whitman",
		Address: &Address{
			City:  "Los Angeles",
			State: "California",
		},
	})

	people = append(people, Person{
		ID:        "3",
		Firstname: "Kurt",
		Lastname:  "Braget",
		Address: &Address{
			City:  "Los Angeles",
			State: "California",
		},
	})

	router := mux.NewRouter()
	router.HandleFunc("/people", GetPeopleEndpoint).Methods("GET")
	router.HandleFunc("/person/{id}", GetPersonEndpoint).Methods("GET")
	router.HandleFunc("/createPerson/{id}", CreatePersonEndpoint).Methods("POST")
	router.HandleFunc("/deletePerson/{id}", DeletePersonEndpoint).Methods("DELETE")
	fmt.Println("Server is up and running at port 6000 ⚡️")
	log.Fatal(http.ListenAndServe(":1234", router))
}
