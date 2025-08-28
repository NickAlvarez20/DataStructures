package main

// Intuition: The problem is to find the first occurence of a substring(needle) within another string(haystack). If the substring exists, return its starting index; otherwise return -1
//Key Observation: We can compare slices of the haystack with the needle to check for equality matches. By iterating through the haystack and increment the slice by one, we can using a sliding window approach to scan each possible position the needle could fit. This approach avoids unneccesary compute as only only compare substrings of the same length as the needle. 

/*Algorithm Design
1. If the needle is an empty string, return 0 because an empty string is considered to be present at the start of any string
2. Iterate through haystack.
 -Use a loop to iterate through all possible starting indices in the haystack where the needle could fit
 -The valid range for the starting index is from 0 to len(haystack)-len(needle) (inclusive)
 3. Substring comparison
 -For each starting index, i , extract the substring with a slice of the haystack/ haystack[i:i+len(needle)]
 -Compare this slice/substring with the needle, if they match, return the current index, i.
 4. If no match is found after iterating the entire string, return -1
 */

func strStr(haystack string, needle string) int {
    // Determine the Edge Cases 
    if len(needle) == 0{
        return 0
    }
    // Iterate through haystack
    for i := 0; i <= len(haystack)-len(needle); i++{
        if haystack[i:i+len(needle)] == needle{
            return i
        }
    }
    return -1
}

/* Algorithm Analysis
1. Time Complexity
-O(n*m)
-Worst Case: We compare the needle with every possible substring that has the same length as needle
-Each comparison takes O(m), where m is the length of the needle
-The number of comparisions is O(n-m+1), where n is the length of the haystack
-Overall time complexity is O((n-m+1)*m), which simplifies to O(n*m) in the worst case 
2. Time complexity
-Algorithim uses constant amount of extra space for variables like the loop index
-Slicing the haystack does not create a new string in Go(it creates a reference), no no additional space is used for slicing
-O(1)

Optimizations:
To optimize the strStr function, consider these key points:

Use KMP Algorithm: Reduces time complexity to O(n + m) by preprocessing the needle to avoid redundant comparisons.
Apply Boyer-Moore Algorithm: Leverages bad character and good suffix rules for potentially faster matching, especially with large alphabets.
Use Two-Pointer Technique: Avoids slicing by comparing characters directly, reducing overhead.
Pre-check Length: Explicitly check if len(needle) > len(haystack) to return -1 early.
Use Built-in Functions: In Go, consider strings.Index for optimized native implementation.
*/