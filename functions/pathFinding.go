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

	 // Breadth-first search
	 for current != nil {
        neighborNode := current.connectedRooms.firstNode
        for neighborNode != nil {
            if !seenRooms[neighborNode.data] {
                if !neighborNode.data.isBeginning {
                    stack.AddRoomToStack(neighborNode.data)
                }
                seenRooms[neighborNode.data] = true
                if farm.roomPaths[current]+1 < farm.roomPaths[neighborNode.data] {
                    farm.roomPaths[neighborNode.data] = farm.roomPaths[current] + 1
                }
            }
            if neighborNode.data == farm.initialRoom {
                exploredRooms[neighborNode.data] = true
            }
            neighborNode = neighborNode.nextConnection
        }
        exploredRooms[current] = true
        current = stack.RemoveRoom()
    }
	  // Check if initial room was explored
	  if !exploredRooms[farm.initialRoom] {
        printAndExit("⛔ Initial Room Not Explored⛔", 0)
    }
}