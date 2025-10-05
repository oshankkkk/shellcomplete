#!/usr/bin/env bash
getuserinput(){
	echo $COMP_LINE
	local suggesions	
	suggesions=$(go run main.go "$COMP_LINE")
	echo "Go output: $suggestions" 
}
complete -D -F getuserinput
