package server

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"sync"
	"syscall"
	"time"

	"github.com/chepsel/home_work_basic/hw13_http/server/source"
)

const (
	DefaultHost string = "default"
	DefaultPort string = "default"
)

type Animals struct {
	db *source.Storage
	mu *sync.Mutex
}

func Server(host string, port string) {
	if host == DefaultHost || port == DefaultPort {
		log.Fatalf("- wrong start params \n- port:\"%s\";\n- host:\"%s\";\n",
			port,
			host)
	}

	database := &Animals{db: source.FileDB(), mu: &sync.Mutex{}}
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT)
	go func() {
		<-sigs
		database.db.SaveBeforeClose()
		log.Println("Stopped after save")
		os.Exit(0)
	}()
	var url strings.Builder
	var address strings.Builder
	url.WriteString("/v1")      // lvl 1
	url.WriteString("/restapi") // lvl 2
	url.WriteString("/animal")  // lvl 3 method
	address.WriteString(host)
	address.WriteString(":")
	address.WriteString(port)
	log.Printf(" - URL is - http://%s%s", address.String(), url.String())
	server := &http.Server{
		Addr:              address.String(),
		ReadHeaderTimeout: 3 * time.Second,
	}

	http.HandleFunc(url.String(), database.Route) // POST
	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}

func (database Animals) Route(resp http.ResponseWriter, req *http.Request) {
	from := time.Now()
	var logString strings.Builder
	resultCode := 200
	logString.WriteString(fmt.Sprintf("- [%s] - | method: %s | url: %s ", req.RemoteAddr, req.Method, req.URL))
	defer func() {
		till := time.Now()
		diff := till.UnixMicro() - from.UnixMicro()
		log.Printf("%s | code: %d | processing: %d |", logString.String(), resultCode, diff)
	}()
	switch req.Method {
	case http.MethodGet:
		params := req.URL.Query()
		id := params.Get("id")
		logString.WriteString(fmt.Sprintf("| id:%s", id))
		res, err := database.db.Get(id, database.mu)
		if err != nil {
			resultCode = http.StatusBadRequest
			http.Error(resp, err.Error(), resultCode)
			return
		}
		json, err := res.MarshalJSON()
		if err != nil {
			resultCode = http.StatusInternalServerError
			http.Error(resp, err.Error(), resultCode)
			return
		}
		resp.Write(json)
		return
	case http.MethodDelete:
		params := req.URL.Query()
		id := params.Get("id")
		logString.WriteString(fmt.Sprintf("| id:%s", id))
		if len(id) == 0 {
			resultCode = http.StatusBadRequest
			http.Error(resp, "Wrong request params", resultCode)
			return
		}
		err := database.db.Delete(id, database.mu)
		switch {
		case errors.Is(err, source.MissingID):
			resultCode = http.StatusBadRequest
			http.Error(resp, source.MissingID.Error(), resultCode)
			return
		case err != nil:
			resultCode = http.StatusServiceUnavailable
			http.Error(resp, err.Error(), resultCode)
			return
		default:
			resp.Write([]byte("done"))
			return
		}
	case http.MethodPost:
		var animal source.Animal
		bodyBytes, err := io.ReadAll(req.Body)
		defer req.Body.Close()
		logString.WriteString(fmt.Sprintf("| body: %s", string(bodyBytes)))
		if err != nil {
			resultCode = http.StatusBadRequest
			http.Error(resp, err.Error(), resultCode)
			return
		}
		err = animal.UnmarshalJSON(bodyBytes)
		if err != nil {
			resultCode = http.StatusBadRequest
			http.Error(resp, err.Error(), resultCode)
			return
		}
		if len(animal.ID) > 0 && len(animal.Name) > 0 {
			err = database.db.Put(animal.ID, animal, database.mu)
			if err != nil {
				resultCode = http.StatusServiceUnavailable
				http.Error(resp, err.Error(), resultCode)
				return
			}
			resp.Write([]byte("success"))
			return
		} else {
			resultCode = http.StatusBadRequest
			http.Error(resp, source.MissingKey.Error(), resultCode)
			return
		}
	default:
		resultCode = http.StatusMethodNotAllowed
		http.Error(resp, "Method not used", resultCode)
		return
	}
}
