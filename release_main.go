//+build release

package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"net"
	"net/http"
	"orderfood/firewall"
	"orderfood/src/config"
	"orderfood/src/util"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

const (
	firewallName = "OrderFood"
	readTimeOut  = 10000
	writeTimeout = 10000
)

var (
	isReleaseMode bool

	isManualListenIP bool
)

func initServer() {
	flagParse()

	isReleaseMode = true

	addFireWall()
}

func flagParse() {
	//flag.BoolVar(&isManualListenIP, "ip", false, "is manual set ip")

	flag.Parse()
}

func addFireWall() {
	appname := "orderfood.exe"
	appname, err := util.GetFilePath(appname)
	if err != nil {
		panic(err)
	}
	dirs := "in"
	action := "allow"

	err = firewall.AddFireWall(firewallName, appname, dirs, action)
	if err != nil {
		panic(err)
	}

	fmt.Println("open fire wall")
}

func removeFireWall() {
	err := firewall.DelFireWall(firewallName)
	if err != nil {
		panic(err)
	}

	fmt.Println("close fire wall")
}

func run(router *gin.Engine, addr string) {
	defer removeFireWall()

	showAddr()

	s := &http.Server{
		Handler:        router,
		Addr:           addr,
		ReadTimeout:    readTimeOut,
		WriteTimeout:   writeTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	go func() {
		if err := s.ListenAndServe(); err != nil {
			fmt.Printf("Listen: %s\n", err)
		}
	}()

	input := ""
	for {
		fmt.Println("input q to quit server")
		fmt.Scanln(&input)
		if input == "q" {
			break
		}
	}

	fmt.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := s.Shutdown(ctx); err != nil {
		fmt.Println("Server Shutdown:", err)
	}

	fmt.Println("Server exiting")
}

func showAddr() {
	releaseIP, err := getIP()
	if err != nil {
		panic(err)
	}

	releaseServer := &config.Config{
		Server: config.Server{
			Host: releaseIP,
			Port: cfg.Port,
		},
	}

	releaseIP = releaseServer.Domain()

	fmt.Println("current addr : " + releaseIP)
}

func getIP() (string, error) {
	const ipStart = "192.168."
	targetMask := net.IPv4Mask(255, 255, 255, 0)
	targetMaskStr := targetMask.String()

	ifaces, err := net.Interfaces()
	if err != nil {
		return "", err
	}

	for _, i := range ifaces {
		addrs, err := i.Addrs()
		if err != nil {
			return "", err
		}

		for _, addr := range addrs {
			if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
				if ipnet.IP.To4() != nil {
					if maskStr := ipnet.Mask.String(); maskStr == targetMaskStr {
						if ip := ipnet.IP.String(); strings.HasPrefix(ip, ipStart) {
							return ip, nil
						}
					}
				}
			}
		}
	}

	return "", errors.New("ip not found")
}
