# Copyright (C) 2017 David Capello

all:
	go build

package:
	env GOOS=darwin GOARCH=386 go build -v -o bin/findfiles *.go
	cd bin && zip findfiles-macosx.zip findfiles && rm findfiles && cd ..
	env GOOS=windows GOARCH=386 go build -v -o bin/findfiles.exe *.go
	cd bin && zip findfiles-windows.zip findfiles.exe && rm findfiles.exe && cd ..
	env GOOS=linux GOARCH=386 go build -v -o bin/findfiles *.go
	cd bin && zip findfiles-linux.zip findfiles && rm findfiles && cd ..
