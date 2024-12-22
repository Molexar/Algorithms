package main

import (
	"molexar.org/library/internal"
)

func main() {
	books := []internal.Book{
		{
			Genre:  "Фентези",
			Author: "Зубов К.",
			Title:  "Как я строил магическую империю",
		},
		{
			Genre:  "Фентези",
			Author: "Зубов К.",
			Title:  "Последний попаданец",
		},
		{
			Genre:  "Фентези",
			Author: "Зубов К.",
			Title:  "Системный приход: Орки под Москвой",
		},
		{
			Genre:  "Фентези",
			Author: "Кольцов И.",
			Title:  "Наследник рода Раджат - 17. Финал",
		},
		{
			Genre:  "Триллер",
			Author: "Я",
			Title:  "Устал копипастить",
		},
		{
			Genre:  "Триллер",
			Author: "Я",
			Title:  "Устал копипастить",
		},
		{
			Genre:  "Триллер",
			Author: "Я",
			Title:  "Устал копипастить",
		},
		{
			Genre:  "Триллер",
			Author: "Я",
			Title:  "Устал копипастить",
		},
		{
			Genre:  "Триллер",
			Author: "Я",
			Title:  "Устал копипастить",
		},
		{
			Genre:  "Триллер",
			Author: "Я",
			Title:  "Устал копипастить",
		},
	}

	internal.Sort(books, func(x, y internal.Book) bool {
		return !func(x, y internal.Book) bool {
			if x.Genre < y.Genre {
				return true
			}
			if x.Genre == y.Genre {
				if x.Author < y.Author {
					return true
				}
				if x.Author == y.Author {
					return x.Title < y.Title
				}
			}
			return false
		}(x, y)
	}, internal.Opts{
		Algorithm: internal.QuickSort,
	})

	library := internal.SetLibrary(books, 3)
	internal.PrintLibrary(library)
}
