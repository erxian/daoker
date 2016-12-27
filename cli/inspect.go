package cli

import (
	"fmt"

	log "github.com/Sirupsen/logrus"
	"../docker"
	"github.com/codegangsta/cli"
)

//inspectContainer inspects container's configuration and serializes it as json.
func inspectContainer(c *cli.Context) {
	if len(c.Args()) != 1 {
		log.Fatalf("Inspect command takes exact one argument. See '%s inspect --help'.", c.App.Name)
	}

	IDOrName := c.Args()[0]
	container, err := docker.GetContainer(IDOrName)
	if err != nil {
		log.Fatal(err.Error())
	}

	containers, err := docker.Containers()  

	if err != nil {
		log.Fatal(err)
	}

	for _, scon := range containers {   
		ID := scon.ID
		if ID != container.ID {
			continue
		}
		fmt.Println("-------------------------------------\n")  
		fmt.Printf("ID:%s \n", scon.ID)    
		fmt.Printf("Created:  %s \n", scon.Created)
		fmt.Printf("Path:%s \n", scon.Path)
		fmt.Printf("Args:%s \n", scon.Args)
		fmt.Printf("State:%s \n", scon.State.StateString())
		fmt.Println("-------------------------------------\n\n")  

	}

}

