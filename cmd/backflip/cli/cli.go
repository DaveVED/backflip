package cli

import (
	"flag"
	"log"

	"github.com/DaveVED/backflip/cmd/backflip/configuration"
	"github.com/DaveVED/backflip/cmd/backflip/server"
)

func CLI() {
	config := flag.String("config", "", "Backflip configuraiton file destination.")
	flag.Parse()

	if *config == "" {
		log.Println("Missing '-config' value. Please provide a config file destination for backflip.")
		return
	}

	backfliConfiguration := configuration.ParseConfigFile(*config)

	server.ServeHTTP(backfliConfiguration.TargetURL)
}
