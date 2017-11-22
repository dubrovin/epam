package main

import (
	"flag"
	"github.com/dubrovin/epam/models"
	"github.com/dubrovin/epam/servers"
	"github.com/dubrovin/epam/services"
	"github.com/qiangxue/fasthttp-routing"
	"log"
	"time"
)

var (
	addr    = flag.String("addr", ":8080", "http service address")
	ttl     = flag.String("ttl", "100s", "time cancel reserve")
	service = flag.String("service", "db", "service for running")
	dbAddr    = flag.String("db_addr", ":8080", "db service address")
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

		controller := servers.NewDBServer(db, routing.New(), defaultTTL)
		controller.Run(*addr, defaultTTL)
		for {
			println("Heartbeat db")
			time.Sleep(time.Second * 10)
			db.AddProduct(&models.Product{TTL: defaultTTL})
		}
	}

	if *service == "watcher" {

		controller := servers.NewWatcherServer(*dbAddr, routing.New())
		controller.Run(*addr)
		for {
			println("Heartbeat watcher")
			time.Sleep(time.Second * 10)
		}
	}

	if *service == "reserver" {
		controller := servers.NewReserverServer(*dbAddr, routing.New())
		controller.Run(*addr)
		for {
			println("Heartbeat reserver")
			time.Sleep(time.Second * 10)
		}
	}

}
