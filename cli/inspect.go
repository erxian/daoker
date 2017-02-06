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

		fmt.Println("\n")
		fmt.Println("*************************************\n")  
		fmt.Printf("ID: %s \n", scon.ID)    
		fmt.Printf("Created:  %s \n", scon.Created)
		fmt.Printf("Path: %s \n", scon.Path)
		fmt.Printf("Args: %s \n", scon.Args)
		fmt.Printf("State:  \n")
		fmt.Printf("Status: %s \n", scon.State.StateString())
		fmt.Printf("Running: %t \n", scon.State.Running)
		fmt.Printf("Paused: %t \n", scon.State.Paused)
		fmt.Printf("Restarting: %t \n", scon.State.Restarting)
		fmt.Printf("OOMKilled: %t \n", scon.State.OOMKilled)
		fmt.Printf("Dead: %t \n", scon.State.Dead)
		fmt.Printf("Pid: %d \n", scon.State.Pid)
		fmt.Printf("ExitCode: %d \n", scon.State.ExitCode)
		fmt.Printf("Error: %d \n", scon.State.Error)
		fmt.Printf("StartedAt: %s \n", scon.State.StartedAt)
		fmt.Printf("FinishedAt: %s \n", scon.State.FinishedAt)
		fmt.Printf("Image: %s \n", scon.Config.Image)
        fmt.Printf("ResolvConfPath: %s \n", scon.ResolvConfPath)
        fmt.Printf("HostnamePath: %s \n", scon.HostnamePath)
        fmt.Printf("HostsPath: %s \n", scon.HostsPath)
        fmt.Printf("LogPath: %s \n", scon.LogPath)
        fmt.Printf("Name: %s \n", scon.Name)
        fmt.Printf("RestartCount: %d \n", scon.RestartCount)
        fmt.Printf("Driver: %s \n", scon.Driver)
        fmt.Printf("MountLabel: %s \n", scon.MountLabel)
        fmt.Printf("ProcessLabel: %s \n", scon.ProcessLabel)
        fmt.Printf("AppArmorProfile: %s \n", scon.AppArmorProfile)
        fmt.Printf("Hostname : %s \n", scon.Config.Hostname)
        fmt.Printf("Domainname : %s \n", scon.Config.Domainname)
        fmt.Printf("User : %s \n", scon.Config.User)
        fmt.Printf("AttachStdin : %t \n", scon.Config.AttachStdin)
        fmt.Printf("AttachStdout : %t \n", scon.Config.AttachStdout)
        fmt.Printf("AttachStderr : %t \n", scon.Config.AttachStderr)
        fmt.Printf("Tty : %t \n", scon.Config.Tty)
        fmt.Printf("OpenStdin : %t \n", scon.Config.OpenStdin)
        fmt.Printf("StdinOnce : %t \n", scon.Config.StdinOnce)
        fmt.Printf("Env : %v \n", scon.Config.Env)
        fmt.Printf("Cmd :%s \n", scon.Config.Cmd)
        fmt.Printf("ArgsEscaped : %t \n", scon.Config.ArgsEscaped)
        fmt.Printf("Image : %s \n", scon.Config.Image)
        fmt.Printf("Volumes : %v \n", scon.Config.Volumes)
        fmt.Printf("WorkingDir : %s \n", scon.Config.WorkingDir)
        fmt.Printf("Entrypoint : %s \n", scon.Config.Entrypoint)
        fmt.Printf("OnBuild : %v \n", scon.Config.OnBuild)
        fmt.Printf("Labels : %v \n", scon.Config.Labels)

        if scon.Networks == nil{
        fmt.Println("no network info \n")
        }

        networks, ok := scon.Networks["bridge"]
        /* 如果 ok 是 true, 则存在，否则不存在 */
        if(ok){
          fmt.Println("NetworkID : %s\n", networks.NetworkID)  
          }else {
          fmt.Println("unknown \n") 
        }

        //fmt.Printf("NetworkID : %s \n", scon.Networks["bridge"].NetworkID) 
		fmt.Println("*************************************\n\n")  

	}

}

