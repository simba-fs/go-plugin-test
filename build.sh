#!/bin/bash 

for i in $(\ls plugins); do
	cmd=$(basename $i .go)
	echo building $cmd
	go build -o dist/$cmd.so -buildmode=plugin ./plugins/$cmd.go
done

