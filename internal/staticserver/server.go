package staticserver

import (
	"fmt"
	"net/http"
	"strings"
	"sync"

	"github.com/fatih/color"
	_ "github.com/gokultp/go-tprof/internal/statik"
	"github.com/rakyll/statik/fs"
)

const (
	dataPlaceHolder = "{{data}}"
	indexFile       = "/index.html"
)

// Start will start the server
func Start(data, port string, wg *sync.WaitGroup) error {
	defer wg.Done()
	statikFS, err := fs.New()
	if err != nil {
		return err
	}
	strHTML, err := getHTMLContent(statikFS, data)
	if err != nil {
		return err
	}

	http.Handle("/_nuxt/", http.FileServer(statikFS))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, strHTML)
	})
	if err := printHelpMsg(port); err != nil {
		return err
	}
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		return err
	}

	return nil

}

func getHTMLContent(hfs http.FileSystem, data string) (string, error) {
	d, err := fs.ReadFile(hfs, indexFile)
	if err != nil {
		return "", err
	}
	return strings.Replace(string(d), dataPlaceHolder, data, 1), nil
}
func printHelpMsg(port string) error {
	clrPrinter := color.New(color.FgMagenta)
	if _, err := clrPrinter.Printf("\n\n\n\n\t Visit http://localhost:%s to see the reports", port); err != nil {
		return err
	}
	if _, err := clrPrinter.Printf("\n\n\t Press Ctrl + C to exit"); err != nil {
		return err
	}
	return nil
}
