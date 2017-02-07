package docker

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"

	"../utils"
	containertypes "github.com/docker/engine-api/types/container"   //containertypes为包别名，包路径为github.com/docker/engine-api/types/container，包名为container 
	networktypes "github.com/docker/docker/api/types/network"
	"github.com/docker/go-connections/nat"
)

type Container struct {
	ID          string
	State       State
	Pnum        int
	Created     time.Time
	Path        string
	Args        []string
	Config      *containertypes.Config 
	NetworkSettings *NetworkSettings     
	MountPoints map[string]MountPoint
	Name        string
	LogPath     string
	ResolvConfPath string
	HostnamePath   string
	HostsPath      string
	RestartCount   int
	Driver         string
	MountLabel     string
	ProcessLabel   string
	AppArmorProfile    string

}

// NetworkSettings exposes the network settings in the api
type NetworkSettings struct {
	NetworkSettingsBase
	DefaultNetworkSettings
	Networks map[string]*networktypes.EndpointSettings
}

// NetworkSettingsBase holds basic information about networks
type NetworkSettingsBase struct {
	Bridge                 string      // Bridge is the Bridge name the network uses(e.g. `docker0`)
	SandboxID              string      // SandboxID uniquely represents a container's network stack
	HairpinMode            bool        // HairpinMode specifies if hairpin NAT should be enabled on the virtual interface
	LinkLocalIPv6Address   string      // LinkLocalIPv6Address is an IPv6 unicast address using the link-local prefix
	LinkLocalIPv6PrefixLen int         // LinkLocalIPv6PrefixLen is the prefix length of an IPv6 unicast address
	Ports                  nat.PortMap // Ports is a collection of PortBinding indexed by Port
	SandboxKey             string      // SandboxKey identifies the sandbox
	SecondaryIPAddresses   []networktypes.Address
	SecondaryIPv6Addresses []networktypes.Address
}

// DefaultNetworkSettings holds network information
// during the 2 release deprecation period.
// It will be removed in Docker 1.11.
type DefaultNetworkSettings struct {
	EndpointID          string // EndpointID uniquely represents a service endpoint in a Sandbox
	Gateway             string // Gateway holds the gateway address for the network
	GlobalIPv6Address   string // GlobalIPv6Address holds network's global IPv6 address
	GlobalIPv6PrefixLen int    // GlobalIPv6PrefixLen represents mask length of network's global IPv6 address
	IPAddress           string // IPAddress holds the IPv4 address for the network
	IPPrefixLen         int    // IPPrefixLen represents mask length of network's IPv4 address
	IPv6Gateway         string // IPv6Gateway holds gateway address specific for IPv6
	MacAddress          string // MacAddress holds the MAC address for the network
}



// Containers returns an array of docker containers unmarshaled from config.json
// in docker root path.
func Containers() ([]Container, error) {
	containersPath := filepath.Join(GetDockerRoot(), "containers")

	containerEntries, err := ioutil.ReadDir(containersPath)
	if err != nil {
		return nil, err
	}

	containers := []Container{}

	for _, entry := range containerEntries {
		ID := entry.Name()
		if len(ID) != len("ffb082df6289394f4d285ef2ea31051deed699f6b352cf4109fb7e97fd15237a") {
			continue
		}

		// FIXME: ID is exactly correct, not to call getContainer to travese
		con, err := getContainer(ID)

		if err != nil {
			continue
		}

		containers = append(containers, con) 
	}
	return containers, nil
}

// GetContainer returns a container by given IDOrName
func GetContainer(IDOrName string) (Container, error) {
	return getContainer(IDOrName)
}

// getContainer is unexported
func getContainer(IDOrName string) (Container, error) {
	containersPath := filepath.Join(GetDockerRoot(), "containers")

	containerEntries, err := ioutil.ReadDir(containersPath)
	if err != nil {
		return Container{}, err
	}

	matchedNum := 0
	var matchID string

	// traverse all entries to match IDOrName
	for _, entry := range containerEntries {
		if strings.HasPrefix(entry.Name(), IDOrName) {
			matchedNum++
			matchID = entry.Name()
		}
	}

	// if more than 1 container ID has the prefix, return an error
	if matchedNum >= 2 {
		return Container{}, fmt.Errorf("More than one container have container ID prefix of %s\n"+
			"Please specify another container ID.", IDOrName)
	}

	// exact 1 container matches IDOrName
	if matchedNum == 1 {
		return getContainerFromConfig(containersPath, matchID)
	}

	// TODO: support container name matching
	if matchedNum == 0 {
		return Container{}, fmt.Errorf("No container with such ID %s", IDOrName)
	}

	return Container{}, nil
}

// getConfigJsonPath gets config json file path for a container
func getContainerFromConfig(containersPath, entryName string) (Container, error) {
	configFilename, err := getConfigFilename()
	if err != nil {
		return Container{}, err
	}

	configJsonPath := filepath.Join(containersPath, entryName, configFilename)

	file, err := os.Open(configJsonPath)
	if err != nil {
		return Container{}, err
	}
	defer file.Close()

	data, err := ioutil.ReadAll(file)
	if err != nil {
		return Container{}, err
	}

	var con Container

	if err := json.Unmarshal(data, &con); err != nil {
		return Container{}, err
	}
    
	return con, nil
}

// getConfigFilename gets container's config file name by docker version
// If docker version differs, json file's name differs, too.
// config.json in 1.10.0-, while config.v2.json in 1.10.0+
func getConfigFilename() (string, error) {
	var configFilename string

	match, err := utils.CompareDockerVersion(GetDockerVersion(), "1.10.0")
	if err != nil {
		return "", err
	}
	// If match is true, it means current docker version is newer or at least equal.
	if match {
		configFilename = "config.v2.json"
	} else {
		configFilename = "config.json"
	}
	return configFilename, nil
}
