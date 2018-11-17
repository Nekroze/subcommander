package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/spf13/afero"
)

func main() {
	err := newExecution(afero.NewOsFs(), getSubcommandsPath(), os.Args[1:]).Run()
	if err != nil {
		panic(err)
	}
}
func getSubcommandsPath() (out string) {
	out = os.Getenv("SUBCOMMANDS")
	if out == "" {
		out = "subcommands"
	}
	return out
}

type execution struct {
	path string
	args []string
	file bool
}

func newExecution(fs afero.Fs, root string, args []string) (out execution) {
	scope := []string{root}
	out.args = []string{}
	for i, part := range args {
		newscope := append(scope, part)
		path := filepath.Join(newscope...)
		if isExecutableFile(fs, path) {
			out.path = path
			out.file = true
			out.args = args[i+1:]
			break
		}
		scope = newscope
		out.args = args[i:]
		if !isDir(fs, path) {
			break
		}
		out.path = path
	}
	return out
}

func (e execution) Run() (err error) {
	if e.file {
		c := exec.Command(e.path, e.args...)
		err = c.Run()
	} else {
		if len(e.args) > 0 {
			if help[e.args[0]] {
				fmt.Println("help", e.path)
			} else if version[e.args[0]] {
				fmt.Println("version", e.path)
			}
		}
	}
	return nil
}

var help = map[string]bool{
	"help":   true,
	"--help": true,
	"-h":     true,
}

var version = map[string]bool{
	"version":   true,
	"--version": true,
	"-v":        true,
}

func isDir(fs afero.Fs, path string) bool {
	file, err := fs.Open(path)
	if err != nil {
		return false
	}
	defer file.Close()

	fi, err := file.Stat()
	if err != nil {
		return false
	}
	return fi.IsDir()
}

func isExecutableFile(fs afero.Fs, path string) bool {
	file, err := fs.Open(path)
	if err != nil {
		return false
	}
	defer file.Close()

	fi, err := file.Stat()
	if err != nil || fi.IsDir() {
		return false
	}
	return fi.Mode().String()[9] == 'x'
}
