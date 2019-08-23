package parser

import (
	"bufio"
	"io"

	"github.com/gokultp/go-tprof/internal/reports"
)

// Scanner will scan stdin and generates reports
type Scanner struct {
	report          *reports.Report
	testMapIterator map[string]*reports.TestFunc
	lastErrorFunc   *string
}

// NewScanner will return  a new instance of Scanner
func NewScanner() *Scanner {
	empty := ""
	return &Scanner{
		report:          &reports.Report{},
		lastErrorFunc:   &empty,
		testMapIterator: map[string]*reports.TestFunc{},
	}
}

// ResetLastErrorFunc will reset the errorfunc
func (sc *Scanner) ResetLastErrorFunc() {
	empty := ""
	*sc.lastErrorFunc = empty
}

// ResetIterator will reset the iterator
func (sc *Scanner) ResetIterator() {
	sc.testMapIterator = map[string]*reports.TestFunc{}
}

// ParseTestReport will parse reports
func (sc *Scanner) ParseTestReport(r io.Reader) *reports.Report {
	s := bufio.NewScanner(r)
	s.Split(bufio.ScanLines)

	for s.Scan() {
		line := s.Text()
		p := GetParser(line, *sc.lastErrorFunc)
		p.UpdateReports(sc)
		p.Println()
	}
	return sc.report
}
