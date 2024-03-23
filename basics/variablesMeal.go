package main

import "fmt"

func main() {
	// declaring and assigning
	var favMeal = "chapo and beans"

	fmt.Println("favMeal")

	//reassigning (when reassigning you dont rewrite the var)
	favMeal = "-------Birani is my fav , wallahi tena---------"

	fmt.Println(favMeal)

	//love another meal
	var otherFavMeal = "chips and chicken"
	fmt.Println(otherFavMeal)

	//variables with data types using examples of swahili restaurant

	var swahiliMeals string
	var mealPrice int
	var isMealAvailable bool
	swahiliMeals = "mahamri"

	fmt.Println(swahiliMeals, mealPrice, isMealAvailable)

	//compound creation (creating multiple variables in a compound )
	//lets say my personality is
	//-----------------------

	// var favNumber = 12
	// var favSport = "chess"
	// var favKicks = "Addidas yeezy "

	// we can do this in compund creation
	var favColor, favTown, favGame = "gray", "Mombasa", "GTA"

	fmt.Println(favColor, favTown, favGame)

	// block creation
	// using my personlaity example we can use block creation to ome up with above data

	var (
		favNumber = 27
		favSport  = "chess"
		favKicks  = "Addidas yeezy"
	)

	fmt.Println(favNumber, favSport, favKicks)

	//using shorcut (:=) instead of using var as keyword
	favPet := "cat"

	fmt.Println(favPet)

	// using shortcut (:=) for compound creation
	favFruits, favLogo, favCompany := "mango", "nike", "payd"

	fmt.Println(favFruits, favLogo, favCompany)

	// using shortcut (:=) for block creation
	favShoes,
		FavAnimal,
		FavCountry := "Timberland", "Lion", "Germany"
	fmt.Println(favShoes, FavAnimal, FavCountry)

	favDrink := "coffee"
	fmt.Println(favDrink)

	cup1, cup2, cup3 := "tea", "black coffee", "orange juice"
	fmt.Println(cup1, cup2, cup3)
	// cup3 := "capucciono cofee"---- you cant reassign like this but you can do like this -
	cup3 = "capucciono cofee"

	fmt.Println(cup3)

	//using constant variables . (const)
	const myName = "olive"
	// myName = "njeri"----------can not use myName varibale coz its constant it cannot change.

}
