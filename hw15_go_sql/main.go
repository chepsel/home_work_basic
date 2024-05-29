package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	config "github.com/chepsel/home_work_basic/hw15_go_sql/config"
	logger "github.com/chepsel/home_work_basic/hw15_go_sql/internal/logger"
	restapi "github.com/chepsel/home_work_basic/hw15_go_sql/internal/restapi"
	source "github.com/chepsel/home_work_basic/hw15_go_sql/internal/source"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func main() {
	database := &source.Database{}
	var err error
	cfg := config.ReadConfig()
	dsn := cfg.ConnectionString()

	database.Logger = logger.New(cfg.LogLevel())
	database.DB, err = sqlx.Connect("postgres", dsn)
	if err != nil {
		database.LogError("connect", err)
		os.Exit(2)
	}
	defer database.DB.Close()

	database.DB.SetMaxOpenConns(cfg.Database.ConnPull)
	database.DB.SetMaxIdleConns(cfg.Database.ConnPull)

	server := &http.Server{
		Addr:         fmt.Sprintf("%s:%s", cfg.Server.Host, cfg.Server.Port),
		ReadTimeout:  time.Duration(cfg.Server.RWTimout * int(time.Second)),
		WriteTimeout: time.Duration(cfg.Server.RWTimout * int(time.Second)),
		IdleTimeout:  time.Duration(cfg.Server.IdleTimeout * int(time.Second)),
	}

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	go func() {
		<-sigs
		ctx, cancel := context.WithTimeout(context.Background(), 10)
		defer cancel()
		if err = server.Shutdown(ctx); err != nil {
			log.Println("Can't stop server")
			return
		}
	}()
	r := restapi.NewRouter(database.Logger)
	url := "/store/v1/restapi"
	http.HandleFunc(fmt.Sprintf("%s/%s", url, "user"), restapi.UsersHandler(r, database))                // User
	http.HandleFunc(fmt.Sprintf("%s/%s", url, "users"), restapi.UsersListHandler(r, database))           // Users
	http.HandleFunc(fmt.Sprintf("%s/%s", url, "product"), restapi.ProductsHandler(r, database))          // Products
	http.HandleFunc(fmt.Sprintf("%s/%s", url, "products"), restapi.ProductsListHandler(r, database))     // Product
	http.HandleFunc(fmt.Sprintf("%s/%s", url, "order"), restapi.OrersHandler(r, database))               // Order
	http.HandleFunc(fmt.Sprintf("%s/%s", url, "order/user"), restapi.UserOrdersHandler(r, database))     // UserOrders
	http.HandleFunc(fmt.Sprintf("%s/%s", url, "statistic/user"), restapi.UserStatHandler(r, database))   // StatisticUser
	http.HandleFunc(fmt.Sprintf("%s/%s", url, "statistic/users"), restapi.UsersStatHandler(r, database)) // StatisticUsers
	if err = server.ListenAndServe(); err != nil {
		log.Println(err)
	}
}
