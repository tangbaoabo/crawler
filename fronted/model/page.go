package model

import "crawler/engine"

type SearchResult struct {
	//总共多少条
	Hints      int64
	Start      int
	Query      string
	PrevFrom   int
	NextFrom   int
	Items      []engine.Item
}
