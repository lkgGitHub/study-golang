package main


// 遍历判断两个map相等，且避免"元素存在但值为零"
func equal(x, y map[string]int) bool {
	if len(x) != len(y) {
		return false
	}
	for k, xv := range x{
		if yv, ok := y[k]; !ok || xv != yv  {
			return false
		}
	}
	return true
}
