package main

import (
	"./cli"
	"./utils"

	log "github.com/Sirupsen/logrus"
)

func main() {
	// check running kernel version
	if err := utils.CheckKernel(3, 10, 0); err != nil {
		log.Fatalf("Daoker should be run on Linux kernel not lower than 3.10.0, error: %v", err)
	}

	// start command line work
	cli.Run()
}
