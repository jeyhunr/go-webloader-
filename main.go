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

	io.Copy(w, resp.Body)
}
