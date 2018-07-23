package metadata

import (
	"../model"
	"time"
)

/*
etcd objects layout:

/influx_cluster/shard_groups
/influx_cluster/shards
/influx_cluster/nodes

*/
type MetaDataManager struct {
}

func NewMetaDataManager() *MetaDataManager {
	mm := &MetaDataManager{}

	return mm
}

func (mm *MetaDataManager) GetOrCreateShardGroup(start, end time.Time) (shardGroup *model.ShardGroup, err error) {
	return nil, nil
}
