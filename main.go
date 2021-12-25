package main

import (
	"flag"
	"fmt"
	"os"

	log "github.com/sirupsen/logrus"

	"github.com/ZMuSiShui/ips/client"
	"github.com/ZMuSiShui/ips/conf"
)

func init() {
	flag.BoolVar(&conf.Help, "-help", false, "help message")
	flag.BoolVar(&conf.Help, "h", false, "help message")
	flag.BoolVar(&conf.Debug, "-debug", false, "start with debug mode")
	flag.BoolVar(&conf.Debug, "d", false, "start with debug mode")
	flag.BoolVar(&conf.Version, "-version", false, "print version info")
	flag.BoolVar(&conf.Version, "v", false, "print version info")
	flag.BoolVar(&conf.OnlyGoogle, "g", false, "only specified Google Cloud IP Ranges")
	flag.BoolVar(&conf.OnlyAmazon, "a", false, "only specified Amazon Cloud IP Ranges")
	flag.BoolVar(&conf.OnlyIPv4, "4", false, "only IPv4")
	flag.BoolVar(&conf.OnlyIPv4, "-ipv4 ", false, "only IPv4")
	flag.BoolVar(&conf.OnlyIPv6, "6", false, "only IPv6")
	flag.BoolVar(&conf.OnlyIPv6, "-ipv6 ", false, "only IPv6")
	flag.StringVar(&conf.Service, "s", "", "specified service")
	flag.StringVar(&conf.Service, "-service", "", "specified service")
	flag.StringVar(&conf.Region, "r", "", "specified region")
	flag.StringVar(&conf.Region, "-region", "", "specified region")
	flag.StringVar(&conf.Scope, "-scope", "", "network border group")
	flag.BoolVar(&conf.WriteToFile, "w", false, "write to the file")
	flag.BoolVar(&conf.WriteToFile, "-write", false, "write to the file")
	flag.Parse()
}

func Init() bool {
	client.InitLog()
	return true
}

func main() {
	if conf.Help {
		fmt.Printf("%v\n", conf.Usage)
		return
	}
	if conf.Version {
		fmt.Printf("IPs Version: %s\n", conf.VERSION)
		return
	}
	if !Init() {
		os.Exit(1)
	}
	if conf.Debug {
		log.Info("Set Debug Mode")
	}
	log.Info("Starting Client")
	client.RunClient()
}
