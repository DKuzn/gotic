#!/bin/bash

archs=(amd64 arm64)

for arch in ${archs[@]};do
	env GOOS=linux GOARCH=${arch} go build -o ./dist/gotic-${VERSION}-linux-${arch}/gotic
	env GOOS=darwin GOARCH=${arch} go build -o ./dist/gotic-${VERSION}-macos-${arch}/gotic
done

env GOOS=windows GOARCH=amd64 go build -o ./dist/gotic-${VERSION}-windows-amd64/gotic.exe

cd ./dist

dirslist=(`ls`)

for dirsitem in ${dirslist[@]}; do
	if [[ $dirsitem == *"linux"* ]]; then
		tar -czf ${dirsitem}.tar.gz $dirsitem
	else
		zip -r ${dirsitem}.zip $dirsitem
	fi
done