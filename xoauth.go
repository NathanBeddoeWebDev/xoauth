package main

import (
	"log"

	"github.com/NathanBeddoeWebDev/xoauth/cmd"
)

func main() {
	// No timestamps
	log.SetFlags(0)

	cmd.Execute()
}
