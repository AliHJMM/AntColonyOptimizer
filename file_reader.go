package functions

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ReadAndParseFile(inputPath string, colony *Farm) (bool, string) {
 // Initialize variables
 firstline := false
 inputFile, err := os.Open(inputPath)
 if err != nil {
	 fmt.Println("ERROR: unable to open file", err)
	 os.Exit(1)
 }
 defer inputFile.Close()
 inputInfo := ""
 lineIndex := 0
 fileScanner := bufio.NewScanner(inputFile)
 isBeginningRoom := false
 isFinalRoom := false
 AntsCount := 0
 startRoomSet := false
 endRoomSet := false

 for fileScanner.Scan() {
	line := fileScanner.Text()
	inputInfo += line + "\n"
	lineIndex++

	// Handle comments and special commands
	if strings.HasPrefix(line, "#") {
		if line == "##start" {
			if startRoomSet {
				fmt.Printf("ERROR: invalid data format, multiple start rooms defined (%s)\n", line)
				os.Exit(1)
			}
			startRoomSet = true
			isBeginningRoom = true
			continue
		} else if line == "##end" {
			if endRoomSet {
				fmt.Printf("ERROR: invalid data format, multiple end rooms defined (%s)\n", line)
				os.Exit(1)
			}
			endRoomSet = true
			isFinalRoom = true
			continue
		} else {
			continue
		}
	}
	 // Parse the first line (number of ants)
	 if !firstline {
		roomFields := strings.Fields(line)
		if len(roomFields) != 1 {
			fmt.Printf("⛔Invalid data format detected⛔ (%s)\n", line)
			os.Exit(1)
		}
		AntsCount, err = strconv.Atoi(roomFields[0])
		if err != nil || AntsCount <= 0 {
			fmt.Printf("ERROR: invalid data format, invalid number of ants (%s)\n", line)
			os.Exit(1)
		}
		firstline = true
		continue
	}
 }
}