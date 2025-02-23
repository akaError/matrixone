// Copyright 2021 Matrix Origin
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package tuplecodec

type KVType int

const (
	KV_MEMORY KVType = iota
	KV_CUBE KVType = iota + 1
)

type KVHandler interface {
	GetKVType() KVType

	// NextID gets the next id of the type
	NextID(typ string) (uint64,error)

	// Set writes the key-value (overwrite)
	Set(key TupleKey,value TupleValue) error

	// SetBatch writes the batch of key-value (overwrite)
	SetBatch(keys []TupleKey,values []TupleValue) []error

	// DedupSet writes the key-value. It will fail if the key exists
	DedupSet(key TupleKey, value TupleValue) error

	// DedupSetBatch writes the batch of keys-values. It will fail if there is one key exists
	DedupSetBatch(keys []TupleKey, values []TupleValue) []error

	// Delete deletes the key
	Delete(key TupleKey) error

	// Get gets the value of the key
	Get(key TupleKey)(TupleValue, error)

	// GetBatch gets the values of the keys
	GetBatch(keys []TupleKey)([]TupleValue, error)

	// GetRange gets the values among the range [startKey,endKey).
	GetRange(startKey TupleKey, endKey TupleKey) ([]TupleValue, error)

	// GetRange gets the values from the startKey with limit
	GetRangeWithLimit(startKey TupleKey, limit uint64) ([]TupleKey,[]TupleValue, error)

	// GetWithPrefix gets the values of the prefix with limit
	// The prefixLen denotes the prefix[:prefixLen] is the real prefix
	GetWithPrefix(prefix TupleKey, prefixLen int, limit uint64) ([]TupleKey, []TupleValue, error)

	// GetShardsWithRange get the shards that holds the range [startKey,endKey)
	GetShardsWithRange(startKey TupleKey, endKey TupleKey) (interface{}, error)

	// GetShardsWithPrefix get the shards that holds the keys with prefix
	GetShardsWithPrefix(prefix TupleKey) (interface{}, error)
}