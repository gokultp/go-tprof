package parser

import (
	"fmt"
	"regexp"
	"sync"

	"github.com/gokultp/go-tprof/internal/reports"
)

var rgxSkippedPackages = regexp.MustCompile(`(?m)\?[ \t]+(?P<package>([a-z.-]*\/)*[a-z.-]+)[ \t]+\[no test files\]$`)

// SpippedPackageParser parses test execution SpippedPackages
type SpippedPackageParser struct {
	text string
}

// NewSpippedPackageParser returns a new instance of SpippedPackageParser
func NewSpippedPackageParser(line string) *SpippedPackageParser {
	return &SpippedPackageParser{
		text: line,
	}
}

// IsAbleToParse will say this parser is able to parse the given text
func (d *SpippedPackageParser) IsAbleToParse() bool {
	return rgxSkippedPackages.MatchString(d.text)
}

// Println will print the line with formatting and colors
func (d *SpippedPackageParser) Println() {
	fmt.Println(d.text)
}

// UpdateReports will update the reports and temp map by reference
func (d *SpippedPackageParser) UpdateReports(s *Scanner, wg *sync.WaitGroup) {
	s.ResetLastErrorFunc()
	var pkg string
	values := rgxSkippedPackages.FindStringSubmatch(d.text)
	for i, key := range rgxSkippedPackages.SubexpNames() {
		if key == "package" {
			pkg = values[i]
			break
		}
	}
	s.report.AddPackage(reports.NewPackage(pkg))
	s.ResetIterator()
}
