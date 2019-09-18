package main

import "fmt"

type storyPage struct {
	text     string
	nextPage *storyPage
}

// func playStory(page *storyPage) {
// 	for page != nil {
// 		fmt.Println(page.text)
// 		page = page.nextPage
// 	}
// }

func (page *storyPage) playStory() {
	for page != nil {
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

// func insertPage(prevP *storyPage, text string) {
// 	newPage := storyPage{text, prevP.nextPage}
// 	prevP.nextPage = &newPage
// }

func deletePageRecursive(startPage, page *storyPage) {
	if startPage == nil || page == nil {
		return
	}
	if startPage.nextPage == page {
		startPage.nextPage = page.nextPage
		return
	}
	deletePageRecursive(startPage.nextPage, page)
}

func main() {
	page1 := storyPage{"It is a dark and stormy night", nil}
	page1.addToEnd("Your mission is to find the sacred helmet before the bad guys do")
	page1.addToEnd("You see a troll ahead")

	// insertPage(&page3, "What will you do now?")

	page1.addAfter("You are tired and hungry")

	page1.playStory()

}
