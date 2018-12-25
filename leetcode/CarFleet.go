package leetcode

import "sort"

func carFleet(target int, position []int, speed []int) int {
	carnum := len(position)
	type car struct {
		position int
		time     float64
	}

	cars := make([]car, carnum)

	for i := 0; i < carnum; i++ {
		cars[i].position = position[i]
		cars[i].time = float64(target-position[i]) / float64(speed[i])
	}

	sort.Slice(cars, func(i, j int) bool {
		return cars[i].position > cars[j].position
	})

	ans := 0
	i := 0
	for i < carnum {
		ans += 1
		j := i + 1
		for j < carnum {
			if cars[i].time >= cars[j].time {

				cars[j].time = cars[i].time
				j += 1
			} else {
				break
			}
		}
		i = j
	}
	return ans
}
