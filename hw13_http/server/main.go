package main

import (
	"flag"
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

var conf *Config

func init() {
	log.Println("- Start")
	conf = &Config{}
	flag.StringVar(&conf.port, "port", defaultPort, "listening port")
	flag.StringVar(&conf.port, "p", conf.port, "listening port")
	flag.StringVar(&conf.host, "host", defaultHost, "listening address")
	flag.StringVar(&conf.host, "h", conf.host, "listening address")
	flag.Parse()
	if conf.host == defaultHost || conf.port == defaultPort {
		log.Fatalf("- wrong start params \n- port:\"%s\";\n- host:\"%s\";\n",
			conf.port,
			conf.host)
	}
	log.Printf("- port:\"%s\", host:\"%s\"\n",
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
	restAPI := v1.Group("/restapi")
	{
		restAPI.GET("/animal", func(c *gin.Context) {
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
		restAPI.DELETE("/animal", func(c *gin.Context) {
			id := c.Query("id")
			if len(id) == 0 {
				c.JSONP(http.StatusBadRequest, "Wrong request params")
				return
			}
			if err := database.Delete(id, &mu); err != nil {
				c.JSONP(http.StatusGone, gin.H{"err": source.MissingID})
			} else {
				c.JSONP(200, "done")
			}
		})
		restAPI.POST("/animal", func(c *gin.Context) {
			var animal source.Animal
			err := c.ShouldBindJSON(&animal)
			if err != nil {
				c.JSONP(http.StatusBadRequest, gin.H{"err": err.Error()})
			}
			if len(animal.ID) > 0 && len(animal.Name) > 0 {
				database.Put(animal.ID, animal, &mu)
				c.JSONP(200, "success")
			} else {
				c.JSONP(http.StatusBadRequest, gin.H{"err": source.MissingKey})
			}
		})
	}
	var address strings.Builder
	address.WriteString(conf.host)
	address.WriteString(":")
	address.WriteString(conf.port)
	server.Run(address.String())
}
