package box

import (
	"fmt"
	"os"
	"strings"
)

type AntsFarm struct {
	Start, End string
	NumberAnts int
	Rooms      []string
	Links      [][]string
}

func (AF *AntsFarm) GetStartEnd(farms []string) {
	var start, end string
	for i, val := range farms {
		if i+1 < len(farms) && val == "##start" {
			start = farms[i+1]
		} else if i+1 <= len(farms) && val == "##end" {
			end = farms[i+1]
		}
	}
	AF.Start = start
	AF.End = end
}

// func (AF *AntsFarm) GetRooms(farms []string) {
// 	var rooms []string
// 	for _, val := range farms {
// 		flag := false
// 	}
// }

// func clerComanter(farms []string) []string {
// }

func ParseFile(File string) {
	contentbyte, err := os.ReadFile(File)
	if err != nil {
		fmt.Printf("\nerror read file %v : %v", File, err)
		return
	}
	content := string(contentbyte)
	contentSplite := strings.Split(content, "##")

		fmt.Print(contentSplite[0])
	

	//antsfarm := &AntsFarm{}

	//antsfarm.GetStartEnd(contentSplite)
}
