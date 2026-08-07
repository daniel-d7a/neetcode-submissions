func groupAnagrams(strs []string) [][]string {

	anagramMap := make(map[[26]int][]string)

	for _, val := range strs {
		charMap := [26]int{}

		for i := range val {
			charMap[val[i] - 'a']++
		}

		anagramMap[charMap] = append(anagramMap[charMap], val)
	}

	answer := [][]string {}

	for key := range anagramMap {
		answer = append(answer, anagramMap[key])
	}
	return answer

}
