package main

import (
	"fmt"
	"os"
	"sort"

	"github.com/spf13/pflag"
)

var (
	countByBytes = pflag.BoolP("bytes", "c", false, "count by bytes")
	countByChars = pflag.BoolP("chars", "m", false, "count by chars")
	countByWords = pflag.BoolP("words", "w", false, "count by words")
	countByLines = pflag.BoolP("lines", "l", false, "count by lines")
	files0From   = pflag.String("files0-from", "", "files from file, - for stdin")
	sortBy       = pflag.String("sort", "fname", "sort by: chars, words, lines, fname")
)

func init() {
	pflag.Parse()

	if !*countByLines && !*countByBytes && !*countByChars {
		*countByWords = true
	}
}

func main() {
	var (
		files []string
		err   error
	)

	files = []string{"/dev/stdin"}
	if len(*files0From) != 0 {
		files, err = collectFiles(*files0From)
		if err != nil {
			fmt.Fprintf(os.Stderr, "collect files error: %v\n", err)
			os.Exit(1)
		}
	}

	var (
		infos = FileStat{}
	)
	for _, f := range files {
		info, err := countFile(f)
		if err != nil {
			fmt.Fprintf(os.Stderr, "count error, file: %s, err: %v", f, err)
		}
		//fmt.Printf("info: %v\n", info)
		infos = append(infos, info)
	}

	sort.Sort(infos)
	infos.Dump()
}
