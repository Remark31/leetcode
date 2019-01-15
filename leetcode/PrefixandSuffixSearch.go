package leetcode

type WordFilter struct {
	prefixsuffix map[string]int
}

func Constructor(words []string) WordFilter {
	t := WordFilter{}
	t.prefixsuffix = make(map[string]int)
	for index := range words {
		for i := 0; i <= 10 && i <= len(words[index]); i++ {
			for j := 0; j <= 10 && j <= len(words[index]); j++ {
				key := words[index][0:i] + "#" + words[index][len(words[index])-j:len(words[index])]
				if _, ok := t.prefixsuffix[key]; !ok {
					t.prefixsuffix[key] = index
				} else {
					if index > t.prefixsuffix[key] {
						t.prefixsuffix[key] = index
					}
				}

			}
		}
	}
	return t
}

func (this *WordFilter) F(prefix string, suffix string) int {
	if k, ok := this.prefixsuffix[prefix+"#"+suffix]; ok {
		return k
	} else {
		return -1
	}
}

/**
 * Your WordFilter object will be instantiated and called as such:
 * obj := Constructor(words);
 * param_1 := obj.F(prefix,suffix);
 */
