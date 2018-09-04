package main

// 快排
func Sort(strs []string) []string {
	if len(strs) == 0 {
		return []string{}
	}
	more := []string{}
	less := []string{}
	pivot := strs[0]
	for i := 1; i < len(strs); i++ {
		str := strs[i]
		if len(pivot) > len(str) {
			less = append(less, str)
		} else {
			more = append(more, str)
		}
	}
	lessRes := []string{}
	for _, str := range sort(less) {
		lessRes = append(lessRes, str)
	}
	lessRes = append(lessRes, pivot)
	for _, str := range sort(more) {
		lessRes = append(lessRes, str)
	}
	return lessRes
}
