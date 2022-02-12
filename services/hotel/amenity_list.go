package hotel

type amenityList map[string]struct{}

func (l amenityList) ToStringSlice() []string {
	s := make([]string, 0, len(l))
	for k := range l {
		s = append(s, k)
	}

	return s
}

func (l amenityList) Add(s string) {
	l[s] = struct{}{}
}

func (l amenityList) Merge(from amenityList) {
	for k := range from {
		l[k] = struct{}{}
	}
}
