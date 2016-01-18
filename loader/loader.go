package loader

import (
	. "github.com/miraclesu/keywords-filter/keyword"
)

type Loader interface {
	Load() ([]*Keyword, []string, error)
}
