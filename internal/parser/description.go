package parser

import (
	"regexp"

	"github.com/fatih/color"
	"github.com/gokultp/go-tprof/internal/reports"
)

var rgxDescrition = regexp.MustCompile(`(?m)^===[ \t]+RUN[ \t]+(?P<testcase>[\w\W]+)$`)

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
	return rgxDescrition.MatchString(d.text)
}

// Println will print the line with formatting and colors
func (d *DescriptionParser) Println() {
	color.New(color.FgBlue).Println(d.text)
}

// UpdateReports will update the reports and temp map by reference
func (d *DescriptionParser) UpdateReports(
	r *reports.Report,
	f map[string]*reports.TestFunc,
	failed *string,
) {
	*failed = ""
	values := rgxDescrition.FindStringSubmatch(d.text)
	for i, key := range rgxDescrition.SubexpNames() {
		if key == "testcase" {
			f[values[i]] = reports.NewTestFunc(values[i])
		}
	}
}
