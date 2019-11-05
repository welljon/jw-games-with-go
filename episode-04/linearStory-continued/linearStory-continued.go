package main

import "fmt"

// Using pointers for fun
// This is a linked list

type storyPage struct {
	text     string
	nextPage *storyPage
}

func (page *storyPage) playStory() {
	// if page == nil { // This example is recursive
	// 	return
	// }
	// fmt.Println(page.text)
	// page.nextPage.playStory()
	for page != nil { // This example is non-revursive
		fmt.Println(page.text)
		page = page.nextPage
	}
}

func (page *storyPage) addToEnd(text string) {
	for page.nextPage != nil {
		page = page.nextPage
	}
	page.nextPage = &storyPage{text, nil}
}

func (page *storyPage) addAfter(text string) {
	newPage := &storyPage{text, page.nextPage}
	page.nextPage = newPage
}

// Delete

func main() {

	//scanner := bufio.NewScanner(os.Stdin)
	// We could do it this way but it's got a lot of repeating
	// fmt.Println("It was a dark and dreary night.")
	// scanner.Scan()
	// fmt.Println("The rain was bouncing from the windscreen of the motor, blocking out the view.")
	// scanner.Scan()
	// fmt.Println("You pull the car up to the kerb and get out.")
	// scanner.Scan()
	// fmt.Println("Your electric umbrella hisses into life, creating a sphere of dry around you.")
	// scanner.Scan()
	// Better to do it like this
	page1 := storyPage{"It was a dark and dreary night.", nil}
	page1.addToEnd("The rain was bouncing from the windscreen of the motor, blocking out the view")
	page1.addToEnd("The car's EPD hums as you pull up to the kerb.")
	page1.addToEnd("Your electric umbrella hisses into life, creating a sphere of dry around you.")

	page1.addAfter("Testing AddAfter")
	page1.playStory()
}

// Functions - has a return value - may also execute commands
// Procedures  - has no return value, just executes commands
// Methods - functions attached to a struct/object/etc
