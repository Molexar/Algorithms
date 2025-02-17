package internal

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ReadMatrix(filename string) ([][]float64, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("ошибка при открытии файла: %w", err)
	}
	defer file.Close()

	var matrix [][]float64
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Fields(line)
		row := make([]float64, len(fields))

		for i, field := range fields {
			val, err := strconv.ParseFloat(field, 64)
			if err != nil {
				return nil, fmt.Errorf("ошибка при преобразовании числа: %w", err)
			}
			row[i] = val
		}

		matrix = append(matrix, row)
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("ошибка при чтении файла: %w", err)
	}

	return matrix, nil
}
