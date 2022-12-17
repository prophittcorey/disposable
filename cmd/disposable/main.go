package main

import (
	"flag"
	"fmt"

	"github.com/prophittcorey/disposable"
)

func main() {
	var domain string
	var domains bool

	flag.StringVar(&domain, "check", "", "an email address or domain to analyze (returns 'true' if it's a known disposable address, 'false' otherwise")
	flag.BoolVar(&domains, "domains", false, "if specified, a list of all known disposable email domains will be printed")

	flag.Parse()

	if len(domain) > 0 {
		if val, err := disposable.Check(domain); err == nil {
			fmt.Println(val)
		}

		return
	}

	if domains {
		for _, address := range disposable.Domains() {
			fmt.Println(address)
		}

		return
	}

	flag.Usage()
}
