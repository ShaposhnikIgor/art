package main

import (
	"fmt"
	"strings"
)

// encode_Art function encodes a message using the ART encoding scheme.
// It takes a string message as input and returns the encoded message as a string.
func encode_Art(message string) string {
	// Check if the message is already encoded.
	if if_encod(message) {
		return "data encoded, try again"
	}

	// Initialize a strings.Builder to efficiently build the result string.
	var result strings.Builder

	// Initialize variables to track the count of consecutive characters and the last seen character.
	count := 1
	var lastSymbol rune

	// Iterate over each character in the message.
	for i, char := range message {
		// If the current character is the same as the last seen character, increment the count.
		if char == lastSymbol {
			count++
		} else {
			// If the current character is different from the last seen character,
			// append the encoded representation of the last seen character to the result string.
			if lastSymbol != 0 && i > 0 {
				if count > 1 {
					// If the count is greater than 1, append the count and the character enclosed in square brackets.
					result.WriteString(fmt.Sprintf("[%d %c]", count, lastSymbol))
				} else {
					// If the count is 1, append only the character.
					result.WriteString(fmt.Sprintf("%c", lastSymbol))
				}
			}
			// Update the last seen character and reset the count to 1.
			lastSymbol = char
			count = 1
		}
	}

	// After processing all characters, append the encoded representation of the last seen character to the result string.
	if lastSymbol != 0 {
		if count > 1 {
			// If the count is greater than 1, append the count and the character enclosed in square brackets.
			result.WriteString(fmt.Sprintf("[%d %c]", count, lastSymbol))
		} else {
			// If the count is 1, append only the character.
			result.WriteString(fmt.Sprintf("%c", lastSymbol))
		}
	}

	// Return the encoded message as a string.
	return result.String()
}
