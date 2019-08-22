package browser

import (
	"fmt"
	"os/exec"
	"runtime"
)

// Open will open the given url in browser
// it is supported linux, windows and mac
func Open(url string) error {
	switch runtime.GOOS {
	case "linux":
		return exec.Command("xdg-open", url).Start()
	case "windows":
		return exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		return exec.Command("open", url).Start()
	default:
		return fmt.Errorf("unsupported platform could not open browser")
	}

}
