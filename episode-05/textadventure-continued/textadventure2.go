// Changed linked lists *choices to an array

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type choice struct {
	cmd         string
	description string
	nextNode    *storyNode
}

type storyNode struct {
	text    string
	choices []*choice
}

func (node *storyNode) addChoice(cmd string, description string, nextNode *storyNode) {
	choice := &choice{cmd, description, nextNode}
	node.choices = append(node.choices, choice)
}

func (node *storyNode) render() {
	fmt.Println(node.text)
	if node.choices != nil {
		for _, choice := range node.choices { //range will loops through all entries from the array. We don't want to see the index so we use an _, otherwise it would look like for i, choice := range node.choices {
			fmt.Println(choice.cmd, choice.description)
		}
	}
}

func (node *storyNode) executeCmd(cmd string) *storyNode {
	for _, choice := range node.choices {
		if strings.ToLower(choice.cmd) == strings.ToLower(cmd) {
			return choice.nextNode
		}
	}
	fmt.Println("Sorry I didn't understand that.")
	return node
}

var scanner *bufio.Scanner

func (node *storyNode) play() {
	node.render()
	if node.choices != nil {
		scanner.Scan()
		node.executeCmd(scanner.Text()).play()

	}
}

func main() {

	scanner = bufio.NewScanner(os.Stdin)

	start := storyNode{text: `
	You wake up in a dark cave.
	You see see that there are 3 passages leading out.
	The north passage leads into darkness
	To the south, a passage appears to head upwards.
	The eastern passage appears flat, but has a disturbing smell coming from it.`}

	darkRoom := storyNode{text: "It is pitch black. You cannot see a thing."}

	darkRoomLit := storyNode{text: "The dark passage is now lit by your torch. You can continue north or head back south"}

	grue := storyNode{text: "Whilst stumbling around in the darkness you are eaten by a grue"}

	trap := storyNode{text: "You trigger a hidden trap and fall to your doom"}

	treasure := storyNode{text: "You arrive at a small chamber, filled with wonderful treasure!"}

	start.addChoice("N", "Go North", &darkRoom)
	start.addChoice("S", "Go South", &darkRoom)
	start.addChoice("E", "Go East", &trap)

	darkRoom.addChoice("S", "Try to go back south", &grue)
	darkRoom.addChoice("O", "Light up Lantern", &darkRoomLit)

	darkRoomLit.addChoice("N", "Go North", &treasure)
	darkRoomLit.addChoice("S", "Try to go back south", &start)

	start.play()

	fmt.Println()
	fmt.Println("The End.")

}
