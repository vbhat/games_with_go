package main

import "fmt"

type storyPage struct {
	text     string
	nextPage *storyPage
}

func tellStory(page *storyPage) {
	if page == nil {
		return
	}
	fmt.Println(page.text)
	tellStory(page.nextPage)
}

func insertPage(prevP *storyPage, text string) {
	newPage := storyPage{text, prevP.nextPage}
	prevP.nextPage = &newPage
}

func deletePage(startPage, page *storyPage) {
	if startPage == nil || page == nil {
		return
	}
	if startPage.nextPage == page {
		startPage.nextPage = page.nextPage
		return
	}
	deletePage(startPage.nextPage, page)
}

func main() {
	page1 := storyPage{"It is a dark and stormy night", nil}
	page2 := storyPage{"Your mission is to find the sacred helmet before the bad guys do", nil}
	page3 := storyPage{"You see a troll ahead", nil}

	page1.nextPage = &page2
	page2.nextPage = &page3

	insertPage(&page3, "What will you do now?")

	tellStory(&page1)

	deletePage(&page1, &page3)

	tellStory(&page1)

}
