package main

import (
	"os"
	"testing"

	"github.com/arschles/assert"
)

func TestCollectEnv(t *testing.T) {
	envMap := collectEnv()
	pwd, ok := envMap["PWD"]
	assert.True(t, ok, "'PWD' not found in the env")
	wd, err := os.Getwd()
	assert.NoErr(t, err)
	assert.Equal(t, pwd, wd, "working dir")
}
