package main

/* Create as REST API to handle customer data - CRUD.
   Start simple, with customer type having:
   - ID (primary key)
   - First name
   - Last name
   - Address -> another struct
   - E-mail (unique)
   - Phone number (unique)
   Address:
   - Street and number
   - Postal code */

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type Customer struct {
	Id        int     `json:"id"`
	FirstName string  `json:"first_name"`
	LastName  string  `json:"last_name"`
	Email     string  `json:"email"`
	Phone     string  `json:"phone"`
	Addr      Address `json:"address"`
}

type Address struct {
	Street   string `json:"street"`
	Postcode string `json:"post"`
}

func main() {
	jsonFile, err := os.Open("customers.json")
	if err != nil {
		fmt.Println(err)
	}

	defer jsonFile.Close()

	jsonBites, _ := io.ReadAll(jsonFile)
	var customers []Customer
	json.Unmarshal(jsonBites, &customers)

	for _, c := range customers {
		fmt.Println(c.FirstName, c.LastName)
		fmt.Println(c.Addr)
		fmt.Println()
	}
}
