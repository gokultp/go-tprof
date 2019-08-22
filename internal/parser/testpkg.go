package parser

import (
	"bufio"
	"fmt"
	"io"
	"os/exec"
	"strings"
	"sync"
	"syscall"
)

type Package struct {
	Name      string
	Functions []*TestFunc
	Status    string
	Time      string
}

func (p *Package) FindCoverage() int {
	var wg sync.WaitGroup
	r, w := io.Pipe()
	wg.Add(1)
	defer wg.Wait()
	cmd := exec.Command("go", "test", p.Name, "-coverprofile=/tmp/1")
	cmd.Stderr = w
	cmd.Stdout = w
	go consume(&wg, r)

	if err := cmd.Run(); err != nil {
		if ws, ok := cmd.ProcessState.Sys().(syscall.WaitStatus); ok {
			return ws.ExitStatus()
		}
		return 1
	}
	return 0
}

func consume(wg *sync.WaitGroup, r io.Reader) {
	defer wg.Done()
	s := bufio.NewScanner(r)

	for s.Scan() {
		fmt.Println(s.Text())
		if strings.Contains(s.Text(), "coverage: ") {
			return
		}
	}
}
