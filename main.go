package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	lines := []string{}
	for {
		lnByte, _, err := reader.ReadLine()
		if errors.Is(err, io.EOF) {
			break
		}
		lines = append(lines, string(lnByte))
	}

	tagStrs := []string{}
	for _, line := range lines {
		if line == "" {
			tagStrs = append(tagStrs, "")
			continue
		}

		lineSplit := strings.Fields(line)
		part1 := lineSplit[0]
		idxs := []int{}
		fcri := 0 // first non whitespace char.
		fcriFound := false
		for i, c := range part1 {
			if !fcriFound {
				fcri = i
				fcriFound = true
				continue
			}

			if i >= 1 && capital(int32(part1[i-1])) {
				continue
			}

			if capital(c) {
				idxs = append(idxs, i)
			}

		}
		idxs = append(idxs, len(part1))

		res := ""
		sp := fcri // starting position.
		for i, idx := range idxs {
			cut := part1[sp:idx]
			sp = idx
			res += strings.ToLower(cut)
			if i < len(idxs)-1 && part1[idx] != '_' {
				res += "_"
			}
		}
		tagStrs = append(tagStrs, fmt.Sprintf("%s\t`json:\"%s\"`", line, res))
	}
	for _, tag := range tagStrs {
		fmt.Println(tag)
	}
}

func capital(c int32) bool {
	if c == '_' || (c >= 65 && c <= 90) {
		return true
	}
	return false
}
