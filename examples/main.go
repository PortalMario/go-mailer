package main

import (
	"github.com/PortalMario/go-mailer/mailer"
	"log"
)

func main() {

	mailer := mailer.NewRelay(
		"example.relay.com",
		"mysenderaddress@example.com",
		"mypasswordforrelay",
		587,
	)

	err := mailer.Send("receiveraddress@example.com", "My cool subject", "My mail body")
	if err != nil {
		log.Fatal("Error, sending mail: ", err)
	}

	log.Println("Mail sent!")
}
