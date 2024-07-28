package functions

import (
	"fmt"
	"os"
)

// RegisterRoom registers a room in the colony with proper error handling.
func (colony *Farm) RegisterRoom(roomName string, roomType string, xCoord int, yCoord int) {
	if _, exists := colony.roomMap[roomName]; exists {
		fmt.Printf("ERROR: invalid data format, duplicate room name '%s'\n", roomName)
		os.Exit(1) // Exit the program with an error code
	}
	switch roomType {
	case "start":
		colony.roomMap[roomName] = &AntRoom{
			roomName:       roomName,
			isBeginning:    true,
			isDestination:  false,
			isUnoccupied:   true,
			xCoordinate:    xCoord,
			yCoordinate:    yCoord,
			connectedRooms: &RoomCollection{},
			accessibility:  make(map[string]bool),
		}
		colony.initialRoom = colony.roomMap[roomName]