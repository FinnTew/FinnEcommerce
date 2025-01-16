package alg

import (
	"hash"
	"hash/fnv"
	"math"
	"sync"
)

type Hash64 interface {
	hash.Hash64
	Reset()
}

type BloomFilter struct {
	bitSet           []uint64
	numBits          uint64
	numHashFunctions int
	hashFuncs        []Hash64
	mu               sync.RWMutex
}

func NewBloomFilter(expectedElements uint64, falsePositiveRate float64, hashFuncs ...Hash64) *BloomFilter {
	numBits := uint64(math.Ceil(float64(-1*float64(expectedElements)*math.Log(falsePositiveRate)) / math.Pow(math.Log(2), 2)))
	numHashFunctions := int(math.Ceil(math.Log(2) * float64(numBits) / float64(expectedElements)))

	if len(hashFuncs) == 0 {
		hashFuncs = []Hash64{fnv.New64a()}
	}

	return &BloomFilter{
		bitSet:           make([]uint64, int(math.Ceil(float64(numBits)/64))),
		numBits:          numBits,
		numHashFunctions: numHashFunctions,
		hashFuncs:        hashFuncs,
		mu:               sync.RWMutex{},
	}
}

func (bf *BloomFilter) Add(data []byte) {
	bf.mu.Lock()
	defer bf.mu.Unlock()

	for i := 0; i < bf.numHashFunctions; i++ {
		h := bf.hashFuncs[i%len(bf.hashFuncs)]
		h.Reset()
		_, _ = h.Write(data)
		sum64 := h.Sum64()
		bf.bitSet[sum64/64] |= (1 << (sum64 % 64))
	}
}

func (bf *BloomFilter) Contains(data []byte) bool {
	bf.mu.RLock()
	defer bf.mu.RUnlock()

	for i := 0; i < bf.numHashFunctions; i++ {
		h := bf.hashFuncs[i%len(bf.hashFuncs)]
		h.Reset()
		_, _ = h.Write(data)
		sum64 := h.Sum64()
		if (bf.bitSet[sum64/64] & (1 << (sum64 % 64))) == 0 {
			return false
		}
	}
	return true
}
