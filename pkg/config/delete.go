package config

import (
	"log"

	"github.com/AlecAivazis/survey/v2"
	"github.com/NathanBeddoeWebDev/xoauth/pkg/db"
)

func ConfirmDelete(database *db.CredentialStore, connection string) {
	confirm := false
	prompt := &survey.Confirm{
		Message: "Are you sure you want to delete this connection?",
	}

	confirmErr := survey.AskOne(prompt, &confirm)

	if confirmErr != nil {
		log.Printf("Exiting without deleting %v\n", confirmErr)
		return
	}

	if !confirm {
		log.Printf("Exiting without deleting")
	}

	_, err := database.DeleteClient(connection)

	if err != nil {
		panic(err)
	}

	log.Println("Connection deleted")
}
