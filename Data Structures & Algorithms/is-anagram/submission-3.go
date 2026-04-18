func isAnagram(s string, t string) bool {
	seen := make(map[rune]int)

	for _, ch := range s {
		seen[ch]++
	}

	for _, ch := range t {
		if n, ok := seen[ch]; ok && n > 0 {
			seen[ch]--

			// no more
			if seen[ch] == 0 {
				delete(seen, ch)
			}
		} else {
			return false
		}
	}

	return len(seen) == 0
}
