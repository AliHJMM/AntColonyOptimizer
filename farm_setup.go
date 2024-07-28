package functions

import "math"

// ColonyWorker represents an individual ant in the colony
type ColonyWorker struct {
	antNumber        int            // Unique identifier for the ant
	currentRoom      *AntRoom       // The room where the ant is currently located
	visitedRoom      map[*AntRoom]bool // Keeps track of rooms visited by the ant
	inMotion         bool           // Indicates if the ant is currently moving
	hasCompletedMove bool           // Indicates if the ant has finished its move
}