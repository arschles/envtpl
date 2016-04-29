package main

import (
	"bytes"
	"os"
	"path/filepath"
	"testing"

	"github.com/arschles/assert"
)

func testDataDir() (string, error) {
	wd, err := os.Getwd()
	if err != nil {
		return "", err
	}
	return filepath.Join(wd, "testdata"), nil
}

func TestCreateAndRenderTpl(t *testing.T) {
	td, err := testDataDir()
	assert.NoErr(t, err)
	templateNames := []string{
		"sprig.tpl",
		"pwd.tpl",
	}
	envs := collectEnv()
	for i, tplName := range templateNames {
		fName := filepath.Join(td, tplName)
		tpl, err := createTpl(filepath.Base(fName), fName)
		if err != nil {
			t.Errorf("Error creating template %d (%s)", i, err)
			continue
		}
		buf := new(bytes.Buffer)
		if err := renderTpl(tpl, buf, envs); err != nil {
			t.Errorf("Error rendering template %d (%s)", i, err)
		}
	}
}
