package topology

import (
	"errors"
	"github.com/m3db/m3/src/cluster/placement"
	"github.com/m3db/m3/src/cluster/services"
	"github.com/m3db/m3/src/cluster/shard"
	xwatch "github.com/m3db/m3/src/x/watch"
	"github.com/uber/aresdb/common"
	aresShard "github.com/uber/aresdb/datanode/shard"
	"github.com/uber/aresdb/utils"
	"go.uber.org/zap"
	"sync"
)

var (
	errInvalidService            = errors.New("service topology is invalid")
	errUnexpectedShard           = errors.New("shard is unexpected")
	errMissingShard              = errors.New("shard is missing")
	errNotEnoughReplicasForShard = errors.New("replicas of shard is less than expected")
	errInvalidTopology           = errors.New("could not parse latest value from config service")
)

type dynamicInitializer struct {
	sync.Mutex

	opts DynamicOptions
	topo Topology
}

// NewDynamicInitializer returns a dynamic topology initializer
func NewDynamicInitializer(opts DynamicOptions) Initializer {
	return &dynamicInitializer{opts: opts}
}

func (i *dynamicInitializer) Init() (Topology, error) {
	i.Lock()
	defer i.Unlock()

	if i.topo != nil {
		return i.topo, nil
	}

	topo, err := newDynamicTopology(i.opts)
	if err != nil {
		return nil, err
	}

	i.topo = topo
	return i.topo, nil
}

func (i *dynamicInitializer) TopologyIsSet() (bool, error) {
	return true, nil
}

type dynamicTopology struct {
	sync.RWMutex

	opts      DynamicOptions
	services  services.Services
	watch     services.Watch
	watchable xwatch.Watchable
	closed    bool
	logger    common.Logger
}

func newDynamicTopology(opts DynamicOptions) (DynamicTopology, error) {
	services, err := opts.ConfigServiceClient().Services(nil)
	if err != nil {
		return nil, err
	}

	logger := utils.GetLogger()
	logger.Info("waiting for dynamic topology initialization, " +
		"if this takes a long time, make sure that a topology/placement is configured")
	watch, err := services.Watch(opts.ServiceID(), opts.QueryOptions())
	if err != nil {
		return nil, err
	}
	<-watch.C()
	logger.Info("initial topology / placement value received")

	m, err := getMapFromUpdate(watch.Get())
	if err != nil {
		logger.Error("dynamic topology received invalid initial value", zap.Error(err))
		return nil, err
	}

	watchable := xwatch.NewWatchable()
	watchable.Update(m)

	dt := &dynamicTopology{
		opts:      opts,
		services:  services,
		watch:     watch,
		watchable: watchable,
		logger:    logger,
	}
	go dt.run()
	return dt, nil
}

func (t *dynamicTopology) isClosed() bool {
	t.RLock()
	closed := t.closed
	t.RUnlock()
	return closed
}

func (t *dynamicTopology) run() {
	for !t.isClosed() {
		if _, ok := <-t.watch.C(); !ok {
			t.Close()
			break
		}

		m, err := getMapFromUpdate(t.watch.Get())
		if err != nil {
			t.logger.Warn("dynamic topology received invalid update", zap.Error(err))
			continue
		}
		t.watchable.Update(m)
	}
}

func (t *dynamicTopology) Get() Map {
	return t.watchable.Get().(Map)
}

func (t *dynamicTopology) Watch() (MapWatch, error) {
	_, w, err := t.watchable.Watch()
	if err != nil {
		return nil, err
	}
	return NewMapWatch(w), err
}

func (t *dynamicTopology) Close() {
	t.Lock()
	defer t.Unlock()

	if t.closed {
		return
	}

	t.closed = true

	t.watch.Close()
	t.watchable.Close()
}

func (t *dynamicTopology) MarkShardsAvailable(
	instanceID string,
	shardIDs ...uint32,
) error {
	opts := placement.NewOptions()
	ps, err := t.services.PlacementService(t.opts.ServiceID(), opts)
	if err != nil {
		return err
	}
	_, err = ps.MarkShardsAvailable(instanceID, shardIDs...)
	return err
}

func getMapFromUpdate(data interface{}) (Map, error) {
	service, ok := data.(services.Service)
	if !ok {
		return nil, errInvalidTopology
	}
	to, err := getStaticOptions(service)
	if err != nil {
		return nil, err
	}

	return NewStaticMap(to), nil
}

func getStaticOptions(service services.Service) (StaticOptions, error) {
	if service.Replication() == nil || service.Sharding() == nil || service.Instances() == nil {
		return nil, errInvalidService
	}
	replicas := service.Replication().Replicas()
	instances := service.Instances()
	numShards := service.Sharding().NumShards()

	allShardIDs, err := validateInstances(instances, replicas, numShards)
	if err != nil {
		return nil, err
	}

	allShards := make([]shard.Shard, len(allShardIDs))
	for i, id := range allShardIDs {
		allShards[i] = shard.NewShard(uint32(id)).SetState(shard.Available)
	}

	allShardSet := aresShard.NewShardSet(allShards)

	hostShardSets := make([]HostShardSet, len(instances))
	for i, instance := range instances {
		hs, err := NewHostShardSetFromServiceInstance(instance)
		if err != nil {
			return nil, err
		}
		hostShardSets[i] = hs
	}

	return NewStaticOptions().
		SetReplicas(replicas).
		SetShardSet(allShardSet).
		SetHostShardSets(hostShardSets), nil
}

func validateInstances(instances []services.ServiceInstance, replicas, numShards int) ([]uint32, error) {
	m := make(map[uint32]int)
	for _, i := range instances {
		if i.Shards() == nil {
			return nil, errInstanceHasNoShardsAssignment
		}
		for _, s := range i.Shards().All() {
			m[s.ID()] = m[s.ID()] + 1
		}
	}
	s := make([]uint32, numShards)
	for i := range s {
		expectShard := uint32(i)
		count, exist := m[expectShard]
		if !exist {
			return nil, errMissingShard
		}
		if count < replicas {
			return nil, errNotEnoughReplicasForShard
		}
		delete(m, expectShard)
		s[i] = expectShard
	}

	if len(m) > 0 {
		return nil, errUnexpectedShard
	}
	return s, nil
}
