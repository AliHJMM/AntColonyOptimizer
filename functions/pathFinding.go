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

// AreAllAntsAtFinalRoom checks if all ants have reached the final room
func (colony *Farm) AreAllAntsAtFinalRoom() bool {
    for _, worker := range colony.workers {
        if worker.currentRoom != colony.finalRoom {
            return false
        }
    }
    return true
}

// AdvanceAntsInFarm moves ants through the farm
func (colony *Farm) AdvanceAntsInFarm(toggle bool) {
    activeWorkers := make([]*ColonyWorker, colony.antCount)
    copy(activeWorkers, colony.workers)

    // Sort workers based on the number of tunnels in their current room
    sort.SliceStable(activeWorkers, func(workerIdx1, workerIdx2 int) bool {
        return colony.CountTunnels(activeWorkers[workerIdx1].currentRoom) < colony.CountTunnels(activeWorkers[workerIdx2].currentRoom)
    })