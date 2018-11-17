package main

import (
	"strings"
	"testing"

	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"
)

func Test_newExecution(t *testing.T) {
	fs := afero.NewMemMapFs()
	base := "/subcommands"
	fs.MkdirAll(base+"/deep", 0755)
	afero.WriteFile(fs, base+"/a", []byte("file a"), 0755)
	afero.WriteFile(fs, base+"/b", []byte("file b"), 0644)
	afero.WriteFile(fs, base+"/deep/c", []byte("file c"), 0755)

	type args struct {
		fs   afero.Fs
		root string
		args []string
	}

	type run struct {
		name    string
		args    args
		wantOut execution
	}
	genTest := func(command string, expected execution) run {
		inputs := strings.Split(command, " ")
		if len(strings.TrimSpace(command)) == 0 {
			inputs = []string{}
			command = "<nil>"
		}
		return run{
			name: command,
			args: args{
				fs:   fs,
				root: "/subcommands",
				args: inputs,
			},
			wantOut: expected,
		}
	}

	tests := []run{
		genTest("", execution{
			path: "",
			args: []string{},
			file: false,
		}),

		genTest("--help", execution{
			path: "",
			args: []string{"--help"},
			file: false,
		}),

		genTest("b", execution{
			path: "",
			args: []string{"b"},
			file: false,
		}),

		genTest("a", execution{
			path: "/subcommands/a",
			args: []string{},
			file: true,
		}),

		genTest("a foo", execution{
			path: "/subcommands/a",
			args: []string{"foo"},
			file: true,
		}),

		genTest("a -h", execution{
			path: "/subcommands/a",
			args: []string{"-h"},
			file: true,
		}),

		genTest("a bar help", execution{
			path: "/subcommands/a",
			args: []string{"bar", "help"},
			file: true,
		}),

		genTest("deep c", execution{
			path: "/subcommands/deep/c",
			args: []string{},
			file: true,
		}),

		genTest("deep c foo", execution{
			path: "/subcommands/deep/c",
			args: []string{"foo"},
			file: true,
		}),

		genTest("deep c -h", execution{
			path: "/subcommands/deep/c",
			args: []string{"-h"},
			file: true,
		}),

		genTest("deep c bar help", execution{
			path: "/subcommands/deep/c",
			args: []string{"bar", "help"},
			file: true,
		}),

		genTest("deep help", execution{
			path: "/subcommands/deep",
			args: []string{"help"},
			file: false,
		}),

		genTest("deep bar --help", execution{
			path: "/subcommands/deep",
			args: []string{"bar", "--help"},
			file: false,
		}),
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.wantOut, newExecution(tt.args.fs, tt.args.root, tt.args.args))
		})
	}
}
