package main

import "fmt"

func main() {
	// manimum age for you to come to my party is 25 years and above
	// the following shows if else statement / conditions of friends and geusts invited to my party
	age := 18

	if age >= 25 {
		fmt.Println("you are allowed to the party")

	} else if age >= 20 {
		fmt.Println("yes you are a teen but we need above 25 years and above  dude ! HAHA")

	} else {
		fmt.Println("under age are not allowed !")
	}

}
