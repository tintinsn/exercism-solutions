package etl

import "strings"

func Transform(in map[int][]string) map[string]int {
	newScores := map[string]int{}
	for points, letters := range in {
		for _, letter := range letters {
			newScores[strings.ToLower(letter)] = points
		}
	}

	return newScores

}
