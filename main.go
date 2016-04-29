package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	in := flag.String("in", "", "the template to render")
	flag.Parse()
	if *in == "" {
		fmt.Fprintln(os.Stderr, "Error: the input file was empty")
		flag.Usage()
		os.Exit(1)
	}

	tpl, err := createTpl(filepath.Base(*in), *in)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: parsing template %s (%s)\n", *in, err)
		os.Exit(1)
	}

	envMap := collectEnv()

	if err := renderTpl(tpl, os.Stdout, envMap); err != nil {
		fmt.Fprintf(os.Stderr, "Rendering template (%s)", err)
		os.Exit(1)
	}
}
