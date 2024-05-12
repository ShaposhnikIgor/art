package main

import (
	"log"
	"runtime"
	"strconv"
	"strings"
)

// decode_Art decodes the given message encoded in the ART cipher.
// It takes a string message as input and returns the decoded message as a string.
func decode_Art(message string) string {
	// Check if the data is already decoded.
	if if_decod(message) {
		return "data decoded, try again" // Return a message indicating that the data is already decoded.
	}

	// Check if the square brackets in the message are balanced.
	if !isBalanced([]byte(message)) {
		return "unbalanced square brackets" // Return a message indicating unbalanced square brackets.
	}

	// Initializing a strings.Builder to efficiently build strings.
	var decodedMessage strings.Builder

	// Variable to keep track of the current position in the input string.
	var i int

	// Loop through the input string.
	for i < len(message) {
		// If the current character is '[', indicating a repeated section.
		if message[i] == '[' {
			// Find the index of the closing ']' for the current section.
			closeBracketIndex := strings.Index(message[i:], "]")
			if closeBracketIndex == -1 { // If ']' is not found, indicating invalid input.
				// Log an error message indicating missing closing bracket.
				_, file, line, _ := runtime.Caller(0)
				log.Printf("%s:%d: error: Missing closing bracket", file, line)
				return "error: Missing closing bracket" // Return an error message.
			}

			closeBracketIndex += i // Adjust the index to get the absolute position of the closing bracket.

			// Split the section into repeat count and content.
			parts := strings.SplitN(message[i+1:closeBracketIndex], " ", 2)
			if len(parts) != 2 || parts[1] == "" { // If the split does not result in two parts or content is empty, indicating invalid input.
				// Log an error message indicating invalid section format.
				_, file, line, _ := runtime.Caller(0)
				log.Printf("%s:%d: error: Invalid section format", file, line)
				return "error: Invalid section format" // Return an error message.
			}

			// Convert repeat count from string to integer.
			repeatCount, err := strconv.Atoi(parts[0])
			if err != nil { // If conversion fails, indicating invalid input.
				// Log an error message indicating invalid repeat count.
				_, file, line, _ := runtime.Caller(0)
				log.Printf("%s:%d: error: Invalid repeat count", file, line)
				return "error: Invalid repeat count" // Return an error message.
			}

			// Repeat the content of the section and append to the decoded message.
			for j := 0; j < repeatCount; j++ {
				decodedMessage.WriteString(parts[1])
			}

			i = closeBracketIndex + 1 // Move the index to the character following the closing bracket.
		} else { // If the current character is not '[', simply append it to the decoded message.
			decodedMessage.WriteByte(message[i])
			i++
		}
	}

	return decodedMessage.String() // Return the decoded message.
}
