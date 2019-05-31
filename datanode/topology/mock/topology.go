package mock

import (
	"github.com/golang/mock/gomock"
	"github.com/m3db/m3/src/cluster/client"
	"github.com/m3db/m3/src/cluster/services"
	"github.com/m3db/m3/src/dbnode/sharding"
	"github.com/uber/aresdb/datanode/topology"
	"github.com/uber/aresdb/utils"
	"reflect"
)

// MockHost is a mock of topology.Host interface
type MockHost struct {
	ctrl     *gomock.Controller
	recorder *MockHostMockRecorder
}

// MockHostMockRecorder is the mock recorder for MockHost
type MockHostMockRecorder struct {
	mock *MockHost
}

// NewMockHost creates a new mock instance
func NewMockHost(ctrl *gomock.Controller) *MockHost {
	mock := &MockHost{ctrl: ctrl}
	mock.recorder = &MockHostMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockHost) EXPECT() *MockHostMockRecorder {
	return m.recorder
}

// ID mocks base method
func (m *MockHost) ID() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ID")
	ret0, _ := ret[0].(string)
	return ret0
}

// ID indicates an expected call of ID
func (mr *MockHostMockRecorder) ID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ID", reflect.TypeOf((*MockHost)(nil).ID))
}

// Address mocks base method
func (m *MockHost) Address() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Address")
	ret0, _ := ret[0].(string)
	return ret0
}

// Address indicates an expected call of Address
func (mr *MockHostMockRecorder) Address() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Address", reflect.TypeOf((*MockHost)(nil).Address))
}

// String mocks base method
func (m *MockHost) String() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "String")
	ret0, _ := ret[0].(string)
	return ret0
}

// String indicates an expected call of String
func (mr *MockHostMockRecorder) String() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "String", reflect.TypeOf((*MockHost)(nil).String))
}

// MockHostShardSet is a mock of topology.HostShardSet interface
type MockHostShardSet struct {
	ctrl     *gomock.Controller
	recorder *MockHostShardSetMockRecorder
}

// MockHostShardSetMockRecorder is the mock recorder for MockHostShardSet
type MockHostShardSetMockRecorder struct {
	mock *MockHostShardSet
}

// NewMockHostShardSet creates a new mock instance
func NewMockHostShardSet(ctrl *gomock.Controller) *MockHostShardSet {
	mock := &MockHostShardSet{ctrl: ctrl}
	mock.recorder = &MockHostShardSetMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockHostShardSet) EXPECT() *MockHostShardSetMockRecorder {
	return m.recorder
}

// topology.Host mocks base method
func (m *MockHostShardSet) Host() topology.Host {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "topology.Host")
	ret0, _ := ret[0].(topology.Host)
	return ret0
}

// topology.Host indicates an expected call of topology.Host
func (mr *MockHostShardSetMockRecorder) Host() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "topology.Host", reflect.TypeOf((*MockHostShardSet)(nil).Host))
}

// ShardSet mocks base method
func (m *MockHostShardSet) ShardSet() sharding.ShardSet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ShardSet")
	ret0, _ := ret[0].(sharding.ShardSet)
	return ret0
}

// ShardSet indicates an expected call of ShardSet
func (mr *MockHostShardSetMockRecorder) ShardSet() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ShardSet", reflect.TypeOf((*MockHostShardSet)(nil).ShardSet))
}

// MockInitializer is a mock of Initializer interface
type MockInitializer struct {
	ctrl     *gomock.Controller
	recorder *MockInitializerMockRecorder
}

// MockInitializerMockRecorder is the mock recorder for MockInitializer
type MockInitializerMockRecorder struct {
	mock *MockInitializer
}

// NewMockInitializer creates a new mock instance
func NewMockInitializer(ctrl *gomock.Controller) *MockInitializer {
	mock := &MockInitializer{ctrl: ctrl}
	mock.recorder = &MockInitializerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockInitializer) EXPECT() *MockInitializerMockRecorder {
	return m.recorder
}

// Init mocks base method
func (m *MockInitializer) Init() (topology.Topology, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Init")
	ret0, _ := ret[0].(topology.Topology)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Init indicates an expected call of Init
func (mr *MockInitializerMockRecorder) Init() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Init", reflect.TypeOf((*MockInitializer)(nil).Init))
}

// TopologyIsSet mocks base method
func (m *MockInitializer) TopologyIsSet() (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TopologyIsSet")
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// TopologyIsSet indicates an expected call of TopologyIsSet
func (mr *MockInitializerMockRecorder) TopologyIsSet() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TopologyIsSet", reflect.TypeOf((*MockInitializer)(nil).TopologyIsSet))
}

// MockTopology is a mock of topology.Topology interface
type MockTopology struct {
	ctrl     *gomock.Controller
	recorder *MockTopologyMockRecorder
}

// MockTopologyMockRecorder is the mock recorder for MockTopology
type MockTopologyMockRecorder struct {
	mock *MockTopology
}

// NewMockTopology creates a new mock instance
func NewMockTopology(ctrl *gomock.Controller) *MockTopology {
	mock := &MockTopology{ctrl: ctrl}
	mock.recorder = &MockTopologyMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockTopology) EXPECT() *MockTopologyMockRecorder {
	return m.recorder
}

// Get mocks base method
func (m *MockTopology) Get() topology.Map {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get")
	ret0, _ := ret[0].(topology.Map)
	return ret0
}

// Get indicates an expected call of Get
func (mr *MockTopologyMockRecorder) Get() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockTopology)(nil).Get))
}

// Watch mocks base method
func (m *MockTopology) Watch() (topology.MapWatch, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Watch")
	ret0, _ := ret[0].(topology.MapWatch)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Watch indicates an expected call of Watch
func (mr *MockTopologyMockRecorder) Watch() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Watch", reflect.TypeOf((*MockTopology)(nil).Watch))
}

// Close mocks base method
func (m *MockTopology) Close() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Close")
}

// Close indicates an expected call of Close
func (mr *MockTopologyMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockTopology)(nil).Close))
}

// MockDynamicTopology is a mock of DynamicTopology interface
type MockDynamicTopology struct {
	ctrl     *gomock.Controller
	recorder *MockDynamicTopologyMockRecorder
}

// MockDynamicTopologyMockRecorder is the mock recorder for MockDynamicTopology
type MockDynamicTopologyMockRecorder struct {
	mock *MockDynamicTopology
}

// NewMockDynamicTopology creates a new mock instance
func NewMockDynamicTopology(ctrl *gomock.Controller) *MockDynamicTopology {
	mock := &MockDynamicTopology{ctrl: ctrl}
	mock.recorder = &MockDynamicTopologyMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockDynamicTopology) EXPECT() *MockDynamicTopologyMockRecorder {
	return m.recorder
}

// Get mocks base method
func (m *MockDynamicTopology) Get() topology.Map {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get")
	ret0, _ := ret[0].(topology.Map)
	return ret0
}

// Get indicates an expected call of Get
func (mr *MockDynamicTopologyMockRecorder) Get() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockDynamicTopology)(nil).Get))
}

// Watch mocks base method
func (m *MockDynamicTopology) Watch() (topology.MapWatch, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Watch")
	ret0, _ := ret[0].(topology.MapWatch)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Watch indicates an expected call of Watch
func (mr *MockDynamicTopologyMockRecorder) Watch() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Watch", reflect.TypeOf((*MockDynamicTopology)(nil).Watch))
}

// Close mocks base method
func (m *MockDynamicTopology) Close() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Close")
}

// Close indicates an expected call of Close
func (mr *MockDynamicTopologyMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockDynamicTopology)(nil).Close))
}

// MarkShardsAvailable mocks base method
func (m *MockDynamicTopology) MarkShardsAvailable(instanceID string, shardIDs ...int) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{instanceID}
	for _, a := range shardIDs {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "MarkShardsAvailable", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// MarkShardsAvailable indicates an expected call of MarkShardsAvailable
func (mr *MockDynamicTopologyMockRecorder) MarkShardsAvailable(instanceID interface{}, shardIDs ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{instanceID}, shardIDs...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MarkShardsAvailable", reflect.TypeOf((*MockDynamicTopology)(nil).MarkShardsAvailable), varargs...)
}

// MockMapWatch is a mock of topology.MapWatch interface
type MockMapWatch struct {
	ctrl     *gomock.Controller
	recorder *MockMapWatchMockRecorder
}

// MockMapWatchMockRecorder is the mock recorder for MockMapWatch
type MockMapWatchMockRecorder struct {
	mock *MockMapWatch
}

// NewMockMapWatch creates a new mock instance
func NewMockMapWatch(ctrl *gomock.Controller) *MockMapWatch {
	mock := &MockMapWatch{ctrl: ctrl}
	mock.recorder = &MockMapWatchMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockMapWatch) EXPECT() *MockMapWatchMockRecorder {
	return m.recorder
}

// C mocks base method
func (m *MockMapWatch) C() <-chan struct{} {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "C")
	ret0, _ := ret[0].(<-chan struct{})
	return ret0
}

// C indicates an expected call of C
func (mr *MockMapWatchMockRecorder) C() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "C", reflect.TypeOf((*MockMapWatch)(nil).C))
}

// Get mocks base method
func (m *MockMapWatch) Get() topology.Map {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get")
	ret0, _ := ret[0].(topology.Map)
	return ret0
}

// Get indicates an expected call of Get
func (mr *MockMapWatchMockRecorder) Get() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockMapWatch)(nil).Get))
}

// Close mocks base method
func (m *MockMapWatch) Close() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Close")
}

// Close indicates an expected call of Close
func (mr *MockMapWatchMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockMapWatch)(nil).Close))
}

// MockMap is a mock of topology.Map interface
type MockMap struct {
	ctrl     *gomock.Controller
	recorder *MockMapMockRecorder
}

// MockMapMockRecorder is the mock recorder for MockMap
type MockMapMockRecorder struct {
	mock *MockMap
}

// NewMockMap creates a new mock instance
func NewMockMap(ctrl *gomock.Controller) *MockMap {
	mock := &MockMap{ctrl: ctrl}
	mock.recorder = &MockMapMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockMap) EXPECT() *MockMapMockRecorder {
	return m.recorder
}

// Hosts mocks base method
func (m *MockMap) Hosts() []topology.Host {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Hosts")
	ret0, _ := ret[0].([]topology.Host)
	return ret0
}

// Hosts indicates an expected call of Hosts
func (mr *MockMapMockRecorder) Hosts() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Hosts", reflect.TypeOf((*MockMap)(nil).Hosts))
}

// HostShardSets mocks base method
func (m *MockMap) HostShardSets() []topology.HostShardSet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HostShardSets")
	ret0, _ := ret[0].([]topology.HostShardSet)
	return ret0
}

// HostShardSets indicates an expected call of HostShardSets
func (mr *MockMapMockRecorder) HostShardSets() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HostShardSets", reflect.TypeOf((*MockMap)(nil).HostShardSets))
}

// LookupHostShardSet mocks base method
func (m *MockMap) LookupHostShardSet(hostID string) (topology.HostShardSet, bool) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LookupHostShardSet", hostID)
	ret0, _ := ret[0].(topology.HostShardSet)
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

// LookupHostShardSet indicates an expected call of LookupHostShardSet
func (mr *MockMapMockRecorder) LookupHostShardSet(hostID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LookupHostShardSet", reflect.TypeOf((*MockMap)(nil).LookupHostShardSet), hostID)
}

// HostsLen mocks base method
func (m *MockMap) HostsLen() int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HostsLen")
	ret0, _ := ret[0].(int)
	return ret0
}

// HostsLen indicates an expected call of HostsLen
func (mr *MockMapMockRecorder) HostsLen() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HostsLen", reflect.TypeOf((*MockMap)(nil).HostsLen))
}

// ShardSet mocks base method
func (m *MockMap) ShardSet() sharding.ShardSet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ShardSet")
	ret0, _ := ret[0].(sharding.ShardSet)
	return ret0
}

// ShardSet indicates an expected call of ShardSet
func (mr *MockMapMockRecorder) ShardSet() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ShardSet", reflect.TypeOf((*MockMap)(nil).ShardSet))
}

// RouteShard mocks base method
func (m *MockMap) RouteShard(shard int) ([]topology.Host, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RouteShard", shard)
	ret0, _ := ret[0].([]topology.Host)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RouteShard indicates an expected call of RouteShard
func (mr *MockMapMockRecorder) RouteShard(shard interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RouteShard", reflect.TypeOf((*MockMap)(nil).RouteShard), shard)
}

// Replicas mocks base method
func (m *MockMap) Replicas() int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Replicas")
	ret0, _ := ret[0].(int)
	return ret0
}

// Replicas indicates an expected call of Replicas
func (mr *MockMapMockRecorder) Replicas() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Replicas", reflect.TypeOf((*MockMap)(nil).Replicas))
}

// MajorityReplicas mocks base method
func (m *MockMap) MajorityReplicas() int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MajorityReplicas")
	ret0, _ := ret[0].(int)
	return ret0
}

// MajorityReplicas indicates an expected call of MajorityReplicas
func (mr *MockMapMockRecorder) MajorityReplicas() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MajorityReplicas", reflect.TypeOf((*MockMap)(nil).MajorityReplicas))
}

// MockStaticOptions is a mock of topology.StaticOptions interface
type MockStaticOptions struct {
	ctrl     *gomock.Controller
	recorder *MockStaticOptionsMockRecorder
}

// MockStaticOptionsMockRecorder is the mock recorder for MockStaticOptions
type MockStaticOptionsMockRecorder struct {
	mock *MockStaticOptions
}

// NewMockStaticOptions creates a new mock instance
func NewMockStaticOptions(ctrl *gomock.Controller) *MockStaticOptions {
	mock := &MockStaticOptions{ctrl: ctrl}
	mock.recorder = &MockStaticOptionsMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockStaticOptions) EXPECT() *MockStaticOptionsMockRecorder {
	return m.recorder
}

// Validate mocks base method
func (m *MockStaticOptions) Validate() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Validate")
	ret0, _ := ret[0].(error)
	return ret0
}

// Validate indicates an expected call of Validate
func (mr *MockStaticOptionsMockRecorder) Validate() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Validate", reflect.TypeOf((*MockStaticOptions)(nil).Validate))
}

// SetShardSet mocks base method
func (m *MockStaticOptions) SetShardSet(value sharding.ShardSet) topology.StaticOptions {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetShardSet", value)
	ret0, _ := ret[0].(topology.StaticOptions)
	return ret0
}

// SetShardSet indicates an expected call of SetShardSet
func (mr *MockStaticOptionsMockRecorder) SetShardSet(value interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetShardSet", reflect.TypeOf((*MockStaticOptions)(nil).SetShardSet), value)
}

// ShardSet mocks base method
func (m *MockStaticOptions) ShardSet() sharding.ShardSet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ShardSet")
	ret0, _ := ret[0].(sharding.ShardSet)
	return ret0
}

// ShardSet indicates an expected call of ShardSet
func (mr *MockStaticOptionsMockRecorder) ShardSet() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ShardSet", reflect.TypeOf((*MockStaticOptions)(nil).ShardSet))
}

// SetReplicas mocks base method
func (m *MockStaticOptions) SetReplicas(value int) topology.StaticOptions {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetReplicas", value)
	ret0, _ := ret[0].(topology.StaticOptions)
	return ret0
}

// SetReplicas indicates an expected call of SetReplicas
func (mr *MockStaticOptionsMockRecorder) SetReplicas(value interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetReplicas", reflect.TypeOf((*MockStaticOptions)(nil).SetReplicas), value)
}

// Replicas mocks base method
func (m *MockStaticOptions) Replicas() int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Replicas")
	ret0, _ := ret[0].(int)
	return ret0
}

// Replicas indicates an expected call of Replicas
func (mr *MockStaticOptionsMockRecorder) Replicas() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Replicas", reflect.TypeOf((*MockStaticOptions)(nil).Replicas))
}

// SetHostShardSets mocks base method
func (m *MockStaticOptions) SetHostShardSets(value []topology.HostShardSet) topology.StaticOptions {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetHostShardSets", value)
	ret0, _ := ret[0].(topology.StaticOptions)
	return ret0
}

// SetHostShardSets indicates an expected call of SetHostShardSets
func (mr *MockStaticOptionsMockRecorder) SetHostShardSets(value interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetHostShardSets", reflect.TypeOf((*MockStaticOptions)(nil).SetHostShardSets), value)
}

// HostShardSets mocks base method
func (m *MockStaticOptions) HostShardSets() []topology.HostShardSet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HostShardSets")
	ret0, _ := ret[0].([]topology.HostShardSet)
	return ret0
}

// HostShardSets indicates an expected call of HostShardSets
func (mr *MockStaticOptionsMockRecorder) HostShardSets() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HostShardSets", reflect.TypeOf((*MockStaticOptions)(nil).HostShardSets))
}

// MockDynamicOptions is a mock of topology.DynamicOptions interface
type MockDynamicOptions struct {
	ctrl     *gomock.Controller
	recorder *MockDynamicOptionsMockRecorder
}

// MockDynamicOptionsMockRecorder is the mock recorder for MockDynamicOptions
type MockDynamicOptionsMockRecorder struct {
	mock *MockDynamicOptions
}

// NewMockDynamicOptions creates a new mock instance
func NewMockDynamicOptions(ctrl *gomock.Controller) *MockDynamicOptions {
	mock := &MockDynamicOptions{ctrl: ctrl}
	mock.recorder = &MockDynamicOptionsMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockDynamicOptions) EXPECT() *MockDynamicOptionsMockRecorder {
	return m.recorder
}

// Validate mocks base method
func (m *MockDynamicOptions) Validate() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Validate")
	ret0, _ := ret[0].(error)
	return ret0
}

// Validate indicates an expected call of Validate
func (mr *MockDynamicOptionsMockRecorder) Validate() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Validate", reflect.TypeOf((*MockDynamicOptions)(nil).Validate))
}

// SetConfigServiceClient mocks base method
func (m *MockDynamicOptions) SetConfigServiceClient(c client.Client) topology.DynamicOptions {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetConfigServiceClient", c)
	ret0, _ := ret[0].(topology.DynamicOptions)
	return ret0
}

// SetConfigServiceClient indicates an expected call of SetConfigServiceClient
func (mr *MockDynamicOptionsMockRecorder) SetConfigServiceClient(c interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetConfigServiceClient", reflect.TypeOf((*MockDynamicOptions)(nil).SetConfigServiceClient), c)
}

// ConfigServiceClient mocks base method
func (m *MockDynamicOptions) ConfigServiceClient() client.Client {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ConfigServiceClient")
	ret0, _ := ret[0].(client.Client)
	return ret0
}

// ConfigServiceClient indicates an expected call of ConfigServiceClient
func (mr *MockDynamicOptionsMockRecorder) ConfigServiceClient() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ConfigServiceClient", reflect.TypeOf((*MockDynamicOptions)(nil).ConfigServiceClient))
}

// SetServiceID mocks base method
func (m *MockDynamicOptions) SetServiceID(s services.ServiceID) topology.DynamicOptions {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetServiceID", s)
	ret0, _ := ret[0].(topology.DynamicOptions)
	return ret0
}

// SetServiceID indicates an expected call of SetServiceID
func (mr *MockDynamicOptionsMockRecorder) SetServiceID(s interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetServiceID", reflect.TypeOf((*MockDynamicOptions)(nil).SetServiceID), s)
}

// ServiceID mocks base method
func (m *MockDynamicOptions) ServiceID() services.ServiceID {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ServiceID")
	ret0, _ := ret[0].(services.ServiceID)
	return ret0
}

// ServiceID indicates an expected call of ServiceID
func (mr *MockDynamicOptionsMockRecorder) ServiceID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ServiceID", reflect.TypeOf((*MockDynamicOptions)(nil).ServiceID))
}

// SetServicesOverrideOptions mocks base method
func (m *MockDynamicOptions) SetServicesOverrideOptions(opts services.OverrideOptions) topology.DynamicOptions {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetServicesOverrideOptions", opts)
	ret0, _ := ret[0].(topology.DynamicOptions)
	return ret0
}

// SetServicesOverrideOptions indicates an expected call of SetServicesOverrideOptions
func (mr *MockDynamicOptionsMockRecorder) SetServicesOverrideOptions(opts interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetServicesOverrideOptions", reflect.TypeOf((*MockDynamicOptions)(nil).SetServicesOverrideOptions), opts)
}

// ServicesOverrideOptions mocks base method
func (m *MockDynamicOptions) ServicesOverrideOptions() services.OverrideOptions {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ServicesOverrideOptions")
	ret0, _ := ret[0].(services.OverrideOptions)
	return ret0
}

// ServicesOverrideOptions indicates an expected call of ServicesOverrideOptions
func (mr *MockDynamicOptionsMockRecorder) ServicesOverrideOptions() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ServicesOverrideOptions", reflect.TypeOf((*MockDynamicOptions)(nil).ServicesOverrideOptions))
}

// SetQueryOptions mocks base method
func (m *MockDynamicOptions) SetQueryOptions(value services.QueryOptions) topology.DynamicOptions {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetQueryOptions", value)
	ret0, _ := ret[0].(topology.DynamicOptions)
	return ret0
}

// SetQueryOptions indicates an expected call of SetQueryOptions
func (mr *MockDynamicOptionsMockRecorder) SetQueryOptions(value interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetQueryOptions", reflect.TypeOf((*MockDynamicOptions)(nil).SetQueryOptions), value)
}

// QueryOptions mocks base method
func (m *MockDynamicOptions) QueryOptions() services.QueryOptions {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "QueryOptions")
	ret0, _ := ret[0].(services.QueryOptions)
	return ret0
}

// QueryOptions indicates an expected call of QueryOptions
func (mr *MockDynamicOptionsMockRecorder) QueryOptions() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QueryOptions", reflect.TypeOf((*MockDynamicOptions)(nil).QueryOptions))
}

// SetInstrumentOptions mocks base method
func (m *MockDynamicOptions) SetInstrumentOptions(value utils.Options) topology.DynamicOptions {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetInstrumentOptions", value)
	ret0, _ := ret[0].(topology.DynamicOptions)
	return ret0
}

// SetInstrumentOptions indicates an expected call of SetInstrumentOptions
func (mr *MockDynamicOptionsMockRecorder) SetInstrumentOptions(value interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetInstrumentOptions", reflect.TypeOf((*MockDynamicOptions)(nil).SetInstrumentOptions), value)
}

// InstrumentOptions mocks base method
func (m *MockDynamicOptions) InstrumentOptions() utils.Options {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InstrumentOptions")
	ret0, _ := ret[0].(utils.Options)
	return ret0
}

// InstrumentOptions indicates an expected call of InstrumentOptions
func (mr *MockDynamicOptionsMockRecorder) InstrumentOptions() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InstrumentOptions", reflect.TypeOf((*MockDynamicOptions)(nil).InstrumentOptions))
}

// SetHashGen mocks base method
func (m *MockDynamicOptions) SetHashGen(h sharding.HashGen) topology.DynamicOptions {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetHashGen", h)
	ret0, _ := ret[0].(topology.DynamicOptions)
	return ret0
}

// SetHashGen indicates an expected call of SetHashGen
func (mr *MockDynamicOptionsMockRecorder) SetHashGen(h interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetHashGen", reflect.TypeOf((*MockDynamicOptions)(nil).SetHashGen), h)
}

// HashGen mocks base method
func (m *MockDynamicOptions) HashGen() sharding.HashGen {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HashGen")
	ret0, _ := ret[0].(sharding.HashGen)
	return ret0
}

// HashGen indicates an expected call of HashGen
func (mr *MockDynamicOptionsMockRecorder) HashGen() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HashGen", reflect.TypeOf((*MockDynamicOptions)(nil).HashGen))
}

// MockMapProvider is a mock of MapProvider interface
type MockMapProvider struct {
	ctrl     *gomock.Controller
	recorder *MockMapProviderMockRecorder
}

// MockMapProviderMockRecorder is the mock recorder for MockMapProvider
type MockMapProviderMockRecorder struct {
	mock *MockMapProvider
}

// NewMockMapProvider creates a new mock instance
func NewMockMapProvider(ctrl *gomock.Controller) *MockMapProvider {
	mock := &MockMapProvider{ctrl: ctrl}
	mock.recorder = &MockMapProviderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockMapProvider) EXPECT() *MockMapProviderMockRecorder {
	return m.recorder
}

// TopologyMap mocks base method
func (m *MockMapProvider) TopologyMap() (topology.Map, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TopologyMap")
	ret0, _ := ret[0].(topology.Map)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// TopologyMap indicates an expected call of TopologyMap
func (mr *MockMapProviderMockRecorder) TopologyMap() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TopologyMap", reflect.TypeOf((*MockMapProvider)(nil).TopologyMap))
}
