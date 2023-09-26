package alg

func FindAnagrams(s string, p string) []int {
	res := []int{}

	destAnagrams := [26]int{}

	n := len(p)

	for i := 0; i < n; i++ {
		destAnagrams[p[i]-'a']++
	}

	for i := 0; i < len(s)-n; i++ {
		nowAnagrams := [26]int{}

		for j := i; j < i+n; j++ {
			nowAnagrams[s[j]-'a']++
		}

		if destAnagrams == nowAnagrams {
			res = append(res, i)
		}

	}

	return res

}
