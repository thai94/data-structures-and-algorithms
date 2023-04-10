package main

import (
	"strconv"
	"strings"
)

func getOpenApplications(commands []string) []string {

	var ret []string
	for _, v := range commands {

		if v == "clear" {
			ret = []string{}
		} else {
			cmd := strings.Split(v, " ")
			if cmd[0] == "open" {
				ret = append(ret, cmd[1])
			} else {
				num, _ := strconv.Atoi(cmd[1])
				if num >= len(ret) {
					ret = []string{}
				} else {
					ret = ret[:len(ret)-num]
				}
			}
		}
	}

	return ret
}

func main() {

}
