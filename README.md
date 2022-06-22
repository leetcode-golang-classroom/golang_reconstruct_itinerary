# golang_reconstruct_itinerary

You are given a list of airline `tickets` where `tickets[i] = [fromi, toi]` represent the departure and the arrival airports of one flight. Reconstruct the itinerary in order and return it.

All of the tickets belong to a man who departs from `"JFK"`, thus, the itinerary must begin with `"JFK"`. If there are multiple valid itineraries, you should return the itinerary that has the smallest lexical order when read as a single string.

- For example, the itinerary `["JFK", "LGA"]` has a smaller lexical order than `["JFK", "LGB"]`.

You may assume all tickets form at least one valid itinerary. You must use all the tickets once and only once.

## Examples

**Example 1:**

![https://assets.leetcode.com/uploads/2021/03/14/itinerary1-graph.jpg](https://assets.leetcode.com/uploads/2021/03/14/itinerary1-graph.jpg)

```
Input: tickets = [["MUC","LHR"],["JFK","MUC"],["SFO","SJC"],["LHR","SFO"]]
Output: ["JFK","MUC","LHR","SFO","SJC"]

```

**Example 2:**

![https://assets.leetcode.com/uploads/2021/03/14/itinerary2-graph.jpg](https://assets.leetcode.com/uploads/2021/03/14/itinerary2-graph.jpg)

```
Input: tickets = [["JFK","SFO"],["JFK","ATL"],["SFO","ATL"],["ATL","JFK"],["ATL","SFO"]]
Output: ["JFK","ATL","JFK","SFO","ATL","SFO"]
Explanation: Another possible reconstruction is ["JFK","SFO","ATL","JFK","ATL","SFO"] but it is larger in lexical order.

```

**Constraints:**

- `1 <= tickets.length <= 300`
- `tickets[i].length == 2`
- `fromi.length == 3`
- `toi.length == 3`
- $`from_i`$ and $`to_i`$ consist of uppercase English letters.
- $`from_i$ != $to_i$`

## 解析

題目給定一個整數矩陣  tickets , 其中每個 entry ticket[i] = [$location_1$, $location_2$] 代表

$location_1$ 到 $location_2$ 有一個 path 可以經過

要求寫一個演算法

找出從 “JFK” 出發按照給定的 tickets 以及地點字母排序 走訪完所有 path 的一個可行順序

首先是這次的 path 是有順序的

所以關鍵是要透過 tickets 做出 adjacency list

然後這些 adjacency list 需要按照字母順序排列

一個可行的作法是先把 tickets 先做字母排序

然後在照順序做 adjacency list

然後依序從 “JFK” 做 DFS

然後每次經過一個地點 就把原本的 adjacency list 的 path消去一個

![](https://i.imgur.com/nK2J1eh.png)

## 程式碼
```go
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

```
## 困難點

1. 理解如何達成有序的找到 Location
2. 對 DFS 需要去理解

## Solve Point

- [x]  需要知道如何找到鄰近的 Location
- [x]  建立 adjacency list 時需要透過 sort 去處理