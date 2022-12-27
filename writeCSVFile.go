package main

import (
	"fmt"
	"log"
	"os"
)

type person struct {
	Name              string
	MobileNo          int
	AlternateMobileNo int
	Address           string
	City              string
}

/*
description: to handle error
param: err error
*/
func errorHandling(err error) {
	if err != nil {
		log.Fatalf("Err: ", err)
	}
}

/*
description: Create addressBook slice of stuct person and add 10 friends data
return: addressBook []person
*/
func addressBook() []person {

	addressBook := []person{

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
	return addressBook
}

/*
description: write address Book data in csv file
param: addressBook []person
*/
func writeFile(addressBook []person) {

	file, err := os.Create("addressBook.csv")
	errorHandling(err)
	defer file.Close()

	for _, data := range addressBook {
		fileData := fmt.Sprintf("%s,%d,%d,%s,%s\n", data.Name, data.MobileNo, data.AlternateMobileNo, data.Address, data.City)
		_, err := file.WriteString(fileData)
		errorHandling(err)
		fmt.Print(fileData)
	}
}

func main() {

	//create Address Book
	addressBook()

	//Write address book data in file
	writeFile(addressBook())
}
