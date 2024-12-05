package main

import (
	"math/rand/v2"
)

// LeetCode

type RandomizedSet struct {
	m map[int]int
	a []int
}

func Constructor() RandomizedSet {
	return RandomizedSet{m: make(map[int]int), a: make([]int, 0)}
}

func (s *RandomizedSet) Insert(v int) bool {
	if _, present := s.m[v]; present {
		return false
	}
	s.a = append(s.a, v)
	s.m[v] = len(s.a)-1
	return true
}

func (s *RandomizedSet) Remove(v int) bool {
	i, present := s.m[v]
	if !present {
		return false
	}

	end := len(s.a)-1
	endValue := s.a[end]

	s.a[i] = s.a[end]
	s.m[endValue] = i
	s.a = s.a[:len(s.a)-1]
	delete(s.m, v)

	return true
}

func (s *RandomizedSet) GetRandom() int {
	if len(s.m) == 0 {
		panic("randomized set is empty")
	}
	i := rand.IntN(len(s.a))
	return s.a[i]
}
