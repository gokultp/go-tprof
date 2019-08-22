package parser

import (
	"bufio"
	"io"

	"github.com/gokultp/go-tprof/internal/reports"
)
// ParseTestReport will parse reports
func ParseTestReport(r io.Reader) *reports.Report {
	reps := &reports.Report{}
	tests := map[string]*reports.TestFunc{}
	s := bufio.NewScanner(r)
	s.Split(bufio.ScanLines)
	var failedFuc string

	for s.Scan() {
		line := s.Text()
		p := GetParser(line, failedFuc)
		p.UpdateReports(reps, tests, &failedFuc)
		p.Println()
	}
	return reps
}
