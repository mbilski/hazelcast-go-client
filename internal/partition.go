package internal

import (
	"github.com/hazelcast/go-client/internal/common"
	. "github.com/hazelcast/go-client/internal/protocol"
	"github.com/hazelcast/go-client/internal/serialization"
	"log"
	"sync"
	"time"
)

const PARTITION_UPDATE_INTERVAL = 5

type PartitionService struct {
	client         *HazelcastClient
	partitions     map[int32]*Address
	partitionCount int32
	mu             sync.Mutex
	cancel         chan bool
	alive          chan bool
}

func NewPartitionService(client *HazelcastClient) *PartitionService {
	return &PartitionService{client: client, partitions: make(map[int32]*Address), cancel: make(chan bool, 0), alive: make(chan bool, 0)}
}
func (partitionService *PartitionService) start() {
	go func() {
		for {
			select {
			case <-partitionService.alive:
				//TODO::
				time.Sleep(time.Duration(PARTITION_UPDATE_INTERVAL) * time.Second)
				go partitionService.doRefresh()
			case <-partitionService.cancel:
				return
			}
		}
	}()
	partitionService.doRefresh()

}
func (partitionService *PartitionService) PartitionCount() int32 {
	partitionService.mu.Lock()
	defer partitionService.mu.Unlock()
	return int32(len(partitionService.partitions))
}

func (partitionService *PartitionService) PartitionOwner(partitionId int32) (*Address, bool) {
	address, ok := partitionService.partitions[partitionId]
	return address, ok
}

func (partitionService *PartitionService) GetPartitionId(keyData *serialization.Data) int32 {
	count := partitionService.PartitionCount()
	if count <= 0 {
		return 0
	}
	return common.HashToIndex(keyData.GetPartitionHash(), count)
}
func (partitionService *PartitionService) doRefresh() {
	partitionService.mu.Lock()
	defer partitionService.mu.Unlock()
	address := partitionService.client.ClusterService.ownerConnectionAddress
	connectionChan := partitionService.client.ConnectionManager.GetConnection(address)
	connection, alive := <-connectionChan
	if !alive {
		log.Print("connection is closed")
		//TODO::Handle connection closed
		return
	}
	if connection == nil {
		//TODO:: Handle error
	}
	request := ClientGetPartitionsEncodeRequest()
	result, err := partitionService.client.InvocationService.InvokeOnConnection(request, connection).Result()
	if err != nil {
		//TODO:: Handle error
	}
	partitionService.processPartitionResponse(result)
	partitionService.alive <- true

}
func (partitionService *PartitionService) processPartitionResponse(result *ClientMessage) {
	partitions := ClientGetPartitionsDecodeResponse(result).Partitions
	partitionService.partitions = make(map[int32]*Address)
	for _, partitionList := range *partitions {
		addr := partitionList.Key().(*Address)
		for _, partition := range partitionList.Value().([]int32) {
			partitionService.partitions[int32(partition)] = addr
		}
	}
}
func (partitionService *PartitionService) shutdown() {
	go func() {
		partitionService.cancel <- true
	}()
}