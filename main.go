package main

import (
	log "github.com/Sirupsen/logrus"
	"github.com/pmyjavec/stashify/stashify"
)

func main() {
	log.SetLevel(log.DebugLevel)
	stashify.Execute()
}
