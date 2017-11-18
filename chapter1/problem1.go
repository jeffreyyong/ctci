package chapter1

// Using map of runes for duplicate detection.
func IsUnique(input string) bool {
	seen := make(map[rune]struct{})
	for _, r := range input {
		if _, ok := seen[r]; ok {
			return false
		} else {
			seen[r] = struct{}{}
		}
	}
	return true
}

/* Note:
   Time complexity for this code is O(n), where n is the length of the string.
   The space complexity is O(1)
   (Could argue that the time complexity is O(1), since the for loop will never iterate through more than 128 characters).
*/

/*
Alternative solutions:
1) Compare every character of the string to every other character of the strnig. This will take O(N2) time and O(1) space.
2) If allowed to modify the input string, coudl sort the string in O(nlog(n)) time and then linearly check the string
	 for neighboring characters that are identical. Careful though: many sorting algorithms take up extra space.
*/
