package pkg

// 给你一个字符串数组，请你将 字母异位词 组合在一起。可以按任意顺序返回结果列表。

// 字母异位词 是由重新排列源单词的所有字母得到的一个新单词。

// 示例 1:

// 输入: strs = ["eat", "tea", "tan", "ate", "nat", "bat"]
// 输出: [["bat"],["nat","tan"],["ate","eat","tea"]]

func GroupAnagrams(strs []string) [][]string {

	hashTable := make(map[[26]int][]string)

	for i := 0; i < len(strs); i++ {
		cnt := [26]int{}
		for _, str := range strs[i] {
			cnt[str-'a']++
		}

		hashTable[cnt] = append(hashTable[cnt], strs[i])
	}

	res := [][]string{}
	for _, v := range hashTable {
		res = append(res, v)
	}

	return res
}
