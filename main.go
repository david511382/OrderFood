package main

import (
	"flag"
	"fmt"
	"net"
	_ "orderfood/docs"
	"orderfood/src/config"
	"orderfood/src/handler"
	"orderfood/src/logic"
	"strings"
)

// @title Order Food API
// @version 1.0
// @description 訂餐系統

// @BasePath /api/

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	flagParse()

	cfg := config.Get()

	logic.Init(cfg)

	router := handler.Init()

	if ips, err := getIP(); err == nil {
		releaseServer := &config.Config{
			Server: config.Server{
				Port: cfg.Port,
			},
		}

		for _, ip := range ips {
			if strings.HasPrefix(ip, "192.168.0.") {
				releaseServer.Server.Host = ip
				break
			}
		}

		if releaseServer.Server.Host == "" {
			releaseServer.Server.Host = ips[len(ips)-1]
		}

		go router.Run(releaseServer.Domain())
	} else {
		fmt.Println(err)
	}

	router.Run(cfg.Domain())
}

func getIP() ([]string, error) {
	ifaces, err := net.Interfaces()
	if err != nil {
		return nil, err
	}

	result := make([]string, 0)
	for _, i := range ifaces {
		addrs, err := i.Addrs()
		if err != nil {
			return nil, err
		}

		for _, addr := range addrs {
			if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
				if ipnet.IP.To4() != nil {
					result = append(result, ipnet.IP.String())
				}
			}
		}
	}

	return result, nil
}

func flagParse() {
	configFileName := "./config/config-develop.yml"
	flag.StringVar(&configFileName, "config-file", "./src/config/config.yml", "config path")

	flag.Parse()

	err := config.ReadConfig(configFileName)
	if err != nil {
		panic(err)
	}
}
