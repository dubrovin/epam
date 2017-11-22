package main

import (
	"epam/models"
	"epam/servers"
	"epam/services"
	"flag"
	"github.com/qiangxue/fasthttp-routing"
	"log"
	"time"
)

var (
	addr    = flag.String("addr", ":8080", "http service address")
	ttl     = flag.String("ttl", "100s", "time cancel reserve")
	service = flag.String("service", "db", "service for running")
)

func main() {
	flag.Parse()
	defaultTTL, err := time.ParseDuration(*ttl)
	if err != nil {
		log.Fatal(err)
	}
	if *service == "db" {
		db := services.NewDataBase()
		db.AddProduct(&models.Product{TTL: defaultTTL})

		controller := servers.NewDBController(db, routing.New(), defaultTTL)
		controller.Run(*addr, defaultTTL)
		for {
			println("Heartbeat")
			time.Sleep(time.Second * 10)
			db.AddProduct(&models.Product{TTL: defaultTTL})
		}
	}

	if *service == "watcher" {
		_ = 0
	}

	if *service == "reserver" {
		_ = 0
	}

}
