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

func StartServer(data, port string, wg *sync.WaitGroup) error {
	defer wg.Done()
	statikFS, err := fs.New()
	if err != nil {
		return err
	}
	d, err := fs.ReadFile(statikFS, "/index.html")
	if err != nil {
		return err
	}
	strHTML := string(d)
	strHTML = strings.Replace(strHTML, "{{data}}", data, 1)
	http.Handle("/_nuxt/", http.FileServer(statikFS))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, strHTML)
	})

	color.New(color.FgMagenta).Printf("\n\n\n\n\t Visit http://localhost:6969 to see the reports \n\n\t Press Ctrl + C to exit")
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		return err
	}

	return nil

}
