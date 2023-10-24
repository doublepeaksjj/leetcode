package greatestcommondivisor

func gcdOfStrings(str1 string, str2 string) string {
	len1 := len(str1)
	len2 := len(str2)
	if len2 > len1 {
		s := str1
		str1 = str2
		str2 = s
	}

	for divisor := len1; divisor >= 1; divisor-- {
		lenDividesExactly := len1%divisor == 0 && len2%divisor == 0
		if !lenDividesExactly {
			continue
		}
		subStr := str1[0:divisor]
		if doesDivide(str1, subStr) && doesDivide(str2, subStr) {
			return subStr
		}
	}
	return ""
}

func doesDivide(s string, sub string) bool {
	for n := 0; n < len(s); n += len(sub) {
		if s[n:n+len(sub)] != sub {
			return false
		}
	}
	return true
}
