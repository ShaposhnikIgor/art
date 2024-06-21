package cypher

import (
	"strings"
)

// if_decod checks if the given message is already decoded.
// It takes a string message as input and returns true if the message is already decoded, false otherwise.
func If_Decod(message string) bool {
	// Check if the message contains both '[' and ']'.
	if strings.Contains(message, "[") && strings.Contains(message, "]") {
		return false // If both '[' and ']' are present, the message is not decoded.
	} else {
		return true // If either '[' or ']' is missing, the message is considered decoded.
	}
}

// isBalanced checks if the square brackets in the given message are balanced.
// It takes a byte slice message as input and returns true if the square brackets are balanced, false otherwise.
func IsBalanced(message []byte) bool {
	// Initializing a variable to keep track of the count of '[' characters.
	count := 0
	// Looping through each character in the message.
	for _, char := range message {
		// If the current character is '[', increment the count.
		if char == '[' {
			count++
		} else if char == ']' { // If the current character is ']', decrement the count.
			count--
		}
		// If the count becomes negative at any point, indicating unbalanced brackets, return false.
		if count < 0 {
			return false
		}
	}
	// If the count is zero at the end of the loop, indicating balanced brackets, return true.
	return count == 0
}

// if_encod checks if the given message is already encoded.
// It takes a string message as input and returns true if the message is already encoded, false otherwise.
func if_encod(message string) bool {
	// Check if the message does not contain both '[' and ']'.
	if !strings.Contains(message, "[") && !strings.Contains(message, "]") {
		return false // If neither '[' nor ']' is present, the message is not encoded.
	} else {
		return true // If either '[' or ']' is present, the message is considered encoded.
	}
}
