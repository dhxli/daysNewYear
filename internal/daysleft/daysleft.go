package daysleft

import "time"

func Until(from time.Time) int {
	from = time.Date(from.Year(), from.Month(), from.Day(), 0, 0, 0, 0, from.Location())
	nextNewYear := time.Date(from.Year()+1, time.January, 1, 0, 0, 0, 0, from.Location())
	return int(nextNewYear.Sub(from).Hours() / 24)
}
