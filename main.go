package main

import (
	"flag"
	"fmt"
	"os"

	template "text/template"

	"github.com/Masterminds/sprig"
)

func main() {
	in := flag.String("in", "", "the template to render")
	flag.Parse()
	if *in == "" {
		fmt.Fprintln(os.Stderr, "Error: the input file was empty")
		flag.Usage()
		os.Exit(1)
	}

	tpl, err := template.ParseFiles(*in)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: parsing template %s (%s)\n", *in, err)
		os.Exit(1)
	}

	envMap := collectEnv()

	if err := tpl.Funcs(sprig.TxtFuncMap()).Execute(os.Stdout, envMap); err != nil {
		fmt.Fprintf(os.Stderr, "Rendering template (%s)", err)
		os.Exit(1)
	}
}
