package parser

import (
	"regexp"

	"github.com/fatih/color"
	"github.com/gokultp/go-tprof/internal/reports"
)

var rgxDescription = regexp.MustCompile(`(?m)^===[ \t]+RUN[ \t]+(?P<testcase>[\w\W]+)$`)

// DescriptionParser parses test execution descriptions
type DescriptionParser struct {
	text string
}

// NewDescriptionParser returns a new instance of DescriptionParser
func NewDescriptionParser(line string) *DescriptionParser {
	return &DescriptionParser{
		text: line,
	}
}

// IsAbleToParse will say this parser is able to parse the given text
func (d *DescriptionParser) IsAbleToParse() bool {
	return rgxDescription.MatchString(d.text)
}

// Println will print the line with formatting and colors
func (d *DescriptionParser) Println() {
	printWithColor(color.FgBlue, d.text)
}

// UpdateReports will update the reports and temp map by reference
func (d *DescriptionParser) UpdateReports(s *Scanner) {
	s.ResetLastErrorFunc()
	values := rgxDescription.FindStringSubmatch(d.text)
	for i, key := range rgxDescription.SubexpNames() {
		if key == "testcase" {
			s.testMapIterator[values[i]] = reports.NewTestFunc(values[i])
		}
	}
}
