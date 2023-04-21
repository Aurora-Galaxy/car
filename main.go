package main

import (
	"car/conf"
	"car/pkg/util"
	"car/router"
)

func main() {
	conf.Init()
	r := router.NewRouter()
	go util.CheckPile()
	r.Run(conf.HttpPort)
}
