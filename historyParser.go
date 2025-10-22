package main

import (
	"sort"
)

type sortedcmds struct {
	cmd string
	distance int
}
func sortbashhistory(usercmd string, bashhistory []string)[]string{
	var sortedbashhistory []sortedcmds

	for _, cmds := range bashhistory {
		sortedbashhistory = append(sortedbashhistory,sortedcmds{
			cmd:      cmds,
			distance: editDistance(cmds, usercmd),
		})
	}
	sort.Slice(sortedbashhistory, func(i, j int) bool {
		return sortedbashhistory[i].distance < sortedbashhistory[j].distance
	})

	result := []string{}
	for _, sc := range sortedbashhistory {
		result = append(result, sc.cmd)
	}

	return result
}


