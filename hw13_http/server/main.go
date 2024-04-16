package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"sync"
	"syscall"

	"github.com/chepsel/home_work_basic/hw13_http/server/source"
	"github.com/gin-gonic/gin"
)

var database *source.Storage

const (
	defaultPort string = "default"
	defaultHost string = "default"
)

type Config struct {
	port string
	host string
}

var conf *Config = &Config{}

func init() {
	log.Println("- Start")
	flag.StringVar(&conf.port, "port", defaultPort, "listening port")
	flag.StringVar(&conf.port, "p", conf.port, "listening port")
	flag.StringVar(&conf.host, "host", defaultHost, "listening address")
	flag.StringVar(&conf.host, "h", conf.host, "listening address")
	log.Println(len(flag.Args()))
	flag.Parse()
	if conf.host == defaultHost || conf.port == defaultPort {
		log.Fatalf("- wrong start params \n- port:\"%s\";\n- host:\"%s\";\n",
			conf.port,
			conf.host)
	}
	fmt.Printf("- port:\"%s\";\n- host:\"%s\";\n",
		conf.port,
		conf.host)
}

func main() {
	database = source.FileDB()
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT)
	go func() {
		<-sigs
		database.SaveBeforeClose()
		log.Println("Stopped after save")
		os.Exit(0)
	}()
	mu := sync.Mutex{}
	server := gin.Default()
	v1 := server.Group("/v1")
	carsApi := v1.Group("/restapi")
	{
		carsApi.GET("/animal", func(c *gin.Context) {
			id := c.Query("id")
			if len(id) == 0 {
				c.JSONP(http.StatusBadRequest, "Wrong request params")
				return
			}
			if res, err := database.Get(id); err == nil {
				c.JSONP(200, res)
				return
			}
			c.JSONP(http.StatusNotFound, "Not found")

		})
		carsApi.DELETE("/animal", func(c *gin.Context) {
			id := c.Query("id")
			if len(id) == 0 {
				c.JSONP(http.StatusBadRequest, "Wrong request params")
				return
			}
			if err := database.Delete(id, &mu); err != nil {
				c.JSONP(http.StatusGone, gin.H{"err": source.MissingId})
			} else {
				c.JSONP(200, "done")
			}
		})
		carsApi.POST("/animal", func(c *gin.Context) {
			var animal source.Animal
			err := c.ShouldBindJSON(&animal)
			if err != nil {
				c.JSONP(http.StatusBadRequest, gin.H{"err": err.Error()})
			}
			if len(animal.Id) > 0 && len(animal.Name) > 0 {
				database.Put(animal.Id, animal, &mu)
			} else {
				c.JSONP(http.StatusBadRequest, gin.H{"err": source.MissingKey})
			}
		})
	}
	var addres strings.Builder
	addres.WriteString(conf.host)
	addres.WriteString(":")
	addres.WriteString(conf.port)
	server.Run(addres.String())
}
