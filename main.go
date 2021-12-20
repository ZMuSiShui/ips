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
	flag.BoolVar(&conf.Update, "update", false, "update software")
	flag.StringVar(&conf.ConfigFile, "-config", "data/config.yml", "config file")
	flag.StringVar(&conf.ConfigFile, "cfg", "data/config.yml", "config file")
	flag.Parse()
}

func Init() bool {
	client.InitLog()
	client.InitConfig()
	client.InitCron()
	client.InitCache()
	return true
}

func main() {
	if conf.Help {
		fmt.Printf("%v\n", conf.Usage)
		return
	}
	if conf.Version {
		fmt.Printf("Built At: %s\nGo Version: %s\nVersion: %s\n", conf.BuiltAt, conf.GoVersion, conf.VERSION)
		return
	}
	if !Init() {
		os.Exit(1)
	}
	if conf.Debug {
		log.Info("Set Debug Mode")
	}
	log.Info("Starting Client")
}
