package filter

import (
	. "github.com/miraclesu/keywords-filter/keyword"
)

type Response struct {
	Threshold int
	Rate      int
	Keywords  []Keyword
}
