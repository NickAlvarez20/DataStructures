package main

// Intuition: Use a stack data structure to keep track of the opening brackets
/* Algorithm: Create a dictionary to map the closing brackets and their corresponding opening brackets for quick lookups. 2. Initialize an empty stack, this will store opening brackets as they appear in the string.
Loop through each character in the string, it it is an opening bracket, push it onto the stack, if it is a closing bracket, check the top of the stack
If the top of the stack is empty of the top element doesn't match the corresponding opening bracket, the string is invalid. Otherwise, pop the top element of the stack, remove the matching bracket.
After the loop completes, if stack not empty, there is an unmatched opening bracket so return false.
otherwise return true */

func isValid(s string) bool {
	// Initialize an empty stack and a hashmap for parenthesis
	stack := []rune{}
	hash := map[rune]rune{'}': '{', ')': '(', ']': '['}
	for _, char := range s {
		if match, found := hash[char]; found {
			// check if length of stack is non empty and matches
			if len(stack) > 0 && stack[len(stack)-1] == match {
				stack = stack[:len(stack)-1]
			} else {
				return false
			}
		} else {
			stack = append(stack, char)
		}

	}
	return len(stack) == 0
}

/* Analysis
Time Complexity
O(n)-where n is the length of the array. Each char is only processed once
Space Complexity
O(n)-The stack may store all possible brackets in the worst case (e.g., all opening brackets )
Optimization
1. Add in early validation checks for odd length arrays and arrays with 0 length.
2.Avoid unneccesary rune conversion and utilize bytes, although this only applies to ASCII characters and not Unicode
3. Create a proallocated stack using a stack := make([]byte, 0, len(s)/2) check to minimize the memory allocation during resizing. 
Overall improvement, time space complexity stays unchanged but memory allocation slightly improved. 
*/