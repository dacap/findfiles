// Copyright (C) 2017 David Capello
//
// This file is released under the terms of the MIT license.
// Read LICENSE.txt for more information.

package main

import (
	"flag"
	"fmt"
	"os"
	"path"
	"path/filepath"
)

var ignoreCase bool
var patterns Patterns

func main() {
	flag.BoolVar(&ignoreCase, "i", false, "ignore case")
	flag.Parse()

	for _, pat := range flag.Args() {
		if ignoreCase {
			pat = "(?i)" + pat
		}
		patterns.append(pat)
	}

	filepath.Walk(".", func(fn string, info os.FileInfo, err error) error {
		if patterns.match(path.Base(fn)) >= 0 {
			fmt.Printf("./%s\n", fn)
		}
		return nil
	})
}
