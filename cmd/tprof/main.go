package main

import (
	"encoding/json"
	"os"
	"sync"

	"github.com/fatih/color"
	"github.com/gokultp/go-tprof/internal/browser"
	"github.com/gokultp/go-tprof/internal/parser"
	"github.com/gokultp/go-tprof/internal/staticserver"
)

func main() {
	var wg sync.WaitGroup
	stat, err := os.Stdin.Stat()
	if err != nil {
		panic(err)
	}
	if (stat.Mode() & os.ModeCharDevice) == 0 {
		scanner := parser.NewScanner()
		res := scanner.ParseTestReport(os.Stdin)
		data, err := json.Marshal(res)
		if err != nil {
			panic(err)
		}
		wg.Add(1)
		go staticserver.StartServer(string(data), "6969", &wg)
		if err := browser.Open("http://localhost:6969"); err != nil {
			panic(err)
		}
		wg.Wait()

		return
	}
	color.New(color.FgRed).Printf("not getting any stream from stdin, you have to run tprof like the following.\n")
	color.New(color.FgGreen).Printf("\t go test <package> -v | tprof\n")

}
