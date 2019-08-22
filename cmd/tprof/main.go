package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/gokultp/go-tprof/internal/parser"
)

func main() {
	pkgs := parser.ParseTestReport(os.Stdin)
	data, err := json.Marshal(pkgs)

	fmt.Println(err, string(data))
}
