package main

import (
	"main/golang/app/config"
	"main/golang/app/listener"
)

func main() {
	listener.NewGinServer().LintenAndServe(config.MainAddress)
}
