package sql

import "strconv"

type SearchOptions struct {
	Page      uint8
	ListSize  uint8
	Sorter    string
	Component string
	Q         string
}

func NewSearchOptions(page, listSize, sorter, component, q string) (options *SearchOptions, err error) {
	var p = 1
	if page != "" {
		p, err = strconv.Atoi(page)
	}
	if err != nil {
		return
	}
	var ls = 5
	if listSize != "" {
		ls, err = strconv.Atoi(listSize)
	}
	if err != nil {
		return
	}

	options = &SearchOptions{
		Page:      uint8(p),
		ListSize:  uint8(ls),
		Sorter:    sorter,
		Component: component,
		Q:         q,
	}
	return
}
