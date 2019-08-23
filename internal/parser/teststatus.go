package parser

import (
	"regexp"
)

const (
	statusFail = "FAIL"
	statusPass = "PASS"
	statusOK   = "ok"
)

var rgxTestStatus = regexp.MustCompile(`(?m)^[ \t]*---[ \t]+(?P<status>(PASS|FAIL|SKIP)): (?P<testcase>[\w\W]+)[ \t]+\((?P<time>[0-9.]+s)\)$`)

// TestStatusParser parses test execution status
type TestStatusParser struct {
	text   string
	status string
}

// NewTestStatusParser returns a new instance of TestStatusParser
func NewTestStatusParser(line string) *TestStatusParser {
	return &TestStatusParser{
		text: line,
	}
}

// IsAbleToParse will say this parser is able to parse the given text
func (d *TestStatusParser) IsAbleToParse() bool {
	return rgxTestStatus.MatchString(d.text)
}

// Println will print the line with formatting and colors
func (d *TestStatusParser) Println() {
	c := getColourByStatus(d.status)
	printWithColor(c, d.text)
}

// UpdateReports will update the reports and temp map by reference
func (d *TestStatusParser) UpdateReports(s *Scanner) {
	s.ResetLastErrorFunc()
	var name, status, time string
	values := rgxTestStatus.FindStringSubmatch(d.text)
	for i, key := range rgxTestStatus.SubexpNames() {
		switch key {
		case "testcase":
			name = values[i]
		case "status":
			status = values[i]
		case "time":
			time = values[i]
		}
	}
	d.status = status
	if status == statusFail {
		*s.lastErrorFunc = name
	}
	if _, ok := s.testMapIterator[name]; ok {
		s.testMapIterator[name].SetResults(status, time)
	}
}
