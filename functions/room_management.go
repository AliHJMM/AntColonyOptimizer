package functions

// CountTunnels counts the number of tunnels leading to rooms closer to the destination
func (farm *Farm) CountTunnels(chamber *AntRoom) int {
	// Initialize the tunnel count
	tunnelCount := 0

	// Start with the first connected room
	connection := chamber.connectedRooms.firstNode