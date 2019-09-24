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

	for _, val := range node.choices {
		fmt.Println(val.cmd, ": ", val.description)
	}
}

func (node *storyNode) executeCmd(cmd string) *storyNode {
	for _, val := range node.choices {
		if strings.ToLower(val.cmd) == strings.ToLower(cmd) {
			return val.nextNode
		}
	}

	fmt.Println("Sorry, I do not understand")
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
	You are in a large chamber, deep underground.
	You see three passages leasing out. A north passage leads into darkness.
	To the south, a passage appears to head upward. The eastern passage appears
	flat and well traveled`}

	darkRoom := storyNode{text: "It is pitch black, You cannot see a thing."}

	darkRoomLit := storyNode{text: "The dark passage is now lit by the lantern. You can continue north or head back south."}

	grue := storyNode{text: "While standing in the darkness, you are eaten by a grue."}

	trapDoor := storyNode{text: "You head down the well traveled path, when all of a sudden a trap door opens and you fall to your death."}

	treasure := storyNode{text: "You arrive at a small chamber, filled with treasure!"}

	start.addChoice("N", "Go North", &darkRoom)
	start.addChoice("S", "Go South", &darkRoom)
	start.addChoice("E", "Go East", &trapDoor)

	darkRoom.addChoice("S", "Try to go back south", &grue)
	darkRoom.addChoice("O", "Turn on lantern", &darkRoomLit)

	darkRoomLit.addChoice("N", "Go North", &treasure)
	darkRoomLit.addChoice("S", "Go South", &start)

	start.play()

	fmt.Println()
	fmt.Println("The End")
}
