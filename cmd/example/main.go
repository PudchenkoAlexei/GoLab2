package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	Lab_2 "github.com/PudchenkoAlexei/GoLang_Lab2"
)

func main() {

	var eFlag = flag.String("e", "", "input from console")
	var fFlag = flag.String("f", "", "input from file")
	var oFlag = flag.String("o", "", "output file")

	flag.Parse()

	if *eFlag == "" && *fFlag == "" {
		fmt.Fprintln(os.Stderr, "error: either -e or -f flag must be used")
		os.Exit(1)
	}

	if *eFlag != "" && *fFlag != "" {
		fmt.Fprintln(os.Stderr, "error: either -e or -f must be used")
		os.Exit(1)
	}
	var handler = Lab_2.ComputeHandler{
		Input:  strings.NewReader(*eFlag),
		Output: os.Stdout,
	}

	if *eFlag == "" && *fFlag != "" {
		inFile, err := os.Open(*fFlag)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error: %v\n", err)
			os.Exit(1)
		}

		defer inFile.Close()
		handler.Input = inFile
	}

	if *oFlag != "" {
		outfile, err := os.Create(*oFlag)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error: %v\n", err)
			os.Exit(1)
		}

		handler.Output = outfile
		defer outfile.Close()
	}

	err := handler.Compute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %s\n", err)
		os.Exit(1)
	}
}
