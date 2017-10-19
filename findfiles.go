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
var findDuplicates bool
var patterns Patterns
var dups map[string][]string

func main() {
	flag.BoolVar(&ignoreCase, "i", false, "ignore case")
	flag.BoolVar(&findDuplicates, "d", false, "find duplicates")
	flag.Parse()

	for _, pat := range flag.Args() {
		if ignoreCase {
			pat = "(?i)" + pat
		}
		patterns.append(pat)
	}

	if findDuplicates {
		dups = make(map[string][]string)
	}

	filepath.Walk(".", func(fn string, info os.FileInfo, err error) error {
		if info.IsDir() { return nil }

		if patterns.match(path.Base(fn)) >= 0 {
			if findDuplicates {
				// Add candidate to calculate duplicate
				sha1 := fileSha1(fn)
				dups[sha1] = append(dups[sha1], fn)
			} else {
				fmt.Printf("./%s\n", fn)
			}
		}
		return nil
	})

	if findDuplicates {
		for _, files := range dups {
			if len(files) > 1 {
				fmt.Printf("%q\n", files)
			}
		}
	}
}
