package config

import (
	"errors"
	"log"

	"github.com/NathanBeddoeWebDev/xoauth/pkg/db"
	"github.com/spf13/cobra"
)

func ValidateSecretCmdArgs(cmd *cobra.Command, args []string) error {
	if len(args) < 1 {
		return errors.New("please supply a client name, e.g, `xero`")
	}

	if len(args) < 2 {
		return errors.New("please supply a client secret, e.g, `secret`")
	}

	return nil
}

func UpdateSecret(database *db.CredentialStore, clientName string, clientSecret string) {
	allClients, clientsErr := database.GetClients()

	if clientsErr != nil {
		log.Fatal(clientsErr)
	}

	client, clientErr := database.GetClientWithoutSecret(allClients, clientName)

	if clientErr != nil {
		log.Fatal(clientErr)
	}

	_, secretErr := database.SetClientSecret(client.Alias, clientSecret)

	if secretErr != nil {
		log.Fatal(secretErr)
	}

	log.Printf("Updated client secret for %s\n", client.Alias)
}
