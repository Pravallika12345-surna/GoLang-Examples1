package main

import "fmt"

type Customer struct {
	CustID int
	Name   string
	Salary float64
}

func main() {
	CreateCust()
	init()
}

func CreateCust() {

	//creating an instance by using struct variable
	var cust Customer
	cust.CustID = 101
	cust.Name = "John"
	cust.Salary = 30000

	//using struct literal
	cust1 := Customer{
		CustID: 102,
		Name:   "charles",
		Salary: 40000,
	}
	// without using fields
	cust2 := Customer{
		103,
		"tony",
		50000,
	}
	//partial execution
	cust3 := Customer{
		Name: "jack",
	}

	fmt.Println(cust)
	fmt.Println(cust1)
	fmt.Printf(fmt.Sprintf("Customer 2 is : %+v\n", cust2))
	fmt.Println(cust3)
}

func init() {
	custMap := make(map[int]Customer)

	custMap[1] = cust
	custMap[2] = cust1
	custMap[3] = cust2
	custMap[4] = cust3

	for key, value := range custMap {

		fmt.Println("Customer :", key, "value :", value)
	}
}

/*
	delete(students, 2)

	//for retreiving
	fmt.Println("Name of Student 1 is : ", students[1])

	//for checking key exists or not

	if _, ok := students[2]; ok {
		fmt.Println("Student Found")
	} else {
		fmt.Println("Student not Found")
	}  */
