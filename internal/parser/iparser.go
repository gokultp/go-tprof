package parser

import "sync"

// Parser is the interface of public methods implemeted by parsers
type Parser interface {
	IsAbleToParse() bool
	UpdateReports(*Scanner, *sync.WaitGroup)
	Println()
}
