//+build release

package main

import (
	"errors"
	"flag"
	"fmt"
	"net"
	"orderfood/src/config"
	"strings"
)

var (
	isReleaseMode bool

	isManualListenIP bool
)

func initServer() {
	flagParse()

	isReleaseMode = true
}

func getAddr() string {
	releaseIP, err := getIP()
	if err != nil {
		fmt.Println(err)
		return ""
	}

	releaseServer := &config.Config{
		Server: config.Server{
			Host: releaseIP,
			Port: cfg.Port,
		},
	}

	releaseIP = releaseServer.Domain()

	fmt.Println("current addr : " + releaseIP)

	return releaseIP
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

func flagParse() {
	//flag.BoolVar(&isManualListenIP, "ip", false, "is manual set ip")

	flag.Parse()
}
