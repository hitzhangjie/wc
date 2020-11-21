package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
)

func countFile(filepath string) (*FileStats, error) {
	fin := os.Stdin
	if filepath != "/dev/stdin" {
		inf, err := os.Lstat(filepath)
		if err != nil || inf.IsDir() {
			return nil, fmt.Errorf("invalid file entry")
		}
		f, err := os.Open(filepath)
		if err != nil {
			return nil, err
		}
		fin = f
	}

	// TODO customize a ScanFunc to read and count only once

	dat, err := ioutil.ReadAll(fin)
	if err != nil {
		return nil, err
	}

	info := FileStats{
		FilePath: filepath,
	}

	// count chars, wc -m
	sc := bufio.NewScanner(bytes.NewBuffer(dat))
	sc.Split(bufio.ScanRunes)
	for sc.Scan() {
		info.CountByChars++
	}

	// count words, wc -w
	sc = bufio.NewScanner(bytes.NewBuffer(dat))
	sc.Split(bufio.ScanWords)
	for sc.Scan() {
		info.CountByWords++
	}

	// count lines, wc -l
	sc = bufio.NewScanner(bytes.NewBuffer(dat))
	sc.Split(bufio.ScanLines)
	for sc.Scan() {
		info.CountByLines++
	}

	// count bytes, wc -c
	sc = bufio.NewScanner(bytes.NewBuffer(dat))
	sc.Split(bufio.ScanBytes)
	for sc.Scan() {
		info.CountByBytes++
	}

	return &info, nil
}
