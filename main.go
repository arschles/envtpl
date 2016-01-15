package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
	"text/template"
)

func main() {
	in := flag.String("in", "", "the template to render")
	flag.Parse()
	if *in == "" {
		fmt.Println("Error: the input file was empty")
		flag.Usage()
		os.Exit(1)
	}

	tpl, err := template.ParseFiles(*in)
	if err != nil {
		fmt.Printf("Error: parsing template %s (%s)\n", *in, err)
		os.Exit(1)
	}

	envMap := make(map[string]string)
	envStrs := os.Environ()
	for _, envStr := range envStrs {
		spl := strings.Split(envStr, "=")
		if len(spl) != 2 {
			continue
		}
		envMap[spl[0]] = spl[1]
	}

	if err := tpl.Execute(os.Stdout, envMap); err != nil {
		fmt.Printf("Rendering template (%s)", err)
		os.Exit(1)
	}
}
