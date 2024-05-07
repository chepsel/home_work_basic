package server

import (
	"errors"
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
	DefaultHost string = "default"
	DefaultPort string = "default"
)

func Server(host string, port string) {
	if host == DefaultHost || port == DefaultPort {
		log.Fatalf("- wrong start params \n- port:\"%s\";\n- host:\"%s\";\n",
			port,
			host)
	}
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
			err := database.Delete(id, &mu)
			switch {
			case errors.Is(err, source.MissingID):
				c.JSONP(http.StatusGone, gin.H{"err": err})
			case err != nil:
				c.JSONP(http.StatusServiceUnavailable, gin.H{"err": err})
			default:
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
				putErr := database.Put(animal.ID, animal, &mu)
				if putErr != nil {
					c.JSONP(http.StatusServiceUnavailable, gin.H{"err": err})
				}
				c.JSONP(200, "success")
			} else {
				c.JSONP(http.StatusBadRequest, gin.H{"err": source.MissingKey})
			}
		})
	}
	var address strings.Builder
	address.WriteString(host)
	address.WriteString(":")
	address.WriteString(port)
	server.Run(address.String())
}
