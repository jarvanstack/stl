package bloom

import (
	"bytes"
	"encoding/binary"
	"math"
	"stl/algorithm/hash"
	"stl/ds/bitmap"
	"stl/utils/sync"
	gosync "sync"
)

const salt = "g9hmj2fhgr"

var defaultLocker sync.FakeLocker

// Options holds BloomFilter's options
type Options struct {
	 //TODO: Complete me!
}

// Option is a function type used to set Options
type Option func(opt *Options)

// WithGoroutineSafe is used to config a BloomFilter with goroutine-safe
func WithGoroutineSafe() Option {
	 //TODO: Complete me!
}

// BloomFilter is an implementation of bloom filter
type BloomFilter struct {
	 //TODO: Complete me!
}

// New creates a new BloomFilter with m bits and k hash functions
func New(m, k uint64, opts ...Option) *BloomFilter {
	 //TODO: Complete me!
}

// NewWithEstimates creates a new BloomFilter with n and fp.
// n is the capacity of the BloomFilter
// fp is the tolerated error rate of the BloomFilter
func NewWithEstimates(n uint64, fp float64, opts ...Option) *BloomFilter {
	 //TODO: Complete me!
}

// NewFromData creates a new BloomFilter from data generated by function 'Data()'
func NewFromData(data []byte, opts ...Option) *BloomFilter {
	 //TODO: Complete me!
}

// EstimateParameters estimates m and k from n and p
func EstimateParameters(n uint64, p float64) (m uint64, k uint64) {
	 //TODO: Complete me!
}

// Add adds val to the BloomFilter
func (bf *BloomFilter) Add(val string) {
	 //TODO: Complete me!
}

// Contains returns true if val is (high probability) in the BloomFilter, otherwise returns false.
func (bf *BloomFilter) Contains(val string) bool {
	 //TODO: Complete me!
}

// Data returns the data of the BloomFilter, it can bee used to creates a new BloomFilter by using function 'NewFromData' .
func (bf *BloomFilter) Data() []byte {
	 //TODO: Complete me!
}
