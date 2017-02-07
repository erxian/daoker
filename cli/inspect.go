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

		fmt.Printf("Bridge : %s \n", scon.NetworkSettings.Bridge) 
		fmt.Printf("SandboxID : %s \n", scon.NetworkSettings.SandboxID) 
		fmt.Printf("HairpinMode : %t \n", scon.NetworkSettings.HairpinMode)
		fmt.Printf("LinkLocalIPv6Address : %s \n", scon.NetworkSettings.LinkLocalIPv6Address)
		fmt.Printf("LinkLocalIPv6PrefixLen : %d \n", scon.NetworkSettings.LinkLocalIPv6PrefixLen)
 
		fmt.Printf("Ports : %v \n", scon.NetworkSettings.Ports)
		fmt.Printf("SandboxKey : %s \n", scon.NetworkSettings.SandboxKey) 
		fmt.Printf("SecondaryIPAddresses : %v \n", scon.NetworkSettings.SecondaryIPAddresses) 
		fmt.Printf("SecondaryIPv6Addresses : %v \n", scon.NetworkSettings.SecondaryIPv6Addresses) 

		if networks, ok := scon.NetworkSettings.Networks["bridge"]; ok {

                        fmt.Printf("IPAMConfig : %v \n", networks.IPAMConfig)
                        fmt.Printf("Links : %s \n", networks.Links)
                        fmt.Printf("Aliases : %s \n", networks.Aliases)
                        fmt.Printf("NetworkID : %s \n", networks.NetworkID)
                        fmt.Printf("EndpointID : %s \n", networks.EndpointID)
                        fmt.Printf("Gateway : %s \n", networks.Gateway)
                        fmt.Printf("IPAddress : %s \n", networks.IPAddress)
                        fmt.Printf("IPPrefixLen : %d \n", networks.IPPrefixLen)
                        fmt.Printf("IPv6Gateway : %s \n", networks.IPv6Gateway)
                        fmt.Printf("GlobalIPv6Address : %s \n", networks.GlobalIPv6Address)
                        fmt.Printf("GlobalIPv6PrefixLen : %d \n", networks.GlobalIPv6PrefixLen)
                        fmt.Printf("MacAddress : %s \n", networks.MacAddress)

                        }else{

                        fmt.Printf("EndpointID : %s \n", scon.NetworkSettings.DefaultNetworkSettings.EndpointID)
                        fmt.Printf("Gateway : %s \n", scon.NetworkSettings.DefaultNetworkSettings.Gateway)
                        fmt.Printf("IPAddress : %s \n", scon.NetworkSettings.DefaultNetworkSettings.IPAddress)
                        fmt.Printf("IPPrefixLen : %d \n", scon.NetworkSettings.DefaultNetworkSettings.IPPrefixLen)
                        fmt.Printf("IPv6Gateway : %s \n", scon.NetworkSettings.DefaultNetworkSettings.IPv6Gateway)
                        fmt.Printf("GlobalIPv6Address : %s \n", scon.NetworkSettings.DefaultNetworkSettings.GlobalIPv6Address)
                        fmt.Printf("GlobalIPv6PrefixLen %d \n: ", scon.NetworkSettings.DefaultNetworkSettings.GlobalIPv6PrefixLen)
                        fmt.Printf("MacAddress : %s \n", scon.NetworkSettings.DefaultNetworkSettings.MacAddress)

                   }

		fmt.Println("*************************************\n\n")   
	 }

}   

