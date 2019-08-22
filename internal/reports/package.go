package reports

import (
	"bufio"
	"io"
	"os/exec"
	"regexp"
	"strconv"
	"sync"
	"syscall"
)

var rgxCoverage = regexp.MustCompile(`(?m)coverage[ \t]*:[ \t]*(?P<coverage>[0-9.]+)\%`)

// Package encapsulates test reports for a package
type Package struct {
	Name      string
	Functions []*TestFunc
	Status    string
	Time      string
	Coverage  float64
}

// NewPackage give a new instace of Package
func NewPackage(name, status, time string, f []*TestFunc) *Package {
	return &Package{
		Name:      name,
		Status:    status,
		Time:      time,
		Functions: f,
	}
}

// FindCoverage will find test coverage for the package
func (p *Package) FindCoverage() int {
	var wg sync.WaitGroup
	r, w := io.Pipe()
	wg.Add(1)
	defer wg.Wait()
	cmd := exec.Command("go", "test", p.Name, "-coverprofile=/tmp/1")
	cmd.Stderr = w
	cmd.Stdout = w
	go p.parseAndSetCoverage(&wg, r)

	if err := cmd.Run(); err != nil {
		if ws, ok := cmd.ProcessState.Sys().(syscall.WaitStatus); ok {
			return ws.ExitStatus()
		}
		return 1
	}
	return 0
}

func (p *Package) parseAndSetCoverage(wg *sync.WaitGroup, r io.Reader) {
	defer wg.Done()
	s := bufio.NewScanner(r)

	for s.Scan() {
		line := s.Text()
		if rgxCoverage.MatchString(line) {
			p.Coverage = getCoverageFromStr(line)
			return
		}
	}
}

func getCoverageFromStr(line string) float64 {
	values := rgxCoverage.FindStringSubmatch(line)
	for i, key := range rgxCoverage.SubexpNames() {
		if key == "coverage" {
			if val, err := strconv.ParseFloat(values[i], 64); err == nil {
				return val
			}
		}
	}
	return 0
}
