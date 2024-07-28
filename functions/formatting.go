package functions

import (
	"fmt"
	"sort"
)

func (colony *Farm) FormatAntLocations() string {
// Initialize an empty string to store formatted locations
formattedLocations := ""

// Create a new slice to hold ordered ants
orderedAnts := make([]*ColonyWorker, colony.antCount)

// Copy the workers to the new slice
copy(orderedAnts, colony.workers)

// Sort the ants by their number
sort.SliceStable(orderedAnts, func(first, second int) bool {
		return orderedAnts[first].antNumber < orderedAnts[second].antNumber
})
}