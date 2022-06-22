package sol

import (
	"sort"
	"strings"
)

type Locations []string

func findItinerary(tickets [][]string) []string {
	sortedTickets := SortInput(tickets)
	adjacencyMap := make(map[string]Locations)
	for _, ticket := range sortedTickets {
		adjacencyMap[ticket[0]] = append(adjacencyMap[ticket[0]], ticket[1])
	}
	result := []string{"JFK"}
	var dfs func(location string) bool
	dfs = func(location string) bool {
		if len(result) == len(tickets)+1 {
			return true
		}
		adjacencyList, ok := adjacencyMap[location]
		if !ok {
			return false
		}
		temp := make([]string, len(adjacencyList))
		copy(temp, adjacencyList)
		for _, loc := range adjacencyList {
			result = append(result, loc)
			adjacencyList = adjacencyList[1:]
			adjacencyMap[location] = adjacencyList
			if dfs(loc) {
				return true
			}
			result = result[:len(result)-1]
			adjacencyList = append(adjacencyList, loc)
			adjacencyMap[location] = adjacencyList
		}
		return false
	}
	dfs("JFK")
	return result
}

func SortInput(tickets [][]string) [][]string {
	temp := make([]string, len(tickets))
	result := make([][]string, len(tickets))
	for idx, ticket := range tickets {
		temp[idx] = ticket[0] + "," + ticket[1]
	}
	sort.Strings(temp)
	for idx, item := range temp {
		result[idx] = strings.Split(item, ",")
	}
	return result
}
