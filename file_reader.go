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

 }
}