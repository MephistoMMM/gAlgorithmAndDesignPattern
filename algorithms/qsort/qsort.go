package qsort

func quickSort(values []int, left, right int) {
	if left >= right {
		return
	}

	var pivot, pindex int

	mid := (left + right) / 2
	a, b, c := values[left], values[mid], values[right]
	switch {
	case a <= b && b <= c || a >= b && b >= c:
		pivot = b
		pindex = mid
	case b <= a && a <= c || b >= a && a >= c:
		pivot = a
		pindex = left
	case a <= c && c <= b || a >= c && c >= b:
		pivot = c
		pindex = right
	}

	i, j := left, right
	values[pindex] = values[i]

	for i < j {
		for i < j && values[j] >= pivot {
			j--
		}
		values[i] = values[j]

		for i < j && values[i] <= pivot {
			i++
		}
		values[j] = values[i]
	}

	values[i] = pivot
	quickSort(values, left, i-1)
	quickSort(values, i+1, right)
}

func QuickSort(values []int) {
	quickSort(values, 0, len(values)-1)
}
