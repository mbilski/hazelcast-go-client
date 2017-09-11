// Copyright (c) 2008-2017, Hazelcast, Inc. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package protocol

type MapAssignAndGetUuidsResponseParameters struct {
	PartitionUuidList *[]Pair
}

func MapAssignAndGetUuidsCalculateSize() int {
	// Calculates the request payload size
	dataSize := 0
	return dataSize
}

func MapAssignAndGetUuidsEncodeRequest() *ClientMessage {
	// Encode request into clientMessage
	clientMessage := NewClientMessage(nil, MapAssignAndGetUuidsCalculateSize())
	clientMessage.SetMessageType(MAP_ASSIGNANDGETUUIDS)
	clientMessage.IsRetryable = true
	clientMessage.UpdateFrameLength()
	return clientMessage
}

func MapAssignAndGetUuidsDecodeResponse(clientMessage *ClientMessage) *MapAssignAndGetUuidsResponseParameters {
	// Decode response from client message
	parameters := new(MapAssignAndGetUuidsResponseParameters)

	partitionUuidListSize := clientMessage.ReadInt32()
	partitionUuidList := make([]Pair, partitionUuidListSize)
	for partitionUuidListIndex := 0; partitionUuidListIndex < int(partitionUuidListSize); partitionUuidListIndex++ {
		var partitionUuidListItem Pair
		partitionUuidListItemKey := clientMessage.ReadInt32()
		partitionUuidListItemVal := UuidCodecDecode(clientMessage)
		partitionUuidListItem.key = partitionUuidListItemKey
		partitionUuidListItem.value = partitionUuidListItemVal
		partitionUuidList[partitionUuidListIndex] = partitionUuidListItem
	}
	parameters.PartitionUuidList = &partitionUuidList

	return parameters
}