//Intuition: Lexicographically sorting the array of strings arranges the array for easier parsing; sorting brings strings with similar prefixes closes together.
/* 1. Sort the array and then compare from the start/end
   2. Compare characters and append matching results until a mismatch
   3. Return the result with all matching appended chars
*/
function longestCommonPrefix(strs: string[]): string {
  //Edge case-if strs is empty return
  if (!strs.length) {
    return "";
  }
  strs.sort((a: any, b: any) => a - b); //ascending
  // Assign variables values
  let first = strs[0],
    last = strs[strs.length - 1],
    result = "";
  // Iterate through strs array, comparing the first and last value, append if matching
  for (let i = 0; i < strs.length; i++) {
    if (first[i] === last[i]) {
      result += first[i];
    } else {
      break;
    }
  }
  return result;
}

/* Analysis:
Time Complexity - O(n*logn*m)
1. JS sort algorithm (Tim/QuickSort) results in O(n*logn) and modifies array in place
2. Each comparison takes O(m) in worst case, where m is the longest string being compared
Space Complexity - O(log n)
1. Due to sorting stack
Optimizations
1. Remove sorting operation. Use nested loops to directly compare all matches across all strings. Reduces time complexity from O(n*logn*m) to O(n*m)
*/
