//--Summary:
//  Write a program to display server status.
//
//--Requirements:
//* Create a function to print server status, including:
//  - Number of servers
//  - Number of servers for each status (Online, Offline, Maintenance, Retired)
//* Store the existing slice of servers in a map
//* Default all of the servers to `Online`
//* Perform the following status changes and display server info:
//  - display server info
//  - change `darkstar` to `Retired`
//  - change `aiur` to `Offline`
//  - display server info
//  - change all servers to `Maintenance`
//  - display server info

package main

import (
	"fmt"
)

const (
	Online      = 0
	Offline     = 1
	Maintenance = 2
	Retired     = 3
)

// * Create a function to print server status, including:
func printStatus(servers map[string]int) {
	//   - Number of servers
	fmt.Println("Number of servers is", len(servers))

	//   - Number of servers for each status (Online, Offline, Maintenance, Retired)
	online, offline, maintenance, retired := 0, 0, 0, 0
	for _, status := range servers {
		if status == Online {
			online++
		} else if status == Offline {
			offline++
		} else if status == Maintenance {
			maintenance++
		} else if status == Retired {
			retired++
		}
	}

	fmt.Println("Number of Online status servers is", online)
	fmt.Println("Number of Offline offline status servers is", offline)
	fmt.Println("Number of Maintenance status servers is", maintenance)
	fmt.Println("Number of Retired status servers is", retired)
}

func main() {
	servers := []string{"darkstar", "aiur", "omicron", "w359", "baseline"}

	//* Store the existing slice of servers in a map
	// serversMap := make(map[string]int)
	var serversMap map[string]int

	//* Default all of the servers to `Online`
	for _, server := range servers {
		serversMap[server] = 0
		// fmt.Println(key, server)
	}

	//* Perform the following status changes and display server info:
	//  - display server info
	printStatus(serversMap)
	//  - change `darkstar` to `Retired`
	serversMap["darkstar"] = 3
	//  - change `aiur` to `Offline`
	serversMap["aiur"] = 1
	//  - display server info
	printStatus(serversMap)
	//  - change all servers to `Maintenance`
	for server := range serversMap {
		serversMap[server] = 2
	}
	//  - display server info
	printStatus(serversMap)
}
