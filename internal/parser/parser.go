package parser

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

func ParseTestReport(r io.Reader) []Package {
	pkgs := []Package{}
	tests := map[string]*TestFunc{}

	s := bufio.NewScanner(r)
	s.Split(bufio.ScanLines)
	prevErr := false
	errFuncName := ""
	for s.Scan() {
		line := s.Text()
		if isTestExecDescription(line) {
			prevErr = false
			tcase := extractTestExecDescription(line)
			tests[tcase] = NewTestFunc(tcase)
			continue
		}

		if isTestStatusDescription(line) {
			prevErr = false
			tcase, status, time := extractTestStatusDescription(line)
			if _, ok := tests[tcase]; ok {
				tests[tcase].SetResults(status, time)
				if status == "FAIL" {
					prevErr = true
					errFuncName = tcase
				}
			}
			continue
		}

		if isTestedPkgDescription(line) {
			prevErr = false
			pkgname, status, time := extractTestedPkgDescription(line)
			pkg := Package{
				Name:      pkgname,
				Status:    status,
				Time:      time,
				Functions: groupTestFunc(tests),
			}
			pkg.FindCoverage()
			tests = map[string]*TestFunc{}
			pkgs = append(pkgs, pkg)
			continue
		}
		if isUntestedPackageDescription(line) {
			prevErr = false
			if pkgName := extractUntestedPkgDescription(line); pkgName != "" {
				pkg := Package{Name: pkgName}
				pkgs = append(pkgs, pkg)
			}
			continue
		}
		fmt.Println(line)

		if _, ok := tests[errFuncName]; ok && prevErr {
			tests[errFuncName].SetError(s.Text())
		}
	}
	return pkgs

}

func groupTestFunc(tests map[string]*TestFunc) []*TestFunc {
	for name, t := range tests {
		if s := strings.Split(name, "/"); len(s) > 1 {
			if _, ok := tests[s[0]]; ok {
				t.Name = strings.Replace(s[1], "_", " ", -1)
				tests[s[0]].AddSubTests(t)
				delete(tests, name)
			}
		}
	}
	arrTest := []*TestFunc{}
	for _, t := range tests {
		arrTest = append(arrTest, t)
	}
	return arrTest
}
