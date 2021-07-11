package main

import (
	"io/ioutil"
	"log"

	lxd "github.com/lxc/lxd/client"
	"github.com/lxc/lxd/shared/api"
)

type containerNetwork struct {
	hostname string
	hwaddr   string
	ip       string
}

func authLxd(connectionType string, connectionParams ...string) (lxd.InstanceServer, error) {

	var connection lxd.InstanceServer
	var err error

	if connectionType == "tcp" {
		if len(connectionParams) < 3 {
			log.Fatal("Error: LXD URL, Cert and KEY files are required for TCP connection")
		}

		lxdURL := connectionParams[0]
		lxdCrt := connectionParams[1]
		lxdKey := connectionParams[2]

		clientCrt, err := ioutil.ReadFile(lxdCrt)
		if err != nil {
			log.Panic(err)
		}

		clientKey, err := ioutil.ReadFile(lxdKey)
		if err != nil {
			log.Panic(err)
		}

		args := lxd.ConnectionArgs{}
		args.TLSClientCert = string(clientCrt)
		args.TLSClientKey = string(clientKey)
		args.InsecureSkipVerify = true
		connection, err = lxd.ConnectLXD(lxdURL, &args)

		if err != nil {
			log.Fatal(err)
		}
	}

	if connectionType == "socket" {
		if len(connectionParams) < 1 {
			log.Fatal("Error: LXD socket path is missing")
		}

		lxdSocket := connectionParams[0]

		// Connect to LXD over the Unix socket
		connection, err = lxd.ConnectLXDUnix(lxdSocket, nil)

		if err != nil {
			log.Fatal(err)
		}
	}

	return connection, nil

}

func getContainersInventory() ([]api.ContainerFull, error) {
	var err error
	containers, err := Connection.GetContainersFull()

	if err != nil {
		log.Fatal(err)
	}

	return containers, nil
}

func getNetworkInventory() map[string]containerNetwork {
	networks, err := Connection.GetNetworkLeases("lxdbr0")
	if err != nil {
		log.Fatal(err)
	}

	inventory := make(map[string]containerNetwork)

	for _, network := range networks {
		containerName := network.Hostname
		inventory[containerName] = containerNetwork{
			hostname: network.Hostname,
			hwaddr:   network.Hwaddr,
			ip:       network.Address,
		}

	}
	return inventory
}

func getContainerIp(name string) string {
	containerIp := getNetworkInventory()[name].ip
	return containerIp
}
