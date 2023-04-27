package main

import "fmt"

const (
	Online = 0
	Offline = 1
	Maintenance = 2
	Retired = 3
)

func printServerStatus(servers map[string]int)  {
	fmt.Println("\nThere are", len(servers), "servers")

	stats := make(map[int]int)
	for _, status := range servers {
		switch status {
		case Online:
			stats[Online] += 1
		case Offline:
			stats[Offline] += 1
		case Maintenance:
			stats[Maintenance] += 1
		case Retired:
			stats[Retired] += 1
		default:
			panic("unhandled server status")
		}
	}
	fmt.Println(stats[Online], "servers are online")
	fmt.Println(stats[Offline], "servers are offline")
	fmt.Println(stats[Maintenance], "servers are undergoing maintenance")
	fmt.Println(stats[Retired], "servers are retired")
}

func main() {
	shoppingList := make(map[string]int)
	shoppingList["eggs"] = 12
	shoppingList["milk"] = 1
	shoppingList["bread"] = 2
	fmt.Println(shoppingList)

	shoppingList["milk"] += 1
	fmt.Println(shoppingList)

	shoppingList["cheese"] = 3
	fmt.Println(shoppingList)

	delete(shoppingList, "cheese")
	fmt.Println("Cheese deleted from shopping list:", shoppingList)

	fmt.Println("need", shoppingList["eggs"], "eggs")

	cereal, found := shoppingList["cereal"]
	fmt.Println(cereal, found)
	fmt.Println("need cereal?")
	if found {
		fmt.Println("nop")
		} else {
		fmt.Println("yup")
	}

	totalItems := 0
	for _, amount := range shoppingList {
		totalItems += amount
	}
	fmt.Println("There are", totalItems, "shopping lists")

	// Exercise
	servers := []string{"darkstar", "aiur", "omicron", "w359", "baseline"}
	serverStatus := make(map[string]int)

	for _, server := range servers {
		serverStatus[server]  = Online
	}

	printServerStatus(serverStatus)

	serverStatus["darkstar"] = Retired
	serverStatus["aiur"] = Offline

	for server := range serverStatus {
		serverStatus[server] = Maintenance
	}

	printServerStatus(serverStatus)
}
