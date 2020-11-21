package main

import (
	"bufio"
	"io"
	"os"
)

const (
	null = byte('\000')
)

func collectFiles(filepath string) ([]string, error) {
	fin := os.Stdin
	if filepath != "-" {
		// 从其他文件读取
		f, err := os.Open(filepath)
		if err != nil {
			return nil, err
		}
		fin = f
	}
	return readFiles(fin)
}

func readFiles(reader io.Reader) ([]string, error) {
	files := []string{}

	r := bufio.NewReader(reader)
	for {
		v, err := r.ReadString(null)
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, err
		}
		files = append(files, v[0:len(v)-1])
	}
	return files, nil
}
