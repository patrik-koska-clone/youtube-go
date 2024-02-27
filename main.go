package main

import (
	"flag"
	"log"

	"github.com/patrik-koska-clone/youtube-go/config"
	"github.com/patrik-koska-clone/youtube-go/desktop"
	"github.com/patrik-koska-clone/youtube-go/youtubeadapter"
)

var configFilePath = "config.yaml"

func init() {
	flag.Parse()
}

func main() {
	c, err := config.ReadConfig(configFilePath)
	if err != nil {
		log.Fatalf("could not read config\n%v", err)
	}

	y, err := youtubeadapter.New(*c)
	if err != nil {
		log.Fatalf("could not initialize youtube adapter\n%v", err)
	}

	desktop.OpenConsole(*y, *c)
}
