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

// AntRoom represents a room in the ant colony
type AntRoom struct {
	connectedRooms *RoomCollection  // Rooms connected to this room
	isBeginning    bool             // True if this is the starting room
	isDestination  bool             // True if this is the destination room
	isUnoccupied   bool             // True if the room is currently empty
	xCoordinate    int              // X-coordinate of the room
	yCoordinate    int              // Y-coordinate of the room
	roomName       string           // Name of the room
	accessibility  map[string]bool  // Indicates which tunnels are accessible
	isEndpoint     bool             // True if this room is an endpoint
}