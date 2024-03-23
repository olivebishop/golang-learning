package main

import "fmt"

//switches are used to create a clean code statement where you need to repeat statement using "if else" you can use switch
// below is an example of sounds produced by different animals.
// default animal is cat , if you change diffent animals you will get different sound when you change to an animals not listed on the swtich you will get "i dont know the animal"

func main() {
	animal := "cat"

	switch animal {
	case "cat":
		fmt.Println("Meow")

	case "dog":
		fmt.Println("Bark")

	case "horse":
		fmt.Println("Neigh")

	case "frog":
		fmt.Println("Ribbit")

	default:
		fmt.Println("Dont know the animal")
	}

}
