package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"math/rand"
	"os"
)

var facts = []string{ //Create array with random facts. array of type string.
	"Until 1994, world maps and globes in Albania only had Albania on them.",
	"February 1865 is the only month in recorded history not to have a full moon.",
	"In the Philippine jungle, the yo-yo was first used as a weapon.",
	"60 Minutes is the only TV show without a theme song.",
	"Each king in a deck of playing cards represents a great king from history. Spades - King David, Clubs - Alexander the Great, Hearts - Charlemagne, and Diamonds - Julius Caesar.",
}

func main() {
	a := app.New()                           //creates app
	w := a.NewWindow("GUI Test Application") //adds window to app
	w.Resize(fyne.NewSize(2000, 2000))       //sets size of window. i set it as a large float since the window automatically scales if it is too large.

	hello := widget.NewLabel("Hello CS495. Label Test.") //creates label and displays it on the window

	radioGroup := widget.NewRadioGroup([]string{"Yes", "No"}, func(string) {}) //Create radio group. Where the radiobuttons are held.
	radioGroup.SetSelected("Yes")                                              //Sets default option selected.
	radioLabel := widget.NewLabel("")                                          //Initialize label as blank

	factText := widget.NewLabel("")

	w.SetContent(container.NewVBox( //Applying the things or concepts to your window. Anything in here will appear in order.
		hello,
		widget.NewButton("Click me", func() { //create a new button. apply the func() when it is pressed.
			hello.SetText("You clicked the button so you changed the text in the hello label.") //this happens when the button is pressed.
		}),
		radioGroup, //apply the radiogroup to the window.
		widget.NewButton("Submit", func() { //create another button with some logic.
			if radioGroup.Selected == "Yes" {
				radioLabel.SetText("You chose 'yes'")
			} else {
				radioLabel.SetText("You chose 'no'")
			}
		}),
		radioLabel,
		widget.NewSeparator(), //This is the line seen on the window. A simple separator.
		widget.NewButton("Get Fact", func() {
			factText.SetText(facts[rand.Intn(3)]) //Randomizes a fact each time the button is pressed.
		}),
		factText,
		widget.NewButton("Exit", func() {
			os.Exit(3) //exit condition.
		}),
	))

	w.ShowAndRun() //runs application
}
