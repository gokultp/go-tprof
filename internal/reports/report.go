package reports

// Report is the final reports contract
type Report struct {
	Packages []*Package `json:"pkgs"`
}
