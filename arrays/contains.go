package arrays

//ContainsString 字符串src是否包含val
func ContainsString(src []string, val string) bool {
	if len(src) == 0 {
		return false
	}

	for i := 0; i < len(src); i++ {
		if src[i] == val {
			return true
		}
	}
	return false
}
