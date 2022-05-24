package dynamic_programming

func Sub_array_for_given_sum(list []int, sum int) []int {

	lft_ptr, rt_ptr, current_sum := 0, 0, 0

	for current_sum != sum && lft_ptr <= len(list)-1 && rt_ptr <= len(list)-1 {

		if current_sum < sum {
			current_sum = current_sum + list[rt_ptr]
			rt_ptr = rt_ptr + 1

		} else if current_sum > sum {
			current_sum = current_sum - list[lft_ptr]
			lft_ptr = lft_ptr + 1
		}

	}

	if current_sum == sum {
		return list[lft_ptr:rt_ptr]
	}

	return []int{}

}
