package main

import (
	"encoding/json"
	"fmt"
	"log"
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

func errorHanding(err error) {
	if err != nil {
		log.Fatalf("Err: ", err)
	}
}

/*
description: convert slice to json Object
param: (slice of struct) addressBook []person
return slicOfStructToJson []byte
*/
func sliceOfStructToJson(addressBook []person) []byte {
	slicOfStructToJson, err := json.MarshalIndent(addressBook, "", "  ")
	errorHanding(err)
	return slicOfStructToJson
}

func main() {

	addressBook := make([]person, 0, 10)

	addressBook = append(addressBook, person{Name: "Vikas", MobileNo: 9359153726, AlternateMobileNo: 9921361827, Address: "Gokul nagar", City: "Gachiroli"})
	addressBook = append(addressBook, person{Name: "Sanket", MobileNo: 8578565435, Address: "Stadium", City: "Gachiroli"})
	addressBook = append(addressBook, person{Name: "Pratik", MobileNo: 8384589657, AlternateMobileNo: 8765478645, Address: "Neharu Ward", City: "Armori"})
	addressBook = append(addressBook, person{Name: "Omkar", MobileNo: 9735345643, Address: "Nandanvan", City: "Chamorshi"})
	addressBook = append(addressBook, person{Name: "Suraj", MobileNo: 9023547354, Address: "Kargil Chowk", City: "Nagpur"})
	addressBook = append(addressBook, person{Name: "Nitin", MobileNo: 8264356455, AlternateMobileNo: 7234684564, Address: "Ram Nagar", City: "Pune"})
	addressBook = append(addressBook, person{Name: "Kobra", MobileNo: 8434753654, AlternateMobileNo: 7565347565, Address: "Ganech colony", City: "Mumbai"})
	addressBook = append(addressBook, person{Name: "Samrit", MobileNo: 7734563454, Address: "Cankai Nagar", City: "Vardha"})
	addressBook = append(addressBook, person{Name: "Sagar", MobileNo: 9854643563, AlternateMobileNo: 8057464578, Address: "Shubhash Ward", City: "Armori"})
	addressBook = append(addressBook, person{Name: "Rahul", MobileNo: 8143564654, AlternateMobileNo: 7464568567, Address: "Gokul nagar", City: "Nagpur"})

	AddressBookInJson := string(sliceOfStructToJson(addressBook))
	fmt.Println("JsonAddressBook: ", AddressBookInJson)
}
