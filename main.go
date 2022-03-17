package main

import ( //Import packages
	"fmt"
	"os"
	"strings"
)

func main() {
	m := make(map[string]int) //Create a new map (key:value data structure)

	m["The Weeknd"] = 79600000 //Enter info into map
	m["Justin Bieber"] = 77800000
	m["Coldplay"] = 57300000
	m["Eminem"] = 54900000
	m["Michael Jackson"] = 28100000
	m["Kanye West"] = 48800000
	m["Drake"] = 53800000

	fmt.Println("Hello. Welcome to higher or lower (Spotify monthly listeners edition).") //fmt.Println() prints stuff to the console
	fmt.Println("")
	fmt.Println("You will be presented with two artists. You have to guess if the artist presented to you has more or less monthly listeners than the previously shown artist by selecting higher or lower.")
	fmt.Println("")
	fmt.Println("Let's get started.")

	var flag = true //Set flag to true to keep the while loop going. I had issues when using flag = false to stop the loop so I resulted in os.Exit(3) instead.

	for flag { //while loop
		fmt.Println("The Weeknd has ", m["The Weeknd"], "monthly listeners. Does Justin Bieber have more or less? (enter h/l)")
		var firstAnswer string //initialize answer string
		fmt.Scanln(&firstAnswer) //gets user input
		if strings.ToLower(firstAnswer) == "l" { //checks answers
			fmt.Println("Correct. Justin Bieber has", m["Justin Bieber"], "monthly listener.")
		} else {
			fmt.Println("Wrong. Justin Bieber has", m["Justin Bieber"], "monthly listener. Cya!")
			os.Exit(3)
		}

		//The rest of the code is repitition from previous comments.
		
		fmt.Println("Justin Bieber has ", m["Justin Bieber"], "monthly listeners. Does Michael Jackson have more or less? (enter h/l)")
		var secondAnswer string
		fmt.Scanln(&secondAnswer)
		if strings.ToLower(secondAnswer) == "l" {
			fmt.Println("Correct. Michael Jackson has", m["Michael Jackson"], "monthly listener.")
		} else {
			fmt.Println("Wrong. Michael Jackson has", m["Michael Jackson"], "monthly listener. Cya!")
			os.Exit(3)
		}

		fmt.Println("Michael Jackson has ", m["Michael Jackson"], "monthly listeners. Does Eminem have more or less? (enter h/l)")
		var thirdAnswer string
		fmt.Scanln(&thirdAnswer)
		if strings.ToLower(thirdAnswer) == "l" {
			fmt.Println("Wrong. Eminem has", m["Eminem"], "monthly listener. Cya!")
			os.Exit(3)
		} else {
			fmt.Println("Correct. Eminem has", m["Eminem"], "monthly listener.")
		}

		fmt.Println("Game over. Purchase the full version for $89.99 by contacting Birgir Gauti Stefánsson.")
		os.Exit(3)
	}
}
