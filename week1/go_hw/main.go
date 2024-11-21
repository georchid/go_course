package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
)

func dirTree(out io.Writer, path string, printFiles bool) error {
	type Entry struct {
		Path      string
		Depth     int
		Prefix    string
		LastInDir bool
	}
	queue := []Entry{{Path: path, Depth: -1, Prefix: "├───", LastInDir: false}}

	for len(queue) > 0 {
		current := queue[len(queue)-1]
		queue = queue[:len(queue)-1]
		info, err := os.Stat(current.Path)
		if err != nil {
			fmt.Errorf("[error] os.stat")
		}

		if info.IsDir() {
			if current.Depth != -1 {
				fmt.Fprintf(out, "%s%s\n", current.Prefix, info.Name())
			}

			entries, err := os.ReadDir(current.Path)

			if err != nil {
				fmt.Errorf("[error] os.ReadDir")
			}

			if !printFiles {
				n := 0
				for _, entry := range entries {
					if entry.IsDir() {
						entries[n] = entry
						n++
					}
				}
				entries = entries[:n]
			}

			sort.Slice(entries, func(i, j int) bool {
				return entries[i].Name() > entries[j].Name()
			})

			for idx, entry := range entries {
				var lastInDir bool
				var Pref string
				if current.Depth != -1 {
					if current.LastInDir {
						Pref += strings.TrimSuffix(current.Prefix, "└───") + "\t"
					} else {
						Pref += strings.TrimSuffix(current.Prefix, "├───") + "│\t"
					}
				}
				if idx != 0 {
					lastInDir = false
					Pref += "├───"
				} else {
					lastInDir = true
					Pref += "└───"
				}
				queue = append(
					queue,
					Entry{
						Path:      filepath.Join(current.Path, entry.Name()),
						Depth:     current.Depth + 1,
						Prefix:    Pref,
						LastInDir: lastInDir,
					},
				)
			}
		} else {
			if current.Depth != -1 {
				var sz string
				if info.Size() != 0 {
					sz = strconv.FormatInt(info.Size(), 10) + "b"
				} else {
					sz = "empty"
				}
				fmt.Fprintf(out, "%s%s (%s)\n", current.Prefix, info.Name(), sz)
			}
		}
	}
	return nil
}

func main() {
	out := os.Stdout
	if !(len(os.Args) == 2 || len(os.Args) == 3) {
		panic("usage go run main.go . [-f]")
	}
	path := os.Args[1]
	printFiles := len(os.Args) == 3 && os.Args[2] == "-f"
	err := dirTree(out, path, printFiles)
	if err != nil {
		panic(err.Error())
	}
}
