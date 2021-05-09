package main

import (
	"fmt"
)

var m int
	
var contact map[string]string = make(map[string]string)

func greetScreen() {
	fmt.Println("Hi, this is simple CLI contact app!")
	fmt.Println("Choose menu:")
	fmt.Println("1. List")
	fmt.Println("2. Create")
	fmt.Println("3. Edit")
	fmt.Println("4. Delete")
	fmt.Println("5. Exit")
}

func mainMenu() {
	greetScreen()
	fmt.Print("Please choose menu: ")
	fmt.Scanf("%d", &m)
	switch m {
	case 1:
		listContact()
	case 2:
		addContact()
	case 3:
		editContact()
	case 4:
		deleteContact()
	case 5:
		fmt.Println("The app was closed")
	default:
		fmt.Println("ouch, you didn't input truth menu!")
		fmt.Println("The app was closed")
	}
}

func returnMainMenu() {
	var yesNo string

	fmt.Print("Back to main menu (y/n): ")
	fmt.Scanf("%s", &yesNo)

	switch yesNo {
	case "y":
		greetScreen()
		mainMenu()
	case "n":
		fmt.Println("App exit")
	default:
		fmt.Println("ouch, you didn't input truth menu!")
		fmt.Println("The app was closed")
	}
}

func listContact() {
	if len(contact) == 0 {
		fmt.Println("The contact list was empty. please add contact first!")
	} else {
		i := 1
		for name, val := range contact {
			fmt.Println(i,".", name, ":", val)
			i++
		}
	}

	returnMainMenu()
}

func addContact() {
	var yesNo string
	
	contactForm()

	fmt.Print("Do you want to add again?(y/n)")
	fmt.Scanf("%s", &yesNo)

	switch yesNo {
	case "y":
		contactForm()
	case "n":
		returnMainMenu()
	default:
		fmt.Println("ouch, you didn't input truth menu!")
		fmt.Println("The app was closed")
	}
}

func contactForm() {
	var shortName string
	var phoneNumber string
	fmt.Print("Input shortname: ")
	fmt.Scanf("%s", &shortName)
	fmt.Print("Input phone number: ")
	fmt.Scanf("%s", &phoneNumber)

	contact[shortName] = phoneNumber
}

func editContact() {
	var contactName string
	fmt.Print("Search contact by name (case sensitive): ")
	fmt.Scanf("%s", &contactName)

	if len(contact) == 0 {
		fmt.Println("The contact list was empty. please add contact first!")
	} else {
		for name, _ := range contact {
	
			if name != contactName {
				fmt.Println("The contact name is not found!")
				fmt.Print("Input again: ")
				fmt.Scanf("%s", &contactName)
				if (name == contactName) {
					break
				}
				returnMainMenu()
			}
		}

		var phoneNumber string
		fmt.Print("Input new phone number: ")
		fmt.Scanf("%s", &phoneNumber)

		contact[contactName] = phoneNumber
	}
	
	returnMainMenu()
}

func deleteContact() {
	var contactName string
	fmt.Print("Search contact by name (case sensitive): ")
	fmt.Scanf("%s", &contactName)

	if len(contact) == 0 {
		fmt.Println("The contact list was empty. please add contact first!")
	} else {
		for name, _ := range contact {
	
			if name != contactName {
				fmt.Println("The contact name is not found!")
				fmt.Print("Input again: ")
				fmt.Scanf("%s", &contactName)
				if (name == contactName) {
					break
				}
				returnMainMenu()
			}
		}

		delete(contact, contactName)
		fmt.Println("Contact with name:", contactName, "was deleted")
	}
	
	returnMainMenu()
}

func main() {
	mainMenu()
}