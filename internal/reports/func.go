package reports

// TestFunc encapsulates a test function's reports
type TestFunc struct {
	Name     string
	Status   string
	Time     string
	Error    string
	SubTests []*TestFunc
}

// NewTestFunc returns a new instance of TestFunc
func NewTestFunc(name string) *TestFunc {
	return &TestFunc{
		Name: name,
	}
}

// SetResults will set parsed results
func (f *TestFunc) SetResults(status, time string) {
	f.Status = status
	f.Time = time
}

// SetError will set the error text
func (f *TestFunc) SetError(err string) {
	f.Error += " " + err
}

// AddSubTests will add subtests for the func
func (f *TestFunc) AddSubTests(t *TestFunc) {
	f.SubTests = append(f.SubTests, t)
}
