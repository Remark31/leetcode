package leetcode

import (
	"fmt"
	"sort"
)

type wordsindex struct {
	s     string
	begin int
	end   int
}

func findword(s, a, b string, index int) []wordsindex {
	al := len(a)
	ans := []wordsindex{}

	i := 0
	if i+al <= len(s) {
		if s[i:i+al] == a {
			t := wordsindex{
				begin: i + index,
				end:   i + index + al,
				s:     b,
			}
			ans = append(ans, t)
		}
	}

	return ans
}

func findReplaceString(S string, indexes []int, sources []string, targets []string) string {

	repmap := []wordsindex{}
	for i := range indexes {
		subs := S[indexes[i]:len(S)]
		a := findword(subs, sources[i], targets[i], indexes[i])

		repmap = append(repmap, a...)
	}

	sort.Slice(repmap, func(i, j int) bool {
		return repmap[i].begin < repmap[j].begin
	})

	fmt.Println(repmap)

	ans := ""
	sbegin := 0
	for i := range repmap {
		if sbegin > repmap[i].begin {
			return S
		}
		ans += S[sbegin:repmap[i].begin] + repmap[i].s
		sbegin = repmap[i].end
	}

	ans += S[sbegin:len(S)]

	return ans
}
