package internal

func bubbleSort(books []Book, cmpFunc func(x, y Book) bool) {
	n := len(books)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if !cmpFunc(books[j], books[j+1]) {
				// Обмен элементов
				books[j], books[j+1] = books[j+1], books[j]
			}
		}
	}
}
