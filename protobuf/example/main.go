package main

import (
	"fmt"
	"log"
	"os"
	"proto-example/employees"
	"strconv"

	"github.com/golang/protobuf/jsonpb"
	"google.golang.org/protobuf/proto"
)

func main() {
	fmt.Println("Protobufs in Go")

	// generate employees
	emps := generateEmployees(100)

	// marshalls employees to wire-format
	dataBytes, err := proto.Marshal(emps)
	if err != nil {
		log.Fatal("error in marshalling employees to bytes")
	}

	// create json marshaller to convert data into json format
	m := jsonpb.Marshaler{}
	dataJSON, err := m.MarshalToString(emps)
	if err != nil {
		log.Fatal("error in marshalling employees to json")
	}

	// write both data formats to disk to compare size
	writeToFile("generated.bytes", dataBytes)
	writeToFile("generated.json", []byte(dataJSON))

	// read employees data from file and unmarshall to new employees
	dataFromFile, err := readFromFile("generated.bytes")
	if err != nil {
		log.Fatal("error in reading employees data from file")
	} 
	empsCopy := &employees.Employees{}
	if err := proto.Unmarshal(dataFromFile, empsCopy); err != nil {
		log.Fatal("error in unmarshalling employees")
	}

	fmt.Println(empsCopy)
}

func generateEmployees(num int) *employees.Employees {
	emps := &employees.Employees{}

	for i := 0; i < num; i++ {
		id := strconv.Itoa(i)
		emps.Employees = append(emps.Employees, &employees.Employee{
			Id: int32(i), 
			Firstname: "FirstName_" + id,
			Lastname: "LastName_" + id,
			Age: 30,
			Address: "Address City, Country",
		})
	}

	return emps
}

func writeToFile(filename string, content []byte) error {
	if err := os.WriteFile(filename, content, 0755); err != nil {
		return err
	}

	fi, err := os.Stat(filename)
	if err != nil {
		return err
	}

	fmt.Printf("%s: %v\n", filename, fi.Size())

	return nil
}

func readFromFile(filename string) ([]byte, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	return data, nil
} 