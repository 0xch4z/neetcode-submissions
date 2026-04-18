func twoSum(nums []int, target int) []int {
    targets := make(map[int]int) // key: target augend, value: addend index

	for i, n := range nums {
		if j, ok := targets[n]; ok {
			return []int{j, i}
		}

		augend := target - n

		targets[augend] = i
	}

	return []int{}
}
