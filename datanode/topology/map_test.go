package topology

import (
	"github.com/m3db/m3/src/cluster/shard"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	aresShard "github.com/uber/aresdb/datanode/shard"
)

func newTestShardSet(shards []uint32) aresShard.ShardSet {
	values := aresShard.NewShards(shards, shard.Available)
	shardSet := aresShard.NewShardSet(values)

	return shardSet
}

var _ = Describe("map and watch", func() {
	It("map", func() {
		hosts := []struct {
			id     string
			addr   string
			shards []uint32
		}{
			{"h1", "h1:9000", []uint32{0}},
			{"h2", "h2:9000", []uint32{1}},
			{"h3", "h3:9000", []uint32{0}},
			{"h4", "h4:9000", []uint32{1}},
		}
		var hostShardSets []HostShardSet
		for _, h := range hosts {
			hostShardSets = append(hostShardSets,
				NewHostShardSet(NewHost(h.id, h.addr), newTestShardSet(h.shards)))
		}
		shardSet := newTestShardSet([]uint32{0, 1})
		opts := NewStaticOptions().
			SetShardSet(shardSet).
			SetHostShardSets(hostShardSets).
			SetReplicas(3)
		m := NewStaticMap(opts)

		Ω(m.HostsLen()).Should(Equal(4))
		for i, h := range hosts {
			Ω(h.id).Should(Equal(m.Hosts()[i].ID()))
			Ω(h.addr).Should(Equal(m.Hosts()[i].Address()))

			Ω(h.id).Should(Equal(m.HostShardSets()[i].Host().ID()))
			Ω(h.addr).Should(Equal(m.HostShardSets()[i].Host().Address()))
			Ω(h.shards).Should(Equal(m.HostShardSets()[i].ShardSet().AllIDs()))
		}

		hostShardSet, ok := m.LookupHostShardSet("h1")
		Ω(ok).Should(Equal(true))
		Ω(hostShardSet.Host().ID()).Should(Equal(hosts[0].id))
		Ω(hostShardSet.Host().Address()).Should(Equal(hosts[0].addr))
		Ω(hostShardSet.ShardSet().AllIDs()).Should(Equal(hosts[0].shards))

		hostShardSet, ok = m.LookupHostShardSet("h0")
		Ω(ok).Should(Equal(false))
		Ω(hostShardSet).Should(BeNil())

		hostsGot, err := m.RouteShard(0)
		Ω(err).Should(BeNil())
		Ω(hostsGot[0].ID()).Should(Equal("h1"))
		Ω(hostsGot[1].ID()).Should(Equal("h3"))

		replicasGot := m.Replicas()
		Ω(replicasGot).Should(Equal(3))
	})
})
