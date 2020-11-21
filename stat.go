package main

import (
	"bytes"
	"fmt"
	"os"
	"strings"
	"text/tabwriter"
)

type FileStats struct {
	FilePath     string
	CountByBytes int
	CountByChars int
	CountByWords int
	CountByLines int
}

type FileStat []*FileStats

func (infs FileStat) Dump() {
	if len(infs) == 0 {
		return
	}
	verbose := false
	if len(infs) > 1 {
		verbose = true
	}

	w := tabwriter.NewWriter(os.Stdout, 0, 8, 0, '\t', tabwriter.DiscardEmptyColumns)

	sum := FileStats{FilePath: "total"}

	buf := bytes.Buffer{}
	for _, f := range infs {
		if *countByLines {
			buf.WriteString(fmt.Sprintf("%d\t", f.CountByLines))
			sum.CountByLines += f.CountByLines
		}
		if *countByWords {
			buf.WriteString(fmt.Sprintf("%d\t", f.CountByWords))
			sum.CountByWords += f.CountByWords
		}
		if *countByChars {
			buf.WriteString(fmt.Sprintf("%d\t", f.CountByChars))
			sum.CountByChars += f.CountByChars
		}
		if *countByBytes {
			buf.WriteString(fmt.Sprintf("%d\t", f.CountByBytes))
			sum.CountByBytes += f.CountByBytes
		}
		if verbose {
			buf.WriteString(f.FilePath + "\t")
		}
		buf.WriteString("\n")
	}

	if verbose {
		if *countByLines {
			buf.WriteString(fmt.Sprintf("%d\t", sum.CountByLines))
		}
		if *countByWords {
			buf.WriteString(fmt.Sprintf("%d\t", sum.CountByWords))
		}
		if *countByChars {
			buf.WriteString(fmt.Sprintf("%d\t", sum.CountByChars))
		}
		if *countByBytes {
			buf.WriteString(fmt.Sprintf("%d\t", sum.CountByBytes))
		}
		buf.WriteString("total\t\n")
	}
	w.Write(buf.Bytes())
	w.Flush()
}

func (infs FileStat) Len() int {
	return len(infs)
}

type sortManner string

const (
	sortByFileName = sortManner("fname")
	sortByChars    = sortManner("chars")
	sortByWords    = sortManner("words")
	sortByLines    = sortManner("lines")
)

func (infs FileStat) Less(i, j int) bool {
	switch sortManner(*sortBy) {
	case sortByFileName:
		return strings.Compare(infs[i].FilePath, infs[j].FilePath) < 0
	case sortByChars:
		return infs[i].CountByChars < infs[j].CountByChars
	case sortByWords:
		return infs[i].CountByWords < infs[j].CountByWords
	case sortByLines:
		return infs[i].CountByLines < infs[j].CountByLines
	default:
		panic("invalid sort manner")
	}

}

func (infs FileStat) Swap(i, j int) {
	infs[i], infs[j] = infs[j], infs[i]
}
