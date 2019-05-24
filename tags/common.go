package tags

import (
	"flag"
	"fmt"
	"orderfood/firewall"
	"orderfood/src/config"
	"orderfood/src/util"

	"github.com/gin-gonic/gin"
)

const (
	firewallName = "OrderFood"
	readTimeOut  = 10000
	writeTimeout = 10000
)

var (
	isReleaseMode    bool
	configPath       string
	cfg              *config.Config
	isManualListenIP bool
)

func InitConfig() (*config.Config, bool) {
	initConfig()
	err := config.ReadConfig(configPath)
	if err != nil {
		panic(err)
	}
	cfg = config.Get()

	flagParse()

	return cfg, isReleaseMode
}

func Run(router *gin.Engine, addr string) {
	run(router, addr)
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
