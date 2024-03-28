package main

import (
	"log"
	"net/http"
	"time"

	"github.com/rohinibarge/productorder/product"
	"github.com/rohinibarge/productorder/scheduler"
)

func main() {
	task := product.Order{
		Url: "https://dummyjson.com/products",
	}
	scheduler := scheduler.NewScheduler(5*time.Second, task)
	go scheduler.Start()
	port := "8300"
	http.Handle("/", http.FileServer(http.Dir("web")))
	log.Printf("Serving %s on HTTP port: %s\n", "web", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
