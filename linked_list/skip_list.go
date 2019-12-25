// 跳表

package linked_list

import (
	// "fmt"
	rand "math/rand"
	// "math"
)

type SkipListNode struct {
	Key     int
	Forword []*SkipListNode
}

func NewSkipListNode(key, level int) *SkipListNode {
	n := new(SkipListNode)
	n.Key = key
	n.Forword = make([]*SkipListNode, level+1)
	return n
}

type SkipList struct {
	Header   *SkipListNode
	MaxLevel int
	Level    int
	P        float64
}

func NewSkipList(maxLevel int, p float64) *SkipList {
	s := new(SkipList)
	s.MaxLevel = maxLevel
	s.P = p
	return s
}

func (s *SkipList) Insert(key int) {
	cur := s.Header

	if cur == nil {
		s.Header = NewSkipListNode(key, 0)
	} else {
		rLevel := s.genRandomLevel()
		n := NewSkipListNode(key, rLevel)
		if s.Level < rLevel {
			s.Header.Forword = append(s.Header.Forword, make([]*SkipListNode, rLevel-s.Level)...)
			s.Level = rLevel
		}

		nodes := make([]*SkipListNode, s.Level+1)
		for i := s.Level; i >= 0; i-- {
			for cur.Forword[i] != nil && cur.Forword[i].Key < key {
				cur = cur.Forword[i]
			}
			nodes[i] = cur
		}

		if cur.Key == key {
			return
		}

		for i := 0; i <= rLevel; i++ {
			n.Forword[i], nodes[i].Forword[i] = nodes[i].Forword[i], n
		}
	}
}

func (s *SkipList) Search(key int) *SkipListNode {
	cur := s.Header
	for i := s.Level; i >= 0; i-- {
		for cur.Forword[i] != nil && cur.Forword[i].Key <= key {
			cur = cur.Forword[i]
		}
	}

	if cur.Key == key {
		return cur
	}

	return nil
}

func (s *SkipList) Delete(key int) {
	if s.Header.Key == key {
		header := s.Header.Forword[0]

		l := len(header.Forword)
		if l < s.Level {
			header.Forword = append(header.Forword, make([]*SkipListNode, s.Level-l+1)...)
		}

		for i := s.Level; i >= 0; i-- {
			if s.Header.Forword[i] == header && header.Forword[i] == nil {
				header.Forword = header.Forword[:i]
			} else {
				for ; i >= 0; i-- {
					if header.Forword[i] == nil {
						header.Forword[i] = s.Header.Forword[i]
					}
				}
			}
		}

		s.Header = header
		s.Level = len(header.Forword) - 1
	} else {
		cur := s.Header
		nodes := make([]*SkipListNode, s.Level+1)
		for i := s.Level; i >= 0; i-- {
			for cur.Forword[i] != nil && cur.Forword[i].Key < key {
				cur = cur.Forword[i]
			}
			nodes[i] = cur
		}

		cur = cur.Forword[0]
		if cur.Key > key {
			return
		}

		l := len(cur.Forword)
		for i := l - 1; i >= 0; i-- {
			if i > 0 && cur.Forword[i] == nil {
				s.Header.Forword = s.Header.Forword[:i]
			} else {
				nodes[i].Forword[i] = cur.Forword[i]
			}
		}
		s.Level = len(s.Header.Forword) - 1
	}
}

func (s *SkipList) genRandomLevel() int {
	var (
		i      int     = 0
		random float64 = rand.Float64()
	)

	for i < s.MaxLevel && s.P < random {
		i, random = i+1, rand.Float64()
	}
	return i
}
