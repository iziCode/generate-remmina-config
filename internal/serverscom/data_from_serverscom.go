package serverscom

import (
	"context"
	"fmt"
	s "github.com/serverscom/serverscom-go-client/pkg"
)

func GetAllServers(token string) []s.Host {
	client := s.NewClient(token)
	hostCollection := client.Hosts.Collection()
	ctx := context.Background()
	firstPageHosts, err := hostCollection.List(ctx)
	if err != nil {
		fmt.Println("error:", err)
		return nil
	}

	var allHosts []s.Host
	allHosts = append(allHosts, firstPageHosts...)
	for {
		if !hostCollection.HasNextPage() {
			break
		}
		nextPageHosts, err := hostCollection.NextPage(ctx)
		if err != nil {
			fmt.Println("error:", err)
			return nil
		}
		allHosts = append(allHosts, nextPageHosts...)
	}

	return allHosts
}
