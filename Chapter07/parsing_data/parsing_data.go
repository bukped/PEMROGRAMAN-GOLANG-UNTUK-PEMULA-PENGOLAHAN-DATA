package main

import (
	"encoding/csv"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
)

type Person struct {
	Name  string `csv:"name" json:"name" xml:"name"`
	Email string `csv:"email" json:"email" xml:"email"`
}

func main() {
	// Parsing data dari file CSV
	csvFile, err := os.Open("data.csv")
	if err != nil {
		fmt.Println("Error opening CSV file:", err)
		return
	}
	defer csvFile.Close()

	csvReader := csv.NewReader(csvFile)
	csvRecords, err := csvReader.ReadAll()
	if err != nil {
		fmt.Println("Error reading CSV file:", err)
		return
	}

	csvPeople := []Person{}
	for _, record := range csvRecords {
		person := Person{
			Name:  record[0],
			Email: record[1],
		}
		csvPeople = append(csvPeople, person)
	}

	// Parsing data dari file JSON
	jsonFile, err := os.Open("data.json")
	if err != nil {
		fmt.Println("Error opening JSON file:", err)
		return
	}
	defer jsonFile.Close()

	jsonData, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		fmt.Println("Error reading JSON file:", err)
		return
	}

	jsonPeople := []Person{}
	err = json.Unmarshal(jsonData, &jsonPeople)
	if err != nil {
		fmt.Println("Error unmarshaling JSON:", err)
		return
	}

	// Parsing data dari file XML
	xmlFile, err := os.Open("data.xml")
	if err != nil {
		fmt.Println("Error opening XML file:", err)
		return
	}
	defer xmlFile.Close()

	xmlData, err := ioutil.ReadAll(xmlFile)
	if err != nil {
		fmt.Println("Error reading XML file:", err)
		return
	}

	xmlPeople := []Person{}
	err = xml.Unmarshal(xmlData, &xmlPeople)
	if err != nil {
		fmt.Println("Error unmarshaling XML:", err)
		return
	}

	fmt.Println("CSV Data:")
	for _, p := range csvPeople {
		fmt.Println("Name:", p.Name)
		fmt.Println("Email:", p.Email)
		fmt.Println()
	}

	fmt.Println("JSON Data:")
	for _, p := range jsonPeople {
		fmt.Println("Name:", p.Name)
		fmt.Println("Email:", p.Email)
		fmt.Println()
	}

	fmt.Println("XML Data:")
	for _, p := range xmlPeople {
		fmt.Println("Name:", p.Name)
		fmt.Println("Email:", p.Email)
		fmt.Println()
	}
}