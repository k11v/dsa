package main

// LeetCode

func maxDistToClosest(seats []int) int {
	maxDistance := 1
	l := -1

	for i := 0; i < len(seats); i++ {
		if seats[i] != 1 {
			continue
		}

		r := i
		if l == -1 {
			maxDistance = max(r, maxDistance)
		} else {
			maxDistance = max((r-l)/2, maxDistance)
		}

		l = r
	}
	maxDistance = max(len(seats)-1-l, maxDistance)

	return maxDistance
}
