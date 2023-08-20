package utils

var CH = new(ch)

type ch struct{}

func (*ch) Len(s string) int {
	rt := []rune(s)
	return len(rt)
}

func (*ch) Cut(s string, start, end int) string {
	rt := []rune(s)
	return string(rt[start:end])
}

var Mark = new(mark)

type mark struct{}

func (*mark) Mark(s string, start, end int) string {
	res := make([]rune, 0)
	for i, it := range []rune(s) {
		if i >= start && i < end {
			res = append(res, []rune("*")[0])
		} else {
			res = append(res, it)
		}
	}
	return string(res)
}
