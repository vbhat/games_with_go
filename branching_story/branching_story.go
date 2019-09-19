package main

import (
	"bufio"
	"fmt"
	"os"
)

type storyNode struct {
	text    string
	yesPath *storyNode
	noPath  *storyNode
}

func (node *storyNode) play() {
	fmt.Println(node.text)

	scanner := bufio.NewScanner(os.Stdin)

	if node.yesPath != nil && node.noPath != nil {
		for {
			scanner.Scan()
			answer := scanner.Text()

			if answer == "yes" {
				node.yesPath.play()
				break
			} else if answer == "no" {
				node.noPath.play()
				break
			} else {
				fmt.Println("Please enter a valid response")
			}
		}
	}
}

func main() {
	root := &storyNode{"You are the entrance of a cave. Do you want to enter it?", nil, nil}
	win := &storyNode{"Yay, you've won!", nil, nil}
	lose := &storyNode{"Oh no! You've lost :(", nil, nil}
	root.yesPath = lose
	root.noPath = win

	root.play()

}
