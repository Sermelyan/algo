package selection

import (
	"cmp"
)

func Selection[T cmp.Ordered](s []T) {
	size := len(s)
	for i := 0; i < size-1; i++ {
		min := i
		for j := i; j < size; j++ {
			if s[j] < s[min] {
				min = j
			}
		}
		s[i], s[min] = s[min], s[i]
	}
}
