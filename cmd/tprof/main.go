package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/fatih/color"
	"github.com/gokultp/go-tprof/internal/parser"
)

func main() {
	stat, err := os.Stdin.Stat()
	if err != nil {
		panic(err)
	}
	if (stat.Mode() & os.ModeCharDevice) == 0 {
		res := parser.ParseTestReport(os.Stdin)
		data, err := json.Marshal(res)

		fmt.Println(err, string(data))
		return
	}
	color.New(color.FgRed).Printf("not getting any stream from stdin, you have to run tprof like the following.\n")
	color.New(color.FgGreen).Printf("\t go test <package> -v | tprof\n")

}
