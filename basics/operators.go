package main

import "fmt"

func main() {
	// arithmetic operators
	//examples of arithemtics :
	// + - / * %
	var num1 = 10
	var num2 = 8
	var sum = num1 + num2
	var minus = num1 - num2
	var divide = num1 / num2
	var multiplication = num1 * num2

	var remainder = num1 % num2

	fmt.Println(sum, minus, divide, multiplication, remainder)

	// relational operators

	var result = num1 > num2
	var isTrue = num2 > num1
	var answer = num1 != num2

	fmt.Println(result, isTrue, answer)

	//logical operators
	//invite guests to my party

	var name = "olive"
	var age = 25

	// i want strickly only those with  names like my name  and those who are above 25 to come to my party..var

	var invitedToParty = name == "olive" && age > 25
	fmt.Println(invitedToParty)

	// i want any other dude with diffrent name but only age should be above 25 ( using or which is ||)
	var inviteToParty = name == "olive" || age > 25
	fmt.Println(inviteToParty)

	// i dont want to be invuted on my own party coz am boring LOL haha !

	var invitesToParty = !(name == "olive") && (age >= 25 && age < 90)
	fmt.Println(invitesToParty)

}
