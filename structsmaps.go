package main

import "fmt"

type Customer struct {
	CustID int
	Name   string
	Salary float64
}

func main() {
	CreateCust()
	add("1", "John")
	get("1", "John")

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
	fmt.Println(cust2)
	//fmt.Printf(fmt.Sprintf("Customer 2 is : %+v\n", cust2))
	fmt.Println(cust3)

	custMap := make(map[int]Customer)

	custMap[1] = cust
	custMap[2] = cust1
	custMap[3] = cust2
	custMap[4] = cust3

	delete(custMap, 3)

	for key, value := range custMap {

		fmt.Println("Customer :", key, "value :", value)
	}

}

var data = make(map[string]string)

func add(k string, v string) {
	if _, ok := data[k]; ok {
		data[k] = v

	}
	fmt.Println(k, v)
}

func get(k string, v string) {
	if _, ok := data[k]; ok {
		data[k] = v
	}
	fmt.Println(k, v)
}
