package main

import (
	"log"

	lxd "github.com/lxc/lxd/client"
	flag "github.com/spf13/pflag"
)

var Connection lxd.InstanceServer
var outputFormat *string
var outputFile *string

func init() {
	var err error
	var connectionMethod *string = flag.String(
		"connector",
		"tcp",
		"LXD connection method")
	var connectionCrt *string = flag.String(
		"crt",
		"./certs/client.crt",
		"LXD connection certificate")
	var connectionKey *string = flag.String(
		"key",
		"./certs/client.key",
		"LXD connection key")
	var connectionUrl *string = flag.String("url",
		"",
		"LXD connection URL. Ex. https://192.168.1.230:8443")
	var connectionSocket *string = flag.String(
		"socket",
		"/var/snap/lxd/common/lxd/unix.socket",
		"LXD connection socket path")

	outputFormat = flag.String(
		"format",
		"table",
		"Output format ['table', 'csv']")

	outputFile = flag.String(
		"output",
		"os.Stdout",
		"Output destination ['output.txt', 'output.csv']")

	flag.Parse()

	// TCP Connections
	if *connectionMethod == "tcp" {

		if len(*connectionUrl) < 7 {
			flag.Usage()
			log.Fatal("LXD URL is required.  Ex. https://192.168.1.230:8443")
		}
		// LXD TCP Connection
		Connection, err = authLxd("tcp",
			*connectionUrl,
			*connectionCrt,
			*connectionKey,
		)

		if err != nil {
			log.Fatal(err)
		}

		// Fall back to socket connection
	} else if *connectionMethod == "socket" {
		Connection, err = authLxd("socket",
			*connectionSocket,
		)

		if err != nil {
			log.Fatal(err)
		}

	} else {
		flag.Usage()
		log.Fatal("Connector is not valid")
	}

}

func main() {
	// Get containers inventory
	containers, err := getContainersInventory()
	if err != nil {
		log.Fatal(err)
	}

	reportContainers(containers, *outputFormat, *outputFile)

}
