package filter

import (
	"sort"
	"sync"
)

type symbols struct {
	lk   *sync.RWMutex
	data []rune
}

// 二分查找数组内的元素，中文的得用rune类型而不是byte
func (s *symbols) search(r rune) (b bool, index int) {
	s.lk.RLock()
	data := s.data
	s.lk.RUnlock()

	if len(data) == 0 {
		return false, 0
	}
	index = sort.Search(len(data)-1, func(i int) bool { return data[i] >= r })
	b = data[index] == r
	return
}

// 往有序的数组内插入一个元素
func (s *symbols) add(r rune) {
	b, index := s.search(r)
	if b {
		return
	}

	data := s.data
	if len(data) == 0 {
		data = append(data, r)
	} else if (data)[index] > r {
		data = append(data[:index], append([]rune{r}, data[index:]...)...)
	}
	data = append(data[:index+1], append([]rune{r}, data[index+1:]...)...)

	s.lk.Lock()
	s.data = data
	s.lk.Unlock()
}

// 删除一个元素
func (s *symbols) remove(r rune) {
	data := s.data
	if len(data) == 0 {
		return
	}
	if b, index := s.search(r); b {
		data = append(data[:index], data[index+1:]...)
	}

	s.lk.Lock()
	s.data = data
	s.lk.Unlock()
}
