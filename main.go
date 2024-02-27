package main

import (
	"log"

	"github.com/patrik-koska-clone/youtube-go/pkg/config"
	"github.com/patrik-koska-clone/youtube-go/pkg/desktop"
	"github.com/patrik-koska-clone/youtube-go/pkg/youtubeadapter"
)

var configFilePath = "config.yaml"

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
