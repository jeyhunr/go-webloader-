package main

import (
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

var (
	flagOutput = flag.String("o", "", "output file")
	flagHeader = flag.Bool("header", false, "print HTTP-header")
)

func main() {
	flag.Parse()
	args := flag.Args()
	if len(args) != 1 {
		fmt.Println("please enter only one url")
		os.Exit(1)
	}
	url := args[0]
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("error reading %s\n%#v", url, err)
	}
	defer resp.Body.Close()
	var w io.Writer
	w = os.Stdout

	if *flagOutput != "" {
		// if the directory does not exist
		err := os.MkdirAll(filepath.Dir(*flagOutput), 0755)
		if err != nil {
			fmt.Printf("error creating directory: %v", err)
			os.Exit(1)
		}
		f, err := os.OpenFile(
			*flagOutput,
			os.O_RDWR|os.O_CREATE,
			0755)
		if err != nil {
			fmt.Println("error reading %s\n%#v", *flagOutput, err)
			os.Exit(1)
		}
		defer f.Close()
		w = f
	}

	if *flagHeader {
		for k, v := range resp.Header {
			fmt.Fprintf(w, "%s :\n", k) // key
			for i, l := range v {
				// row
				fmt.Fprintf(w, " %03d: %s \n", i+1, l)
			}
		}
		os.Exit(0)
	}
	io.Copy(w, resp.Body)
}
