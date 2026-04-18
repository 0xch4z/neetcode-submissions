import (
	"maps"
	"slices"
)

type anagramKey [26]int

func groupAnagrams(strs []string) [][]string {
	m := make(map[anagramKey][]string)

	for _, str := range strs {
		var key anagramKey
		for _, ch := range str {
			key[ch-'a']++
		}

		m[key] = append(m[key], str)
	}

	return slices.Collect(maps.Values(m))
}
