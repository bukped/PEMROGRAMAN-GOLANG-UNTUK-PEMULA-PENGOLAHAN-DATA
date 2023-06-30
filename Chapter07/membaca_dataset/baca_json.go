package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Person struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func main() {
	file, err := ioutil.ReadFile("data.json")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	var people []Person
	err = json.Unmarshal(file, &people)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
// contoh nama kolom disini yaitu Name dan Email, sesuaikan dengan kolom data sendiri
	for _, p := range people {
		fmt.Println("Name:", p.Name)
		fmt.Println("Email:", p.Email)
		fmt.Println()
	}
}
