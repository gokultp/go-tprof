package parser

import (
	"github.com/fatih/color"
)

// ErrorParser parser will parse error messages
type ErrorParser struct {
	text string
}

// NewErrorParser returns a new instance of ErrorParser
func NewErrorParser(line string) *ErrorParser {
	return &ErrorParser{
		text: line,
	}
}

// IsAbleToParse will say this parser is able to parse the given text
func (d *ErrorParser) IsAbleToParse() bool {
	return true
}

// Println will print the line with formatting and colors
func (d *ErrorParser) Println() {
	printWithColor(color.FgRed, d.text)
	underline := make([]rune, len(d.text))
	for i := 0; i < len(d.text); i++ {
		if d.text[i] == ' ' || d.text[i] == '\t' {
			underline[i] = rune(d.text[i])
			continue
		}
		underline[i] = '^'
	}
	printWithColor(color.FgRed, string(underline))
}

// UpdateReports will update the reports and temp map by reference
func (d *ErrorParser) UpdateReports(s *Scanner) {
	if s.lastErrorFunc != nil {
		if _, ok := s.testMapIterator[*s.lastErrorFunc]; ok {
			s.testMapIterator[*s.lastErrorFunc].SetError(d.text)
		}
	}

}
