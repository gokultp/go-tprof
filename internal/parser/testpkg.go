package parser

type Package struct {
	Name      string
	Functions []*TestFunc
	Status    string
	Time      string
}
