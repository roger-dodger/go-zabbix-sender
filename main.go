package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/adubkov/go-zabbix"
)

var (
	zabbixHost = flag.String("server", "localhost", "Hostname or IP of the Zabbix server")
	zabbixPort = flag.Int("port", 10051, "Port of the Zabbix server trapper")
	host       = flag.String("host", "localhost", "Specify the host name the item belongs to")
	key        = flag.String("key", "", "Item key")
	value      = flag.String("value", "", "Item value")
)

func main() {
	flag.Parse()

	if *key == "" {
		fmt.Println("Key is missing - terminating")
		os.Exit(1)
	}

	if *value == "" {
		fmt.Println("Value is missing - terminating")
		os.Exit(1)
	}

	// Prepare Metric
	var metrics []*zabbix.Metric
	metrics = append(metrics, zabbix.NewMetric(*host, *key, *value))

	// Create instance of Packet class
	packet := zabbix.NewPacket(metrics)

	// Send packet to Zabbix
	z := zabbix.NewSender(*zabbixHost, *zabbixPort)
	z.Send(packet)

	fmt.Printf("Sent metric %q with value %q for host %q to Zabbix server \"%v:%v\"\n", *key, *value, *host, *zabbixHost, *zabbixPort)
}
