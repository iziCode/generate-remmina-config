package main

import (
	"fmt"
	"generate-remmina-config/internal/remmina"
	"generate-remmina-config/internal/serverscom"
)

func main() {
	generateRemminaConfigFromServerscom()
}

func generateRemminaConfigFromServerscom() {
	ServerscomToken := ""
	allServers := serverscom.GetAllServers(ServerscomToken)
	fmt.Println(len(allServers))
	fmt.Println("----------------------")
	for i, server := range allServers {
		fmt.Println(i)
		fmt.Println(server)
		fmt.Println("----------------------")
	}

	var remminaConfig []remmina.Config
	for _, server := range allServers {
		remminaConfig = append(remminaConfig, remmina.Config{
			Domain: server.Title,
			Ip:     *server.PublicIPv4Address,
		})
	}
	remmina.GenerateRemminaFiles(remminaConfig)
}

func generateRemminaConfigFromHetzner() {

}
