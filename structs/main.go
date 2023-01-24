package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode string
}
type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

func main() {
	myinfo := person{
		firstName: "sanket",
		lastName:  "Mote",
		contact: contactInfo{
			email:   "sanket@gmail.com",
			zipCode: "415409",
		},
	}
	fmt.Println(myinfo)

}
