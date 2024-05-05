package insertion

import (
	"cmp"
)

func Insertion[T cmp.Ordered](s []T) {
	for i := 1; i < len(s); i++ {
		k := s[i]

		j := i - 1
		for ; j >= 0 && s[j] > k; j-- {
			s[j+1] = s[j]
		}

		s[j+1] = k
	}
}
