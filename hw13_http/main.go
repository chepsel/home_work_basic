package main

import (
	"flag"
	"log"

	"github.com/chepsel/home_work_basic/hw13_http/client"
	"github.com/chepsel/home_work_basic/hw13_http/server"
)

type Config struct {
	port   string
	host   string
	url    string
	server bool
}

var conf *Config

func init() {
	log.Println("- Start")
	conf = &Config{}
	flag.StringVar(&conf.port, "port", server.DefaultPort, "listening port")
	flag.StringVar(&conf.port, "p", conf.port, "listening port")
	flag.StringVar(&conf.host, "host", server.DefaultHost, "listening address")
	flag.StringVar(&conf.host, "h", conf.host, "listening address")
	flag.StringVar(&conf.url, "url", client.DefaultURL, "request url")
	flag.StringVar(&conf.url, "u", conf.url, "request url")
	flag.BoolVar(&conf.server, "server", false, "is server - bool")
	flag.BoolVar(&conf.server, "s", false, "is server - bool")
	flag.Parse()
	if conf.server {
		log.Printf(" - running server - port:\"%s\", host:\"%s\"\n",
			conf.port,
			conf.host)
	} else {
		log.Printf(" - running client - url:\"%s\"\n",
			conf.url)
	}
}

func main() {
	if conf.server {
		server.Server(conf.host, conf.port)
	} else {
		client.Client(conf.url)
	}
}
