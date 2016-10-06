package algorithms

// Use Dynamic Programming Table Algo
func IsInterleave(s1 string, s2 string, s3 string) bool {
	if len(s3) != len(s1)+len(s2) {
		return false
	}
	table := make([][]bool, len(s1)+1)
	for i := 0; i < len(table); i++ {
		table[i] = make([]bool, len(s2)+1)
	}
	for i := 0; i < len(s1)+1; i++ {
		for j := 0; j < len(s2)+1; j++ {
			if i == 0 && j == 0 {
				table[i][j] = true
			} else if i == 0 {
				table[i][j] = (table[i][j-1] && s2[j-1] == s3[i+j-1])
			} else if j == 0 {
				table[i][j] = (table[i-1][j] && s1[i-1] == s3[i+j-1])
			} else {
				table[i][j] = (table[i-1][j] && s1[i-1] == s3[i+j-1]) || (table[i][j-1] && s2[j-1] == s3[i+j-1])

			}
		}
	}

	return table[len(s1)][len(s2)]
}
