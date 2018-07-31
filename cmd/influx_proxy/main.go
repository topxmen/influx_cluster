package main

import (
	"flag"
	"strings"
	"time"

	"../../pkg/core"
)

func main() {
	var (
		etcdEndPoints   string
		address         string
		serverID        uint64
		etcdDailTimeout int64
	)

	flag.StringVar(&etcdEndPoints, "etcd", "http://127.0.0.1:2379,http://127.0.0.1:4001", "etcd endpoints, using , to split")
	flag.Int64Var(&etcdDailTimeout, "etcd-dail-timeout", 5, "etcd dail timeout")
	flag.StringVar(&address, "address", ":8086", "http api address")
	flag.Uint64Var(&serverID, "server_id", 1, "server id for id generator, avoid duplication")
	flag.Parse()

	ps := core.NewProxyServer(address, strings.Split(etcdEndPoints, ","), time.Duration(etcdDailTimeout)*time.Second, serverID)
	ps.Start()

	core.SignalHandling()
}
