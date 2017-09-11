package hazelcast

import (
	"github.com/hazelcast/go-client/config"
	"github.com/hazelcast/go-client/core"
	"github.com/hazelcast/go-client/internal"
)

func NewHazelcastClient() IHazelcastInstance {
	return NewHazelcastClientWithConfig(config.NewClientConfig())
}

func NewHazelcastClientWithConfig(config *config.ClientConfig) IHazelcastInstance {
	return internal.NewHazelcastClient(config)
}

type IHazelcastInstance interface {
	GetMap(name *string) core.IMap
	Shutdown()
}