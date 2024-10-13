package common

func trimLeadingZeros(s []int) []int {
	for len(s)>0 && s[0] == 0 {
		s = s[1:]
	}
	return s
}