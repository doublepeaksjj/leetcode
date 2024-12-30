package arrows

import (
	"sort"
)

func MinArrows(points [][]int) int {
	sort.Slice(points, func(i, j int) bool {
		return points[i][1] < points[j][1]
	})
	curArrow := points[0][1]
	numArrows := 1
	for _, p := range points {
		if curArrow >= p[0] && curArrow <= p[1] {
			// This balloon is taken out by the current arrow
			continue
		}
		// Start a new arrow for this balloon
		curArrow = p[1]
		numArrows++
	}
	return numArrows
}
