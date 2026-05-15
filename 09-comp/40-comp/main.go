package main

import "fmt"

type Address struct {
	Street  string
	City    string
	State   string
	ZipCode string
}

func (a Address) FullAddress() string {
	if a.Street == "" && a.City == "" {
		return "No address provided"
	}

	return fmt.Sprintf("%s, %s, %s %s", a.Street, a.City, a.State, a.ZipCode)
}

type Customer struct {
	CustomerID      int
	Name            string
	Mail            string
	BillingAddress  Address
	ShippingAddress Address
}

func (c Customer) PrintDetails() {
	fmt.Println("Customer ID     :", c.CustomerID)
	fmt.Println("Customer Name   :", c.Name)
	fmt.Println("Customer eMail  :", c.Mail)
	fmt.Println("Billing Address :", c.CustomerID)
	fmt.Println("Shipping Address:", c.ShippingAddress.FullAddress())
	fmt.Println("Billing  Address:", c.BillingAddress.FullAddress())
}

func main() {
	cust := Customer{
		CustomerID: 1,
		Name:       "Mario",
		Mail:       "mario.lazzari@gmail.com",
		BillingAddress: Address{
			Street:  "stree 1",
			City:    "City 1",
			State:   "State 1",
			ZipCode: "12345",
		},
		ShippingAddress: Address{
			Street:  "stree 2",
			City:    "City 2",
			State:   "State 2",
			ZipCode: "54321",
		},
	}

	cust.PrintDetails()
}
