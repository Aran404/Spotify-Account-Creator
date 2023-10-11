package utils

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

func Run(name string, arg ...string) {
	cmd := exec.Command(name, arg...)
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func Clear() {
	switch runtime.GOOS {
	case "darwin":
		Run("clear")
	case "linux":
		Run("clear")
	case "windows":
		Run("cmd", "/c", "cls")
	default:
		Run("clear")
	}
}

func SetTitle(format string, content ...interface{}) {
	switch runtime.GOOS {
	case "darwin":
		Run("printf", "\033]0;"+fmt.Sprintf(format, content...)+"\007")
	case "linux":
		Run("printf", "\033]0;"+fmt.Sprintf(format, content...)+"\007")
	case "windows":
		Run("cmd", "/c", "title", fmt.Sprintf(format, content...))
	default:
		fmt.Printf(format, content...)
	}
}
