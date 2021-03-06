package main

import (
	"fmt"
	"os"

	networkmanager "github.com/phoracek/networkmanager-go/src"
)

func main() {
	connectionIDToRemove := os.Args[1]

	client, err := networkmanager.NewClient()
	if err != nil {
		panic(err)
	}
	defer client.Close()

	// remove all connections with given ID
	for connection := findConnection(client, connectionIDToRemove); connection != nil; connection = findConnection(client, connectionIDToRemove) {
		err := connection.Delete()
		if err != nil {
			fmt.Fprintln(os.Stderr, "Failed", err)
			os.Exit(1)
		}
	}
}

func findConnection(client *networkmanager.Client, connectionID string) *networkmanager.Connection {
	connections, err := client.ListConnections()
	if err != nil {
		panic(err)
	}

	for _, connection := range connections {
		settings, _ := connection.GetSettings()
		if settings["connection"]["id"] == connectionID {
			return connection
		}
	}

	return nil
}
