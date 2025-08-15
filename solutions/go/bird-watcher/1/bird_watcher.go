package birdwatcher

// TotalBirdCount return the total bird count by summing
// the individual day's counts.
func TotalBirdCount(birdsPerDay []int) int {
	total := 0
    for _, birds := range birdsPerDay {
        total += birds
    }
    return total
}

// BirdsInWeek returns the total bird count by summing
// only the items belonging to the given week.
func BirdsInWeek(birdsPerDay []int, week int) int {
    startIndex := 7 * (week -1)
    endIndex := 7 * week
    if week < 1 || startIndex >= len(birdsPerDay) {
        return 0
    }
    if endIndex > len(birdsPerDay){
        endIndex = len(birdsPerDay)
    }
	total := 0
    for _, birds := range birdsPerDay[startIndex:endIndex] {
        total += birds
    }
    return total
}

// FixBirdCountLog returns the bird counts after correcting
// the bird counts for alternate days.
func FixBirdCountLog(birdsPerDay []int) []int {
	for idx := range birdsPerDay {
        if idx%2 == 0 {
            birdsPerDay[idx]++
        }
    }
    return birdsPerDay
}
