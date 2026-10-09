package store

import (
	"fmt"
	"slices"
	"strings"
)

/*
Note:
In order for clients to talk to the OAuth server, the OAuth server needs to assign a unique client identifier to each client.
Static registration here is temporary, as well as keeping such data in variable instead of DB.
*/
type ClientData struct {
	ClientId     string
	ClientSecret string
	RedirectUris []string
}

// Note: In a regular OAuth system, the authorization server issues the clientID and client secret to the client software
var clientsData = []ClientData{
	{
		ClientId:     "oauth-client-1",
		ClientSecret: "oauth-client-secret-1",
		RedirectUris: []string{"http://example.com/oauth-redirect"},
	},
}

// GetClientData Gets client data by clientId from the store.
func GetClientData(clientId string) (ClientData, error) {
	i := slices.IndexFunc(clientsData, func(c ClientData) bool {
		return strings.Compare(c.ClientId, clientId) == 0
	})

	if i == -1 {
		return ClientData{}, fmt.Errorf("client not found")
	}

	return clientsData[i], nil
}
