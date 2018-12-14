package leetcode

func ReachNumber(target int) int {
	if target < 0 {
		target = 0 - target
	}
	step := 1
	sum := 0

	sum += step

	for sum < target {
		step += 1
		sum += step
	}

	for (sum-target)%2 != 0 {
		step += 1
		sum += step
	}

	return step
}
