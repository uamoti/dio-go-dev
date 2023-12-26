package main

/* Create as REST API to handle customer data - CRUD
On this version, I want to use only the standard packages.*/

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
)

type Customer struct {
	Id        int     `json:"id,omitempty"`
	FirstName string  `json:"first_name,omitempty"`
	LastName  string  `json:"last_name,omitempty"`
	Email     string  `json:"email,omitempty"`
	Phone     string  `json:"phone,omitempty"`
	Addr      Address `json:"address,omitempty"`
}

type Address struct {
	Street   string `json:"street,omitempty"`
	Postcode string `json:"post,omitempty"`
}

var customers []Customer

func CustomersRootPath(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(customers)
	case "POST":
		var newCustomer Customer
		err := json.NewDecoder(r.Body).Decode(&newCustomer)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		customers = append(customers, newCustomer)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(newCustomer)
	}
}

func CustomerById(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		id := r.URL.Path
		for _, c := range customers {
			if cid := strconv.Itoa(c.Id); cid == id {
				w.Header().Set("Content-Type", "application/json")
				json.NewEncoder(w).Encode(c)
				return
			}
		}
		json.NewEncoder(w).Encode("error: customer not found")
	case "PUT":
		id := r.URL.Path
		for i, c := range customers {
			if cid := strconv.Itoa(c.Id); cid == id {
				w.Header().Set("Content-Type", "application/json")
				err := json.NewDecoder(r.Body).Decode(&c)
				if err != nil {
					http.Error(w, err.Error(), http.StatusBadRequest)
					return
				}
				customers = append(customers[:i], customers[i+1:]...)
				customers = append(customers, c)
				json.NewEncoder(w).Encode(c)
				return
			}
		}
		json.NewEncoder(w).Encode("error: customer not found")
	case "DELETE":
		id := r.URL.Path
		for i, c := range customers {
			if cid := strconv.Itoa(c.Id); cid == id {
				w.Header().Set("Content-Type", "application/json")
				customers = append(customers[:i], customers[i+1:]...)
				json.NewEncoder(w).Encode(c)
				return
			}
		}
		json.NewEncoder(w).Encode("error: customer not found")
	}
}

func main() {
	jsonFile, err := os.Open("customers.json")
	if err != nil {
		fmt.Println(err)
	}

	defer jsonFile.Close()

	jsonBites, _ := io.ReadAll(jsonFile)
	json.Unmarshal(jsonBites, &customers)
	http.HandleFunc("/customers", CustomersRootPath)
	http.Handle("/customers/", http.StripPrefix("/customers/", http.HandlerFunc(CustomerById)))
	http.ListenAndServe(":8000", nil)
}
