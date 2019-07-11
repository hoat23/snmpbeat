package main

import (
	"os"

	"github.com/elastic/beats/libbeat/beat"

	"github.com/hoat23/snmpbeat/beater"
)

func main() {
	err := beat.Run("snmpbeat", "", beater.New)
	if err != nil {
		os.Exit(1)
	}
}
