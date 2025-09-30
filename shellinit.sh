#!/usr/bin/env bash
gettab(){
	echo "COMP_LINE : $COMP_LINE"
	echo "COMP_WORDS : ${COMP_WORDS[*]}"
	echo "COMP_CWORDS: $COMP_CWORDS"
}
complete -F gettab mycmd 
