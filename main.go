package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
)

//Employee ... your employee field structure
type Employee struct {
	FirstName string
	LastName  string
	Address   string
	Age       int
	Email     string
}

func main() {
	csvFile, err := os.Open("./data.csv")
	if err != nil {
		fmt.Print(err.Error())
	}
	defer csvFile.Close()

	reader := csv.NewReader(csvFile)
	reader.FieldsPerRecord = -1

	csvData, err := reader.ReadAll()
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	var emp Employee
	var employees []Employee

	for _, each := range csvData {
		emp.FirstName = each[0]
		emp.LastName = each[1]
		emp.Address = each[2]
		emp.Age, _ = strconv.Atoi(each[3])
		emp.Email = each[4]
		//Append the instance of the current emp model values to a slice of Employee
		employees = append(employees, emp)
	}

	jsonData, err := json.Marshal(employees)
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	fmt.Print(string(jsonData))

	jsonFile, err := os.Create("./data.json")
	if err != nil {
		fmt.Print(err.Error())
	}
	defer jsonFile.Close()

	jsonFile.Write(jsonData)
	jsonFile.Close()

}
