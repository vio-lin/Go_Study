package bubblesort

// Bubblesort 用于冒泡排序
func Bubblesort(values []int) {
	flag := true
	for i := 0; i < len(values)-1; i++ {
		flag = true
		for j := 0; j < len(values)-i-1; j++ {
			if values[j] > values[j+1] {
				values[j], values[j+1] = values[j+1], values[j]
				flag = false
			} // end if
		} // end for j = ...
		if flag == true {
			break
		}
	}
}
