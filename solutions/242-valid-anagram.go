package solutions

func IsAnagram(s string, t string) bool {
	// O(s,n) linear space
	var map1 = make(map[string]int)
	var map2 = make(map[string]int)

	if len(s) != len(t) {
		return false
	}

	// O(s,t) linear time
	for i := 0; i < len(s); i++ {
		charS := string(s[i])
		charT := string(t[i])
		map1[charS]++
		map2[charT]++
	}

	// O(s,t) linear time
	for key, value := range map1 {
		if map2[key] != value {
			return false
		}
	}

	return true

}
