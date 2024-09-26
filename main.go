package main

import (
	"fmt"
	"os"
	"strings"

	"box/lemin"
	box "box/parseFile"
)

func main() {
	if len(os.Args) != 2{
		fmt.Println("ERROR: inter name file")
		return
	}
	name := os.Args[1]
	graph := &lemin.Graph{}

	farms, err := box.ParseFile("exampls/"+name)
	if err != nil {
		fmt.Println("ERROR: invalid data format: ", err)
		return
	}

	for _, val := range farms.Rooms {
		err := graph.AddVertex(val)
		if err != nil{
			fmt.Println("ERROR: invalid data format: not found end: ",err)
		}
	}
	for _, val := range farms.Links {
		valsp := strings.Split(val, "-")
		err := graph.AddEdge(valsp[0], valsp[1])
		if err != nil{
			fmt.Println("ERROR: invalid data format: not found end: ",err)
			return
		}
	}

	graph.Print()
	println("****************************************")

	start := graph.GetVertex(farms.Start)
	
	paths := [][]string{}
	stack := []string{}
	graph.DFS(&paths, &stack, start, farms.End)
	if len(paths) == 0 {
		fmt.Println("ERROR: invalid data format: not found end")
		return
	}

	paths = lemin.SortPaths(paths)
	for _, v := range paths {
		fmt.Println(v)
	}
	println("****************************************")

	bestPath := lemin.Chouse(paths, farms.NumberAnts)
	println("****************************************")
	lemin.PrintAntMovements(bestPath, farms.NumberAnts, farms.Start, farms.End)
}

