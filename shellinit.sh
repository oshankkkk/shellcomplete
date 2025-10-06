#!/usr/bin/env bash
getuserinput(){
	s=$(go run . "$COMP_LINE")
#	echo "Go output: $s"
}
complete -D -F getuserinput
