package main

import (
	"log"
	"os"
	"strconv"
)

// Config encapsulates the fields used to configure the service.
type Config struct {
	WithRelay                bool
	RelayIngressAddress      string
	RelayEgressAddress       string
	EgressDestinationAddress string
}

// NewConfig constructs a new Configuration for the service.
func NewConfig() Config {
	withRelay := false
	withRelayEnv := os.Getenv("WITH_RELAY")
	if withRelayEnv != "" {
		var err error
		withRelay, err = strconv.ParseBool(withRelayEnv)
		if err != nil {
			log.Printf("Error parsing WITH_RELAY: %v", err)
		}
	}

	return Config{
		WithRelay:                withRelay,
		RelayIngressAddress:      os.Getenv("RELAY_INGRESS_ADDRESS"),
		RelayEgressAddress:       os.Getenv("RELAY_EGRESS_ADDRESS"),
		EgressDestinationAddress: os.Getenv("EGRESS_DESTINATION_ADDRESS"),
	}
}
