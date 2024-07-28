package functions

import (
	"fmt"
	"math"
	"os"
	"sort"
)

var movementOccurred = false // Global flag to track if any ant movement occurred

// printAndExit prints a message and exits the program with the given exit code
func printAndExit(message string, exitCode int) {
    fmt.Println(message)
    os.Exit(exitCode)
}

// FindShortestPath calculates the shortest path from the final room to all other rooms
func (farm *Farm) FindShortestPath() {
    stack := &RoomStack{}
    exploredRooms := make(map[*AntRoom]bool)
    seenRooms := make(map[*AntRoom]bool)

    // Start from the final room
    stack.AddRoomToStack(farm.finalRoom)
    current := stack.RemoveRoom()
    farm.roomPaths[current] = 1
    seenRooms[current] = true