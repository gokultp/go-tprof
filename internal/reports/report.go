package reports

// Report is the final reports contract
type Report struct {
	Packages []*Package `json:"pkgs"`
}

// AddPackage will append the package
func (r *Report) AddPackage(pkg *Package) {
	r.Packages = append(r.Packages, pkg)
}
