//+build release

package main

import (
	"flag"
	"fmt"
	"net"
	"orderfood/src/config"
	"strconv"
	"strings"
	"time"
)

var (
	isReleaseMode bool

	isManualListenIP bool
	ipChan           chan string
	defaultIP        string
)

func initServer() {
	flagParse()

	isReleaseMode = true
	ipChan, defaultIP = listenIP()
}

func getAddr() string {
	releaseIP := ""

	for releaseIP == "" {
		select {
		case releaseIP = <-ipChan:
		case <-time.After(10 * time.Second):
			releaseIP = defaultIP
			fmt.Println("use default addr")
		}
	}

	releaseServer := &config.Config{
		Server: config.Server{
			Host: releaseIP,
			Port: cfg.Port,
		},
	}

	releaseIP = releaseServer.Domain()

	fmt.Println("use addr : " + releaseIP)

	return releaseIP
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

func listenIP() (ipChan chan string, defaultIP string) {
	ipChan = make(chan string)

	ips, err := getIP()
	if err != nil {
		fmt.Println(err)
		ipChan <- ""
		return
	}

	const ipStart = "192.168."
	defaultIP = ips[len(ips)-1]
	for _, ip := range ips {
		if strings.HasPrefix(ip, ipStart) {
			defaultIP = ip
			break
		}
	}

	go func() {
		defer close(ipChan)

		if isManualListenIP {
			output := make([]string, 0)
			output = append(output, "all ips")
			for i, ip := range ips {
				output = append(output, "id:"+strconv.Itoa(i)+"  ip:"+ip)
			}
			output = append(output, "input ip id")
			outputStr := strings.Join(output, "\n")
			fmt.Println(outputStr)

			input := ""
			fmt.Scanln(&input)
			iip, err := strconv.Atoi(input)
			if err == nil {
				if iip >= len(ips) || iip < 0 {
					fmt.Print("out of range")
				} else {
					defaultIP = ips[iip]
				}
			} else {
				fmt.Println("wrong input " + input)
			}
		}

		ipChan <- defaultIP
	}()

	return
}

func flagParse() {
	flag.BoolVar(&isManualListenIP, "ip", true, "is manual set ip")

	flag.Parse()
}
