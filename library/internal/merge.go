package internal

func merge(left, right []Book, cmpFunc func(x, y Book) bool) []Book {
	result := make([]Book, 0, len(left)+len(right))
	i, j := 0, 0

	for i < len(left) && j < len(right) {
		if cmpFunc(left[i], right[j]) {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}

	// Добавляем оставшиеся элементы из левого массива
	for i < len(left) {
		result = append(result, left[i])
		i++
	}

	// Добавляем оставшиеся элементы из правого массива
	for j < len(right) {
		result = append(result, right[j])
		j++
	}

	return result
}

func mergeSort(books []Book, cmpFunc func(x, y Book) bool) []Book {
	if len(books) <= 1 {
		return books
	}

	mid := len(books) / 2
	left := mergeSort(books[:mid], cmpFunc)
	right := mergeSort(books[mid:], cmpFunc)

	return merge(left, right, cmpFunc)
}
