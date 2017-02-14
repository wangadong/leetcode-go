package algorithms

func CanCross(stones []int) bool {
	length := len(stones)
	table := make([][]bool, length)
	for x := 0; x < len(table); x++ {
		table[x] = make([]bool, length)
	}
	if stones[1] != 1 {
		return false
	}
	table[0][1] = true

	for i := 0; i < length; i++ {
		flag := false
		for j := i + 1; j < length; j++ {
			units := stones[j] - stones[i]
			for k := 0; k < i; k++ {
				lastUnits := stones[i] - stones[k]
				if table[k][i] == true && (lastUnits == units-1 || lastUnits == units || lastUnits == units+1) {
					table[i][j] = true
					flag = true
				}
			}
			if j == length-1 && table[i][j] == true {
				return true
			}
		}
		if !flag && i != 0 {
			return false
		}
	}
	return false

}
