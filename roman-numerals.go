package main

func romanToInt(s string) int {
	//Create map
	romanInts := map[rune]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
	//Initialize counter var
	result := 0
	// Iterate s
	for i := 0; i < len(s); i++ {
		curr := romanInts[rune(s[i])]
		next := romanInts[rune(s[i+1])]
		if curr < next {
			result += curr - next
		} else {
			result += curr
		}
	}
	return result

}
