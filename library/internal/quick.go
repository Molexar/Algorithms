package internal

func partition(books []Book, cmpFunc func(x, y Book) bool, low, high int) int {
	pivot := books[high]
	i := low - 1

	for j := low; j < high; j++ {
		if cmpFunc(books[j], pivot) {
			i++
			books[i], books[j] = books[j], books[i]
		}
	}
	books[i+1], books[high] = books[high], books[i+1]
	return i + 1
}

func quickSort(books []Book, cmpFunc func(x, y Book) bool, low, high int) {
	if low < high {
		pi := partition(books, cmpFunc, low, high)
		quickSort(books, cmpFunc, low, pi-1)
		quickSort(books, cmpFunc, pi+1, high)
	}
}
