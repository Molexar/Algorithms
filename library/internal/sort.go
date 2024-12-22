package internal

type Algorithm int

type Opts struct {
	Algorithm Algorithm
}

const (
	QuickSort Algorithm = iota
	BubbleSort
	MergeSort
)

func Sort(books []Book, cmpFunc func(x, y Book) bool, opts Opts) {
	switch opts.Algorithm {
	case QuickSort:
		quickSort(books, cmpFunc, 0, len(books)-1)
	case MergeSort:
		res := mergeSort(books, cmpFunc)
		copy(books, res)
	case BubbleSort:
		bubbleSort(books, cmpFunc)
	}
}
