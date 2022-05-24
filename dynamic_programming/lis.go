package dynamic_programming

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func max_from_list(arr []int) int {

	max_num := 0

	for i, _ := range arr {
		if i+1 == len(arr) {
			break
		} else {
			max_num = max(arr[i], arr[i+1])
		}

	}

	return max_num
}

func Longest_Increasing_SubSequence(list []int) int {

	LIS := make([]int, len(list))

	for i, _ := range list {
		LIS[i] = 1
	}

	for index := 1; index < len(list); index++ {
		for lesser_index := 0; lesser_index < index; lesser_index++ {

			if list[lesser_index] < list[index] {
				LIS[index] = max(LIS[index], 1+LIS[lesser_index])
			}
		}

	}
	return max_from_list(LIS)

}
