package leetmerge

func mergeAlternately(word1 string, word2 string) string {
	len1 := len(word1)
	len2 := len(word2)
	res := make([]byte, len1+len2)
	onePos := 0
	twoPos := 0
	next := 1
	for i := 0; i < len(res); i++ {
		if next == 1 {
			res[i] = word1[onePos]
			onePos++
			if twoPos < len2 {
				next = 2
			}
		} else {
			res[i] = word2[twoPos]
			twoPos++
			if onePos < len1 {
				next = 1
			}
		}
	}
	return string(res)
}
