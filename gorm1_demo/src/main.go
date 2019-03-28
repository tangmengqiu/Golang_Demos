package main

import (
	"flag"
	"gorm_demo/config"
	"gorm_demo/src/model"
)

var (
	cfg = flag.String("config", "conf/cfg.yaml", "gorm demo config path")
)

func main() {

	//init
	if err := config.Init(*cfg); err != nil {
		panic(err)
	}
	model.InitDB()

	return
}
