package parser

type TestFunc struct {
	Name     string
	Status   string
	Time     string
	Error    string
	SubTests []*TestFunc
}

func NewTestFunc(name string) *TestFunc {
	return &TestFunc{
		Name: name,
	}
}

func (f *TestFunc) SetResults(status, time string) {
	f.Status = status
	f.Time = time
}

func (f *TestFunc) SetError(err string) {
	f.Error = err
}

func (f *TestFunc) AddSubTests(t *TestFunc) {
	f.SubTests = append(f.SubTests, t)
}
