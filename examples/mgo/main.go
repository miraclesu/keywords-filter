package main

import (
	"fmt"

	"github.com/miraclesu/keywords-filter"
	"github.com/miraclesu/keywords-filter/loader/mgo.load"
)

func main() {
	// db.keywords.insert({word:'xxoo',kind:'porn',rate:100})
	// db.symbols.insert({word:'*'})
	loader, err := load.New("mgo.json")
	if err != nil {
		fmt.Printf("new loader err: %s\n", err.Error())
		return
	}

	f, err := filter.New(100, loader)
	if err != nil {
		fmt.Printf("new filter err: %s\n", err.Error())
		return
	}

	content := "test *xx**oo something"
	resp := f.Filter(content)
	fmt.Printf("%s is %+v\n", content, resp)
}
