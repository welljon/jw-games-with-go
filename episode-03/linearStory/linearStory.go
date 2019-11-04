package main

import "fmt"

// Using pointers for fun
// This is a linked list

type storyPage struct {
	text     string
	nextPage *storyPage
}

func playStory(page *storyPage) {
	if page == nil {
		return
	}
	fmt.Println(page.text)
	playStory(page.nextPage)
}

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
	page2 := storyPage{"The rain was bouncing from the windscreen of the motor, blocking out the view", nil}
	page3 := storyPage{"The car hums as you pull up to the kerb.", nil}
	page4 := storyPage{"Your electric umbrella hisses into life, creating a sphere of dry around you.", nil}

	page1.nextPage = &page2
	page2.nextPage = &page3
	page3.nextPage = &page4

	playStory(&page1)
}
