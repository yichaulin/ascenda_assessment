package string_list

type StringList map[string]struct{}

func New(items ...string) StringList {
	sList := StringList{}
	for _, item := range items {
		sList.Add(item)
	}

	return sList
}

func (l StringList) ToSlice() []string {
	s := make([]string, 0, len(l))
	for k := range l {
		s = append(s, k)
	}

	return s
}

func (l StringList) Add(items ...string) {
	for _, item := range items {
		l[item] = struct{}{}
	}
}

func (l StringList) Merge(from StringList) {
	for item := range from {
		l[item] = struct{}{}
	}
}

func (l StringList) Delete(item string) {
	delete(l, item)
}
