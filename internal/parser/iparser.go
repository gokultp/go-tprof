package parser

import (
	"github.com/gokultp/go-tprof/internal/reports"
)

// Parser is the interface of public methods implemeted by parsers
type Parser interface {
	IsAbleToParse() bool
	UpdateReports(*reports.Report, map[string]*reports.TestFunc, *string)
	Println()
}
