package main

import (
	"encoding/xml"
	"fmt"
	"os"
)

type Person struct {
	Name  string `xml:"name"`
	Email string `xml:"email"`
}

type People struct {
	XMLName xml.Name `xml:"people"`
	Persons []Person `xml:"person"`
}

func main() {
	xmlFile, err := os.Open("data.xml")
	if err != nil {
		fmt.Println("Error opening XML file:", err)
		return
	}
	defer xmlFile.Close()

	var people People
	err = xml.NewDecoder(xmlFile).Decode(&people)
	if err != nil {
		fmt.Println("Error decoding XML:", err)
		return
	}

	for _, p := range people.Persons {
		fmt.Println("Name:", p.Name)
		fmt.Println("Email:", p.Email)
		fmt.Println()
	}
}
