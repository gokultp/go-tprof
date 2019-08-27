package parser

import (
	"fmt"
	"regexp"
	"strings"
	"sync"

	"github.com/fatih/color"
	"github.com/gokultp/go-tprof/internal/reports"
)

var rgxPackageStatus = regexp.MustCompile(`(?m)^(?P<status>(FAIL|PASS|ok))[ \t]+(?P<package>([a-z.]*\/)*[a-z.]+)[ \t]+(?P<time>[\(\)a-z0-9.]+)$`)

// PackageStatusParser parses package status
type PackageStatusParser struct {
	text   string
	status string
}

// NewPackageStatusParser returns a new instance of PackageStatusParser
func NewPackageStatusParser(line string) *PackageStatusParser {
	return &PackageStatusParser{
		text: line,
	}
}

// IsAbleToParse will say this parser is able to parse the given text
func (d *PackageStatusParser) IsAbleToParse() bool {
	return rgxPackageStatus.MatchString(d.text)
}

// Println will print the line with formatting and colors
func (d *PackageStatusParser) Println() {
	c := getColourByStatus(d.status)
	printWithColor(c, d.text)
}
func printWithColor(c color.Attribute, msg string) {
	if _, err := color.New(c).Println(msg); err != nil {
		fmt.Println(msg)
	}
}

func getColourByStatus(status string) color.Attribute {
	switch status {
	case statusFail:
		return color.FgRed
	case statusPass:
		return color.FgGreen
	default:
		return color.FgYellow
	}
}

// UpdateReports will update the reports and temp map by reference
func (d *PackageStatusParser) UpdateReports(s *Scanner, wg *sync.WaitGroup) {
	s.ResetLastErrorFunc()
	var pkg, status, time string
	values := rgxPackageStatus.FindStringSubmatch(d.text)
	for i, key := range rgxPackageStatus.SubexpNames() {
		switch key {
		case "package":
			pkg = values[i]
		case "status":
			status = values[i]
		case "time":
			time = values[i]
		}
	}
	if status == statusOK {
		status = statusPass
	}
	d.status = status
	p := reports.NewPackage(pkg)
	p.SetResults(status, time, groupTestFunc(s.testMapIterator))
	wg.Add(1)
	go p.FindCoverage(wg)
	p.GetNumberOfPassedCases()
	s.report.AddPackage(p)
	s.ResetIterator()
}

func groupTestFunc(tests map[string]*reports.TestFunc) []*reports.TestFunc {
	for name, t := range tests {
		if s := strings.Split(name, "/"); len(s) > 1 {
			if _, ok := tests[s[0]]; ok {
				t.Name = strings.Replace(s[1], "_", " ", -1)
				tests[s[0]].AddSubTests(t)
				delete(tests, name)
			}
		}
	}
	arrTest := []*reports.TestFunc{}
	for _, t := range tests {
		arrTest = append(arrTest, t)
	}
	return arrTest
}
