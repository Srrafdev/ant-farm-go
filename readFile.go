package main

import (
	"errors"
	"os"
	"strconv"
	"strings"
)

var (
	Start = ""
	End   = ""
	N     = 0
)

func ReadFile(fileName string) ([][]string, error) {
	content, e := os.ReadFile(fileName)
	if e != nil {
		return nil, e
	}
	if content == nil {
		e = errors.New("empty File")
		return nil, e
	}
	arr := strings.Split(string(content), "\n")
	N, e = strconv.Atoi(arr[0])
	if e != nil || N == 0 {
		e = errors.New("error in number of ants")
		return nil, e
	}

	links := [][]string{}

	for i := 0; i < len(arr); i++ {

		if arr[i] == "##start" || arr[i] == "##end" {
			if i == len(arr)-1 {
				e = errors.New("can't find start or end")
				return nil, e
			}
			if arr[i] == "##start" {
				line := strings.Split(arr[i+1], " ")
				Start = line[0]
			} else if arr[i] == "##end" {
				line := strings.Split(arr[i+1], " ")
				End = line[0]
			}
		}

		b := strings.Split(arr[i], "-")
		if len(b) == 2 {
			links = append(links, []string{b[0], b[1]})
		}
	}
	if len(links) == 0 {
		e = errors.New("there is no links")
		return nil, e
	}
	return links, nil
}
