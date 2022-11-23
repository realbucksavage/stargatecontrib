package main

import (
	"log"
	"net/http"
	"time"

	"github.com/realbucksavage/stargate"
	"github.com/realbucksavage/stargatecontrib/lister/eureka"
)

func main() {

	ls := eureka.NewLister("http://some-eureka-1/eureka", "http://some-eureka-2/eureka")
	sg, err := stargate.NewRouter(ls)
	if err != nil {
		log.Fatalf("cannot create stargate proxy: %v", err)
	}

	go func() {
		for {
			time.Sleep(30 * time.Second)

			if err := sg.Reload(); err != nil {
				// sg.Reload() causes Eureka Lister to rebuild its routes based on applications registered in it.
				log.Printf("cannot query eureka for new routes: %v", err)
			}
		}
	}()

	log.Fatal(http.ListenAndServe(":8080", sg))
}
