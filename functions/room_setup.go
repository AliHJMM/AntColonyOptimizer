package functions

// CheckDistinctRoomCoordinates verifies that all rooms have unique coordinates
func (colony Farm) CheckDistinctRoomCoordinates() bool {
	// Iterate through all rooms in the colony
	for _, currentChamber := range colony.roomMap {
			// Compare each room with every other room
			for _, nextChamber := range colony.roomMap {
					// Skip comparing a room with itself
					if currentChamber.roomName != nextChamber.roomName {
							// Check if the coordinates are the same
							if currentChamber.xCoordinate == nextChamber.xCoordinate && currentChamber.yCoordinate == nextChamber.yCoordinate {
									// If coordinates match, rooms are not distinct
									return false
							}
					}
			}
	}
	// All rooms have distinct coordinates
	return true
}