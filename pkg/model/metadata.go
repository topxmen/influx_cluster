package model

import (
	"fmt"
	"time"
)

/*
ShardGroup contains multiple shards to maximize the number of data nodes utilized.
If there are N data nodes in the cluster and the replication factor is X, then N/X shards will be created in each shard group, discarding any fractions.
This means that a new shard group will get created for each day of data that gets written in.
*/
type ShardGroup struct {
	ID     uint64
	Start  time.Time
	End    time.Time
	Shards []Shard
}

/*
Shard is sharded from ShardGroup according to hash value
*/
type Shard struct {
	Replicas []Replica
}

/*
Replica is an individual replica from a Shard, each replica is mapping to a storage data in a Node.
*/
type Replica struct {
	NodeID uint64
}

/*
Node is an individual host machine running influxdb instance
*/
type Node struct {
	ID      uint64 `json:"id"`
	Address string `json:"address"`
	Host    string `json:"host"`
}

func (n Node) HTTPAddress() string {
	return fmt.Sprintf("http://%s", n.Address)
}
