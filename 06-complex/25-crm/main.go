package main

import "fmt"

type Contact struct {
	ID    int
	Name  string
	Email string
	Phone string
}

var contactList []Contact
var contactIndexByName map[string]int
var nextID int = 1

func init() {
	contactList = make([]Contact, 0)
	contactIndexByName = make(map[string]int)
}

func addContact(name, email, phone string) {
	if _, exists := contactIndexByName[name]; exists {
		fmt.Printf("Contact %s exists\n", name)
		return
	}

	contactList = append(contactList, Contact{
		ID:    nextID,
		Name:  name,
		Email: email,
		Phone: phone,
	})
	contactIndexByName[name] = nextID

	fmt.Printf("Contact %s added\n", name)
	nextID++
}

func findContact(name string) *Contact {
	if idx, ok := contactIndexByName[name]; ok {
		return &contactList[idx]
	}
	return nil
}

func listContacts() {
	if len(contactList) == 0 {
		fmt.Println("No contact found")
	}

	for _, c := range contactList {
		fmt.Println(c)
	}
}

func main() {
	listContacts()
	addContact("Mario", "mario@lazzari.it", "123-1234")
	addContact("Maria", "maria@lazzari.it", "123-4321")
	listContacts()

	mario := findContact("Mario")
	if mario == nil {
		fmt.Println("Mario not found")
	} else {
		fmt.Println("Found:", *mario)
	}
}
