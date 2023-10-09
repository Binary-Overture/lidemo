package main

import (
	"cncyx.xyz/router"
	"log"
)

func main() {
	r := router.Router()
	port := ":8080"
	log.Println("端口号为", port)
	err := r.Run(port)
	if err != nil {
		return
	}
}
