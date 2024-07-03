package main

import (
	"eCommerce/config"
	"eCommerce/init/app"
	"flag"
)

var envFlag = flag.String("config", "./.env.toml", "env not found")

func main() {
	flag.Parse()
	c := config.NewConfig(*envFlag)
	app.NewApplication(c)
}
