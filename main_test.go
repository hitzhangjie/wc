package main

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_collectFiles(t *testing.T) {
	var dat = "a.go\000b.go\000c/d/e.go\000"
	files, err := readFiles(bytes.NewBufferString(dat))
	t.Logf("readFiles, err:%v, files:%v", err, files)

	assert.Nil(t, err)
	assert.NotNil(t, files)
	assert.Len(t, files, 3)
	assert.Equal(t, "a.go", files[0])
	assert.Equal(t, "b.go", files[1])
	assert.Equal(t, "c/d/e.go", files[2])
}

func TestCountFile(t *testing.T) {
	info, err := countFile("go.sum")
	t.Logf("count file, info: %v, err: %v", info, err)
	assert.Nil(t, err)
	assert.NotNil(t, info)
}
