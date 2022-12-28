package main

import (
	"fmt"
	"log"
	"os"
)

type addressBook struct {
	Name              string
	MobileNo          int
	AlternateMobileNo int
	Address           string
	City              string
}

/*
description: If error occur print the error and terminate the program
param: err error
*/
func errorHandling(err error) {
	if err != nil {
		log.Fatalf("Err: ", err)
	}
}

/*
description: create address Book of friends
return: addressBookOfFriends []addressBook
*/
func createAddressBook() []addressBook {

	var AddressBookOfFriends []addressBook

	AddressBookOfFriends = []addressBook{

		{Name: "Vikas", MobileNo: 9359153726, AlternateMobileNo: 9921361827, Address: "Gokul nagar", City: "Gachiroli"},
		{Name: "Sanket", MobileNo: 8578565435, Address: "Stadium", City: "Gachiroli"},
		{Name: "Pratik", MobileNo: 8384589657, AlternateMobileNo: 8765478645, Address: "Neharu Ward", City: "Armori"},
		{Name: "Omkar", MobileNo: 9735345643, Address: "Nandanvan", City: "Chamorshi"},
		{Name: "Suraj", MobileNo: 9023547354, Address: "Kargil Chowk", City: "Nagpur"},
		{Name: "Nitin", MobileNo: 8264356455, AlternateMobileNo: 7234684564, Address: "Ram Nagar", City: "Pune"},
		{Name: "Kobra", MobileNo: 8434753654, AlternateMobileNo: 7565347565, Address: "Ganech colony", City: "Mumbai"},
		{Name: "Samrit", MobileNo: 7734563454, Address: "Cankai Nagar", City: "Vardha"},
		{Name: "Sagar", MobileNo: 9854643563, AlternateMobileNo: 8057464578, Address: "Shubhash Ward", City: "Armori"},
		{Name: "Rahul", MobileNo: 8143564654, AlternateMobileNo: 7464568567, Address: "Gokul nagar", City: "Nagpur"},
	}
	return AddressBookOfFriends
}

/*
description: Write addressBook data of friends in file
param: addressBookOfFriends []addressBook
*/
func writeFile(addressBookOfFriends []addressBook) {

	file, err := os.Create("addressBookFile.txt")
	errorHandling(err)
	defer file.Close()

	fileData := "Name  Mobile Numbers  Alternate Mobile Number  Address City\n\n\n"
	_, err = file.WriteString(fileData)
	errorHandling(err)
	fmt.Println(fileData)

	for _, data := range addressBookOfFriends {
		fileData = fmt.Sprintf("%s  %d  %d  %s  %s\n\n\n", data.Name, data.MobileNo, data.AlternateMobileNo, data.Address, data.City)
		_, err = file.WriteString(fileData)
		errorHandling(err)
		fmt.Println(fileData)
	}

}

func main() {

	addressBookOfFriends := createAddressBook()

	writeFile(addressBookOfFriends)

}
