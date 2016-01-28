package filter

import (
	"github.com/miraclesu/keywords-filter/keyword"
)

// Response is used to return the filter result
type Response struct {
	Threshold int
	Rate      int
	Keywords  []keyword.Keyword
}
