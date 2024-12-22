package internal

import "fmt"

type Book struct {
	Genre  string
	Author string
	Title  string
}

func SetLibrary(books []Book, shelfSize int) [][]Book {
	shelfs := make([][]Book, 0)
	currentShelf := make([]Book, 0, shelfSize)
	for _, book := range books {
		currentShelf = append(currentShelf, book)

		if len(currentShelf) == shelfSize {
			shelfs = append(shelfs, currentShelf)
			currentShelf = make([]Book, 0, shelfSize)
		}
	}

	if len(currentShelf) > 0 {
		shelfs = append(shelfs, currentShelf)
	}
	return shelfs
}

func PrintLibrary(lirary [][]Book) {
	for _, shelf := range lirary {
		fmt.Println(shelf)
	}
}
