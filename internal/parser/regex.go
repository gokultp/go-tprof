package parser

import (
	"regexp"
)

var (
	rgxUntestedPackages = regexp.MustCompile(`(?m)^\?[ ]+(?P<package>([a-z.]*\/)*[a-z.]+)[ ]+\[no test files\]$`)
	rgxTestedPackages   = regexp.MustCompile(`(?m)^(?P<status>(FAIL|PASS))[ ]+(?P<package>([a-z.]*\/)*[a-z.]+)[ ]+(?P<time>[0-9.]+s)$`)
	rgxExecLine         = regexp.MustCompile(`(?m)^===[ ]+RUN[ ]+(?P<testcase>[a-zA-Z0-9_\/]+)$`)
	rgxTestStatus       = regexp.MustCompile(`(?m)^[ ]*---[ ]+(?P<status>(PASS|FAIL)): (?P<testcase>[a-zA-Z0-9_\/]+)[ ]+\((?P<time>[0-9.]+s)\)$`)
)

func isTestedPkgDescription(line string) bool {
	return rgxTestedPackages.MatchString(line)
}

func isUntestedPackageDescription(line string) bool {
	return rgxUntestedPackages.MatchString(line)
}

func isTestExecDescription(line string) bool {
	return rgxExecLine.MatchString(line)
}

func isTestStatusDescription(line string) bool {
	return rgxTestStatus.MatchString(line)
}

func extractTestExecDescription(line string) string {
	values := rgxExecLine.FindStringSubmatch(line)
	for i, key := range rgxExecLine.SubexpNames() {
		if key == "testcase" {
			return values[i]
		}
	}
	return ""
}

func extractTestStatusDescription(line string) (string, string, string) {
	var name, status, time string
	values := rgxTestStatus.FindStringSubmatch(line)
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
	return name, status, time
}

func extractTestedPkgDescription(line string) (string, string, string) {
	var pkg, status, time string
	values := rgxTestedPackages.FindStringSubmatch(line)
	for i, key := range rgxTestedPackages.SubexpNames() {
		switch key {
		case "package":
			pkg = values[i]
		case "status":
			status = values[i]
		case "time":
			time = values[i]
		}
	}
	return pkg, status, time
}
