package leetCode

//双指针移除字符串

func RemoveSpace(words string) string {
	slow, fast := 0, 0
	wordb := []byte(words)
	for ; fast < len(wordb); fast++ {
		if wordb[fast] != ' ' {
			wordb[slow] = wordb[fast]
			slow++
		}
	}
	return string(wordb[:slow])
}
