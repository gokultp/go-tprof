package parser

import (
	"fmt"

	"github.com/gokultp/go-tprof/internal/reports"
)

// DefaultParser parser is a fallback parser
type DefaultParser struct {
	text string
}

// NewDefaultParser returns a new instance of DefaultParser
func NewDefaultParser(line string) *DefaultParser {
	return &DefaultParser{
		text: line,
	}
}

// IsAbleToParse will say this parser is able to parse the given text
func (d *DefaultParser) IsAbleToParse() bool {
	return true
}

// Println will print the line with formatting and colors
func (d *DefaultParser) Println() {
	fmt.Println(d.text)
}

// UpdateReports will update the reports and temp map by reference
func (d *DefaultParser) UpdateReports(
	r *reports.Report,
	f map[string]*reports.TestFunc,
	failed *string,
) {

}
