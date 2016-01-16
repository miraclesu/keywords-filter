package filter

import (
    "sort"
)

type symbols []rune

// 二分查找数组内的元素，中文的得用rune类型而不是byte
func (s *symbols) search(r rune) (b bool, index int) {
    if len(*s) == 0 {
        return false, 0
    }
    index = sort.Search(len(*s)-1, func(i int) bool { return (*s)[i] >= r })
    b = (*s)[index] == r
    return
}

// 往有序的数组内插入一个元素
func (s *symbols) add(r rune) {
    b, index := s.search(r)
    if b {
        return
    }
    if len(*s) == 0 {
        *s = append(*s, r)
        return
    }
    if (*s)[index] > r {
        *s = append((*s)[:index], append([]rune{r}, (*s)[index:]...)...)
        return
    }
    *s = append((*s)[:index+1], append([]rune{r}, (*s)[index+1:]...)...)
}

// 删除一个元素
func (s *symbols) remove(r rune) {
    if len(*s) == 0 {
        return
    }
    if b, index := s.search(r); b {
        *s = append((*s)[:index], (*s)[index+1:]...)
    }
}
