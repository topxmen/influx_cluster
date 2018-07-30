package model

import (
	"../util"
	"time"

	"github.com/coreos/etcd/client"
	"github.com/golang/glog"
)

/*
etcd objects layout:

/influx_cluster/shard_groups
/influx_cluster/shards
/influx_cluster/nodes

*/
type MetaDataManager struct {
	etcdCli client.Client
	idGen   *util.IDGenerator
}

func NewMetaDataManager(etcdEndpoints []string, dialTimeout time.Duration, idGen *util.IDGenerator) *MetaDataManager {
	cli, err := client.New(client.Config{
		Endpoints: etcdEndpoints,
	})
	if err != nil {
		glog.Fatal(err)
	}
	mm := &MetaDataManager{
		etcdCli: cli,
		idGen:   idGen,
	}

	return mm
}

func (mm *MetaDataManager) GetOrCreateShardGroup(start, end time.Time) (shardGroup *ShardGroup, err error) {
	return nil, nil
}

func (mm *MetaDataManager) AddNode(node Node) (id int64, err error) {
	return 0, nil
}

func (mm *MetaDataManager) UpdateNode(node Node) (err error) {
	return nil
}

func (mm *MetaDataManager) ListNodes() []Node {
	return nil
}
