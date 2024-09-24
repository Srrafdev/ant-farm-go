package box

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type AntsFarm struct {
	Start, End string
	NumberAnts int
	Rooms      []string
	Links      []string
}

func (AF *AntsFarm) GetData(scanner *bufio.Scanner) error {
	for scanner.Scan() {
		line := scanner.Text()
		switch {
		case line == "##start":
			if !scanner.Scan() {
				return fmt.Errorf("unexpected end of file after ##start")
			}
			parts := strings.Fields(scanner.Text())
			if len(parts) < 1 {
				return fmt.Errorf("invalid start room format")
			}
			AF.Start = parts[0]
			AF.Rooms = append(AF.Rooms, parts[0])
		case line == "##end":
			if !scanner.Scan() {
				return fmt.Errorf("unexpected end of file after ##end")
			}
			parts := strings.Fields(scanner.Text())
			if len(parts) < 1 {
				return fmt.Errorf("invalid end room format")
			}
			AF.End = parts[0]
			AF.Rooms = append(AF.Rooms, parts[0])
		case strings.Contains(line, "-"):
			AF.Links = append(AF.Links, line)
		case !strings.HasPrefix(line, "#") && strings.Contains(line, " ") || !strings.HasPrefix(line, "L"):
			parts := strings.Fields(line)
			if len(parts) < 1 {
				return fmt.Errorf("invalid room format")
			}
			AF.Rooms = append(AF.Rooms, parts[0])
		}
	}

	if err := scanner.Err(); err != nil {
		return fmt.Errorf("error reading input: %v", err)
	}

	return nil
}

func ParseFile(filename string) (*AntsFarm, error) {
	contentBytes, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("error reading file %v: %v", filename, err)
	}
	content := string(contentBytes)
	scanner := bufio.NewScanner(strings.NewReader(content))

	if !scanner.Scan() {
		return nil, fmt.Errorf("file is empty")
	}
	firstLine := scanner.Text()
	numb, err := strconv.Atoi(strings.TrimSpace(firstLine))
	if err != nil {
		return nil, fmt.Errorf("invalid number of ants: %v", err)
	}

	antsFarm := &AntsFarm{NumberAnts: numb}
	if err := antsFarm.GetData(scanner); err != nil {
		return nil, err
	}

	if antsFarm.Start == "" {
		return nil, fmt.Errorf("start room not found")
	}
	if antsFarm.End == "" {
		return nil, fmt.Errorf("end room not found")
	}
	if antsFarm.NumberAnts <= 0{
		return nil, fmt.Errorf("number ants not corect")
	}

	return antsFarm, nil
}
