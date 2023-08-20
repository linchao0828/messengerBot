package utils

import "strings"

func MaskAddress(s string) string {
	if s == "" {
		return s
	}
	ss := strings.Split(s, ":")
	if len(ss) != 2 {
		return s
	}
	s = ss[1]
	start := s[0:4]
	end := s[len(s)-4:]
	return ss[0] + ":" + start + "..." + end
}

func MaskMobile(s string) string {
	if s == "" {
		return s
	}
	start := s[0:3]
	end := s[len(s)-4:]
	return start + "****" + end
}
