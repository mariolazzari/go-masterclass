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

type ContactInfo struct {
	Email string
	Phone string
}

func (ci ContactInfo) DisplayContact() string {
	return fmt.Sprintf("%s, %s", ci.Email, ci.Phone)
}

type Company struct {
	Name string
	Address
	ContactInfo
	BusinessType string
}

func (c Company) GetProfile() {
	fmt.Println("Company Name:", c.Name)
	fmt.Println("Address     :", c.FullAddress())
	fmt.Println("Company Street (promoted):", c.Street)

}

type CompanyWithOwnMail struct {
	Company
	Email string // shadowing
}

func main() {
	comp := Company{
		Name: "Company 1",
		Address: Address{
			City:    "City 1",
			Street:  "Street 1",
			State:   "State 1",
			ZipCode: "12345",
		},
		BusinessType: "Type 1",
	}

	comp.GetProfile()

	comp2 := CompanyWithOwnMail{
		Company: Company{
			Name: "Company 12",
			ContactInfo: ContactInfo{
				Email: "original@mail.com",
			},
			Address: Address{
				City:    "City 1",
				Street:  "Street 1",
				State:   "State 1",
				ZipCode: "12345",
			},
			BusinessType: "Type 1",
		},
		Email: "shadowing@mail.com",
	}

	comp2.GetProfile()
	fmt.Println("Shadowing Email:", comp2.Email)

}
