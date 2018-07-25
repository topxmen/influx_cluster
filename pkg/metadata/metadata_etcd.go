package metadata

import (
	"time"

	"../model"
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
}

func NewMetaDataManager(etcdEndpoints []string, dialTimeout time.Duration) *MetaDataManager {
	cli, err := client.New(client.Config{
		Endpoints: etcdEndpoints,
	})
	if err != nil {
		glog.Fatal(err)
	}
	mm := &MetaDataManager{
		etcdCli: cli,
	}

	return mm
}

func (mm *MetaDataManager) GetOrCreateShardGroup(start, end time.Time) (shardGroup *model.ShardGroup, err error) {
	return nil, nil
}

func (mm *MetaDataManager) AddNode(addr string, hostName string) (id int64, err error) {
	return 0, nil
}

func (mm *MetaDataManager) UpdateNode(id int64, addr string, hostName string) (err error) {
	return nil
}

func (mm *MetaDataManager) ListNodes() []model.Node {
	return nil
}
