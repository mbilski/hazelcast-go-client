// Copyright (c) 2008-2020, Hazelcast, Inc. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License")
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

package predicate

import (
	"github.com/mbilski/hazelcast-go-client/serialization"
)

const FactoryID = -32

type Factory struct {
}

func NewFactory() *Factory {
	return &Factory{}
}

func (pf *Factory) Create(id int32) serialization.IdentifiedDataSerializable {
	switch id {
	case sqlID:
		return &SQL{}
	case andID:
		return &And{}
	case betweenID:
		return &Between{}
	case equalID:
		return &Equal{}
	case greaterlessID:
		return &GreaterLess{}
	case likeID:
		return &Like{}
	case ilikeID:
		return &ILike{}
	case inID:
		return &In{}
	case instanceOfID:
		return &InstanceOf{}
	case notEqualID:
		return &NotEqual{}
	case notID:
		return &Not{}
	case orID:
		return &Or{}
	case regexID:
		return &Regex{}
	case falseID:
		return &False{}
	case trueID:
		return &True{}
	default:
		return nil
	}
}
