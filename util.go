package rectpack

func quickRemove[T any](s *[]T, idx int) {
	l := len(*s)
	if l < 2 {
		*s = (*s)[:0]
		return
	}
	if idx >= l {
		return
	}
	rep := (*s)[l-1]
	(*s)[idx] = rep
	*s = (*s)[:l-1]
}
