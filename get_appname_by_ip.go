package main

import (
	"git.qutoutiao.net/infr/ab/abserver/app/components/hosts"
)

func main () {
	hosts.InitHosts("http://cmdb.qutoutiao.net/api/v0.1/servers")
	println("hosts")
	println(hosts.GetHost(string("172.16.109.142")))
}
