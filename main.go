package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"

  "github.com/someone-stole-my-name/mkdn/formatter"
)

func main() {

	flag.Usage = func() {
		fmt.Printf("Usage: mkdn [options] PEM_FILE\n\n")
		flag.PrintDefaults()
	}

	format := flag.String("format", "go", "DN format")

	flag.Parse()

	if flag.NArg() == 0 {
	    flag.Usage()
	    os.Exit(1)
	}

	fmtter, err := formatter.New(*format)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

  certBytes, err := ioutil.ReadFile(flag.Arg(0))
  if err != nil {
    fmt.Println(err.Error())
    os.Exit(1)
  }

	dn, err := fmtter.Format(string(certBytes))

  if err != nil {
    fmt.Println(err.Error())
    os.Exit(1)
  }

	fmt.Println(dn)
}
