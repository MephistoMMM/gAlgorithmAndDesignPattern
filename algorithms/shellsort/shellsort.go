// Package shellsort provide algorithm of shellsort
//
// Author: Mephis Pheies <mephistommm@gmail.com>
package shellsort

func shellInsert(values []int, offset int) {
	length := len(values)

	for i := offset; i < length; i += offset {
		if values[i] < values[i-offset] {
			temp := values[i]
			j := i
			for ; j >= offset && temp < values[j-offset]; j -= offset {
				values[j] = values[j-offset]
			}

			values[j] = temp
		}
	}
}

func ShellSort(values []int) {
	for offset := len(values) / 2; offset > 0; offset /= 2 {
		shellInsert(values, offset)
	}
}
