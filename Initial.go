package main

import (
	"fmt"
	"sort"
)

type Server struct {
	Name   string
	Region string
	Load   float64
	Online bool
}

type ServerManager struct {
	servers []Server
}

func (m *ServerManager) AddServer(server Server) {
	m.servers = append(m.servers, server)
}

func (m *ServerManager) SortByLoad() {
	sort.Slice(m.servers, func(i, j int) bool {
		return m.servers[i].Load < m.servers[j].Load
	})
}

func (m ServerManager) ShowReport() {
	online := 0
	totalLoad := 0.0

	fmt.Println("Server Manager")
	fmt.Println("==============")

	for _, server := range m.servers {
		status := "Offline"
		if server.Online {
			status = "Online"
			online++
			totalLoad += server.Load
		}

		fmt.Printf(
			"%s | %s | Load: %.1f%% | %s\n",
			server.Name,
			server.Region,
			server.Load,
			status,
		)
	}

	fmt.Println()
	fmt.Printf("Total servers: %d\n", len(m.servers))
	fmt.Printf("Online servers: %d\n", online)

	if online > 0 {
		fmt.Printf("Average load: %.1f%%\n", totalLoad/float64(online))
	}
}

func main() {
	manager := ServerManager{}

	manager.AddServer(Server{"api-01", "EU", 42.5, true})
	manager.AddServer(Server{"api-02", "US", 71.3, true})
	manager.AddServer(Server{"db-01", "EU", 88.7, true})
	manager.AddServer(Server{"cache-01", "ASIA", 15.2, false})
	manager.AddServer(Server{"worker-01", "EU", 56.8, true})

	manager.SortByLoad()
	manager.ShowReport()
}