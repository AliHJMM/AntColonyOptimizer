package functions

// CountTunnels counts the number of tunnels leading to rooms closer to the destination
func (farm *Farm) CountTunnels(chamber *AntRoom) int {
	// Initialize the tunnel count
	tunnelCount := 0

	// Start with the first connected room
	connection := chamber.connectedRooms.firstNode

	// Iterate through all connected rooms
	for connection != nil {
		// If the connected room is closer to the destination, increment the count
		if farm.roomPaths[connection.data] < farm.roomPaths[chamber] {
				tunnelCount++
		}
		// Move to the next connection
		connection = connection.nextConnection
}

// Return the total count of tunnels leading to closer rooms
return tunnelCount
}