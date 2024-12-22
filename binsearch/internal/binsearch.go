package internal

import (
	"cmp"
	"sort"
)

type Sortable interface {
	cmp.Ordered
	comparable
}

type result struct {
	targetIndex int
	index       int
}

func Search[T Sortable](inputs, targets []T, maxWorkers int) []int {

	sort.Slice(targets, func(i, j int) bool {
		return targets[i] < targets[j]
	})

	results := make(chan result, len(targets))
	queue := make(chan func(), maxWorkers)

	go binSearchMany(inputs, targets, 0, len(inputs)-1, 0, len(targets)-1, results, queue)

	go func() {
		for task := range queue {
			task()
		}
	}()

	indexes := make([]int, len(targets))
	i := 0
	for res := range results {
		indexes[res.targetIndex] = res.index
		i += 1

		if i == len(targets) {
			close(results)
			close(queue)
		}
	}
	return indexes
}

func binSearchMany[T Sortable](inputs, targets []T, left, right, leftTarget, rightTarget int, results chan<- result, queue chan func()) {

	if leftTarget == rightTarget {
		target := targets[leftTarget]

		queue <- func() {
			results <- result{
				index:       binSearch(inputs, target, left, right),
				targetIndex: leftTarget,
			}
		}
		return
	}

	if (rightTarget - leftTarget) == 1 {

		go func() {
			queue <- func() {
				results <- result{
					index:       binSearch(inputs, targets[leftTarget], left, right),
					targetIndex: leftTarget,
				}
				<-queue
			}
		}()

		results <- result{
			index:       binSearch(inputs, targets[rightTarget], left, right),
			targetIndex: rightTarget,
		}
		return
	}

	midTarget := (leftTarget + rightTarget) / 2
	mid := (left + right) / 2
	if targets[midTarget] < inputs[mid] {

		go func() {
			queue <- func() { binSearchMany(inputs, targets, left, mid-1, leftTarget, midTarget, results, queue) }
		}()
		binSearchMany(inputs, targets, left, right, midTarget+1, rightTarget, results, queue)
		return
	}

	go func() {
		queue <- func() { binSearchMany(inputs, targets, left, right, leftTarget, midTarget, results, queue) }
	}()
	binSearchMany(inputs, targets, mid, right, midTarget+1, rightTarget, results, queue)
}

func binSearch[T Sortable](inputs []T, target T, left, right int) int {
	if left > right {
		return -1
	}

	if left == right {
		if inputs[left] == target {
			return left
		} else {
			return -1
		}
	}

	mid := (left + right) / 2
	if inputs[mid] == target {
		return mid
	} else if inputs[mid] <= target {
		return binSearch(inputs, target, mid+1, right)
	} else {
		return binSearch(inputs, target, left, mid-1)
	}
}
