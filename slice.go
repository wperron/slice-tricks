package slice

func AppendVector[T any](a, other []T) []T {
	return append(a, other...)
}

func Copy[T any](a []T) []T {
	c := make([]T, len(a))
	copy(c, a)
	return c
}

func Cut[T any](a []T, from, to int) []T {
	return append(a[:from], a[to:]...)
}

func CutGC[T any](a []T, from, to int) []T {
	copy(a[from:], a[to:])
	for k, n := len(a)-to+from, len(a); k < n; k++ {
		var n T
		a[k] = n
	}
	a = a[:len(a)-to+from]
	return a
}

func Delete[T any](a []T, i int) []T {
	a = a[:i+copy(a[i:], a[i+1:])]
	return a
}

func DeleteGC[T any](a []T, i int) []T {
	if i < len(a)-1 {
		copy(a[i:], a[i+1:])
	}
	var n T
	a[len(a)-1] = n
	a = a[:len(a)-1]
	return a
}

func DeleteUnordered[T any](a []T, i int) []T {
	a[i] = a[len(a)-1]
	a = a[:len(a)-1]
	return a
}

func DeleteUnorderedGC[T any](a []T, i int) []T {
	a[i] = a[len(a)-1]
	var n T
	a[len(a)-1] = n
	a = a[:len(a)-1]
	return a
}

func Expand[T any](a []T, i, j int) []T {
	return append(a[:i], append(make([]T, j), a[i:]...)...)
}

func Extend[T any](a []T, j int) []T {
	return append(a, make([]T, j)...)
}

func Filter[T any](a []T, keep func(x T) bool) []T {
	n := 0
	for _, x := range a {
		if keep(x) {
			a[n] = x
			n++
		}
	}
	return a[:n]
}

func Insert[T any](a []T, at int, n T) []T {
	return append(a[:at], append([]T{n}, a[at:]...)...)
}

func InsertVector[T any](a, b []T, i int) []T {
	return append(a[:i], append(b, a[i:]...)...)
}

func Pop[T any](a []T) (T, []T) {
	return a[len(a)-1], a[:len(a)-1]
}

func Push[T any](a []T, x T) []T {
	return append(a, x)
}

func Shift[T any](a []T) (T, []T) {
	return a[0], a[1:]
}

func Unshift[T any](a []T, x T) []T {
	return append([]T{x}, a...)
}
