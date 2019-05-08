package firewall

import (
	"os"
	"os/exec"
)

func AddFireWall(names, appnames, dirs, actions string) error {
	c := exec.Command("netsh", "advfirewall", "firewall", "add", "rule",
		"name="+names,
		"dir="+dirs,
		"action="+actions,
		"program="+appnames,
	)
	c.Stdout = os.Stdout
	c.Stderr = os.Stderr
	return c.Run()
}

func DelFireWall(name string) error {
	cmd := exec.Command("netsh", "advfirewall", "firewall", "delete", "rule", "name="+name)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
