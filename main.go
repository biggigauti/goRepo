package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	m := make(map[string]int)

	m["The Weeknd"] = 79600000
	m["Justin Bieber"] = 77800000
	m["Coldplay"] = 57300000
	m["Eminem"] = 54900000
	m["Michael Jackson"] = 28100000
	m["Kanye West"] = 48800000
	m["Drake"] = 53800000

	fmt.Println("Hello. Welcome to higher or lower (Spotify monthly listeners edition).")
	fmt.Println("")
	fmt.Println("You will be presented with two artists. You have to guess if the artist presented to you has more or less monthly listeners than the previously shown artist by selecting higher or lower.")
	fmt.Println("")
	fmt.Println("Let's get started.")

	var flag = true

	for flag {
		fmt.Println("The Weeknd has ", m["The Weeknd"], "monthly listeners. Does Justin Bieber have more or less? (enter h/l)")
		var firstAnswer string
		fmt.Scanln(&firstAnswer)
		if strings.ToLower(firstAnswer) == "l" {
			fmt.Println("Correct. Justin Bieber has", m["Justin Bieber"], "monthly listener.")
		} else {
			fmt.Println("Wrong. Justin Bieber has", m["Justin Bieber"], "monthly listener. Cya!")
			os.Exit(3)
		}

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
