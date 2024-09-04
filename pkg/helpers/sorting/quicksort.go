package sorting

func QuickSort[T int | int64 | int32 | float64 | float32 | string](partition []T) []T {
	switch len(partition) {
	case 0, 1:
		return partition
	case 2:
		if partition[0] > partition[1] {
			temp := partition[0]
			partition[0] = partition[1]
			partition[1] = temp
		}

		return partition
	}

	pivot := partition[len(partition)-1]
	var lhs, rhs []T
	var pivotOccurrences int

	for _, el := range partition {
		if el < pivot {
			lhs = append(lhs, el)
		}

		if el > pivot {
			rhs = append(rhs, el)
		}

		// In case the pivot appears more than once,
		// keep track of how many times it appears
		if el == pivot {
			pivotOccurrences += 1
		}
	}

	lhs = QuickSort(lhs)
	rhs = QuickSort(rhs)

	for range pivotOccurrences {
		// this is equivalent to prepending the pivot to the rhs
		lhs = append(lhs, pivot)
	}

	return append(lhs, rhs...)
}
