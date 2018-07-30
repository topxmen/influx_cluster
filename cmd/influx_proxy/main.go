package main

import (
	"flag"

	"../../pkg/core"
)

func main() {
	var (
		etcdEndPoints string
		address       string
		serverID      uint64
	)

	flag.StringVar(&etcdEndPoints, "etcd", "http://127.0.0.1:4001", "etcd endpoints, using , to split")
	flag.StringVar(&address, "address", ":10000", "http api address")
	flag.Uint64Var(&serverID, "server_id", 1, "server id for id generator, avoid duplication")
	flag.Parse()

	ps := core.NewProxyServer(address)
	ps.Start()

	core.SignalHandling()
}
