package model

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"../util"

	"github.com/coreos/etcd/clientv3"
	"github.com/golang/glog"
)

/*
etcd objects layout:

/influx_cluster/shard_groups
/influx_cluster/shards
/influx_cluster/nodes

*/
const (
	prefix    = "/influx_cluster"
	nodesPath = prefix + "/nodes"
)

type MetaDataManager struct {
	etcdCli clientv3.KV
	idGen   *util.IDGenerator
}

func NewMetaDataManager(etcdEndpoints []string, dialTimeout time.Duration, idGen *util.IDGenerator) *MetaDataManager {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   etcdEndpoints,
		DialTimeout: dialTimeout,
	})
	if err != nil {
		glog.Fatal(err)
	}
	mm := &MetaDataManager{
		etcdCli: clientv3.NewKV(cli),
		idGen:   idGen,
	}

	return mm
}

func (mm *MetaDataManager) GetOrCreateShardGroup(start, end time.Time) (shardGroup *ShardGroup, err error) {
	return nil, nil
}

func (mm *MetaDataManager) AddNode(node Node) (id uint64, err error) {
	node.ID = mm.idGen.New()
	err = mm.UpdateNode(node)
	if err != nil {
		return 0, err
	}

	return node.ID, nil
}

func (mm *MetaDataManager) UpdateNode(node Node) (err error) {
	js, _ := json.Marshal(node)
	_, err = mm.etcdCli.Put(context.TODO(), fmt.Sprintf("%s/%d", nodesPath, node.ID), string(js))
	if err != nil {
		return err
	}

	return nil
}

func (mm *MetaDataManager) ListNodes() ([]Node, error) {
	//TODO:
	return nil, nil
}
