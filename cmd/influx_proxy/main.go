package main

import (
	"flag"

	"../../pkg/core"
)

func main() {
	var (
		etcdEndPoints string
		address       string
	)

	flag.StringVar(&etcdEndPoints, "etcd", "http://127.0.0.1:4001", "etcd endpoints, using , to split")
	flag.StringVar(&address, "address", ":10000", "http api address")

	ps := core.NewProxyServer(address)
	ps.Start()

	core.SignalHandling()
}
