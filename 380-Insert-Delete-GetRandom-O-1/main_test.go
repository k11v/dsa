package main

import (
	"math/rand/v2"
	"hash/maphash"
	"encoding/binary"
)

// LeetCode

type RandomizedSet struct {
	buckets [][]int
	maxBucketSize int
	length int
	h *maphash.Hash
}

func Constructor() RandomizedSet {
	return RandomizedSet{buckets: make([][]int, 1), maxBucketSize: 8, length: 0, h: &maphash.Hash{}}
}

func (s *RandomizedSet) Insert(v int) bool {
	b := s.hash(v) % uint64(len(s.buckets))

	for i := 0; i < len(s.buckets[b]); i++ {
		if s.buckets[b][i] == v {
			return false
		}
	}

	if len(s.buckets[b]) == s.maxBucketSize {
		s.grow()
		b = s.hash(v) % uint64(len(s.buckets))
	}

	s.buckets[b] = append(s.buckets[b], v)
	s.length++
	return true
}

func (s *RandomizedSet) Remove(v int) bool {
	b := s.hash(v) % uint64(len(s.buckets))

	for i := 0; i < len(s.buckets[b]); i++ {
		if s.buckets[b][i] == v {
			s.buckets[b] = append(s.buckets[b][:i], s.buckets[b][i+1:]...)
			s.length--
			return true
		}
	}

	return false
}

func (s *RandomizedSet) GetRandom() int {
	if s.length == 0 {
		panic("randomized set is empty")
	}

	for {
		b := rand.IntN(len(s.buckets))
		if len(s.buckets[b]) != 0 {
			i := rand.IntN(len(s.buckets[b]))
			return s.buckets[b][i]
		}
	}
}

func (s *RandomizedSet) grow() {
	newBuckets := make([][]int, len(s.buckets)*2)
	for b := 0; b < len(s.buckets); b++ {
		for i := 0; i < len(s.buckets[b]); i++ {
			newB := s.hash(s.buckets[b][i]) % uint64(len(newBuckets))
			newBuckets[newB] = append(newBuckets[newB], s.buckets[b][i])
		}
	}
	s.buckets = newBuckets
}

func (s *RandomizedSet) hash(v int) uint64 {
	s.h.Reset()
	binary.Write(s.h, binary.LittleEndian, v)
	return s.h.Sum64()
}
