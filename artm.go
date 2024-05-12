package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
)

// UserChoice struct represents the user's choice during interactions.
type UserChoice struct {
	Choice string // Choice holds the selected option.
	GoBack bool   // GoBack indicates whether to return to the previous step.
}

func main() {

	// Open or create a log file for recording errors and informational messages.
	logFile, err := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal("Failed to open log file:", err)
	}
	defer logFile.Close()

	// Set log output to the log file.
	log.SetOutput(logFile)

	// Parse command-line arguments.
	args := os.Args[1:]
	if len(args) == 0 {
		_, file, line, _ := runtime.Caller(0)
		log.Printf("%s:%d: Empty string", file, line)
		fmt.Println("Empty string")
		os.Exit(0)
	}

	// Check if the "-o" flag is provided.
	if len(args) > 0 && args[0] != "-o" {

		// If no flag is provided, assume the input is the encoded text.
		encodedText := args[0]

		// Decode the text and print the result.
		decodedText := decode_Art(encodedText)
		fmt.Println(decodedText)
	}

	// Parse the "-o" flag to call the main menu.
	mainmenuFlag := flag.Bool("o", false, "Call main menu")
	flag.Parse()

	// If the "-o" flag is provided, display the main menu and handle user interactions.
	if *mainmenuFlag {
		fmt.Println("Welcome to the ART Cypher Tool!")

		for {

			// Get user input for encryption/decryption, input/output methods, and message/file selection.
			toEncrypt, inputMethod, outputMethod, message := getInput()

			if toEncrypt.GoBack {
				continue
			}

			var result string

			// Perform encryption or decryption based on user choices.
			switch inputMethod.Choice {
			case "1":
				if toEncrypt.Choice == "1" {
					result = encode_Art(message)
				} else {
					result = decode_Art(message)
				}
			case "2":
				if toEncrypt.Choice == "1" {
					result = encode_Art(message)
				} else {
					result = decode_Art(message)
				}
			default:
				fmt.Println("Invalid input method selection. Please try again.")
				continue
			}

			// Display or write the result based on the output method selected by the user.
			switch outputMethod.Choice {
			case "1":
				fmt.Println(result)
			case "2":
				writeToFile("output.txt", result)
				fmt.Println("Result has been written to output.txt")
			default:
				fmt.Println("Invalid output method selection. Please try again.")
				continue
			}
			os.Exit(0)
		}
	}
}

// getInput gets user input for encryption/decryption, input/output methods, and message/file selection.
func getInput() (toEncrypt UserChoice, inputMethod UserChoice, outputMethod UserChoice, message string) {

	toEncrypt = getOperation()

	if toEncrypt.GoBack {
		return UserChoice{}, UserChoice{}, UserChoice{}, ""
	}

	inputMethod = getInputMethod()

	if inputMethod.GoBack {
		return UserChoice{}, UserChoice{}, UserChoice{}, ""
	}

	outputMethod = getOutputMethod()

	if outputMethod.GoBack {
		return UserChoice{}, UserChoice{}, UserChoice{}, ""
	}

	// Prompt the user to enter a message or select a file based on the chosen input method.
	if inputMethod.Choice == "1" {
		fmt.Println("Enter the message (press Enter twice to finish):")
		reader := bufio.NewReader(os.Stdin)
		var lines []string
		for {
			line, err := reader.ReadString('\n')
			if err != nil {
				_, file, line, _ := runtime.Caller(0)
				log.Printf("%s:%d: Error reading input:", file, line)
				fmt.Println("Error reading input:", err)
				return UserChoice{}, UserChoice{}, UserChoice{}, ""
			}

			if line == "\n" {
				break
			}
			lines = append(lines, line)
		}

		message = strings.Join(lines, "")
	} else if inputMethod.Choice == "2" {

		// Select a file from the specified directory based on the encryption/decryption choice.
		var directory string
		if toEncrypt.Choice == "1" {
			directory = "decodeArtfile"
		} else {
			directory = "encodeArtfile"
		}
		files := getFileList(directory)
		if len(files) == 0 {
			_, file, line, _ := runtime.Caller(0)
			log.Printf("%s:%d: No files found in the directory.", file, line)
			fmt.Println("No files found in the directory.")
			os.Exit(9)
		}

		message = selectFileAndRead(files, directory)
		if message == "" {
			return UserChoice{}, UserChoice{}, UserChoice{}, ""
		}
	}

	return toEncrypt, inputMethod, outputMethod, message
}

// getOutputMethod prompts the user to select the output method (terminal or file).
func getOutputMethod() UserChoice {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("Select the output method (1/2):")
		fmt.Println("1. Terminal.")
		fmt.Println("2. File.")
		fmt.Println("0. Back to top")
		outputMethodChoice, _ := reader.ReadString('\n')
		outputMethodChoice = strings.TrimSpace(outputMethodChoice)

		switch outputMethodChoice {
		case "1", "2":
			return UserChoice{Choice: outputMethodChoice}
		case "0":
			return UserChoice{GoBack: true}
		default:
			fmt.Println("Invalid output method selection. Please try again.")
		}
	}
}

// getOperation prompts the user to select encryption or decryption.
func getOperation() UserChoice {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("Select operation (1/2):")
		fmt.Println("1. Encode.")
		fmt.Println("2. Decode.")
		opChoice, _ := reader.ReadString('\n')
		opChoice = strings.TrimSpace(opChoice)

		switch opChoice {
		case "1", "2":
			return UserChoice{Choice: opChoice}
		case "0":
			return UserChoice{GoBack: true}
		default:
			fmt.Println("Invalid operation selection. Please try again.")
		}
	}
}

// getInputMethod prompts the user to select the input method (terminal or file).
func getInputMethod() UserChoice {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("Select the input method (1/2):")
		fmt.Println("1. Art_terminal.")
		fmt.Println("2. Art_IO.")
		fmt.Println("0. Back to top")
		inputMethodChoice, _ := reader.ReadString('\n')
		inputMethodChoice = strings.TrimSpace(inputMethodChoice)

		switch inputMethodChoice {
		case "1", "2":
			return UserChoice{Choice: inputMethodChoice}
		case "0":
			return UserChoice{GoBack: true}
		default:
			fmt.Println("Invalid input method selection. Please try again.")
		}
	}
}

// getFileList retrieves a list of files from the specified directory.
func getFileList(directory string) []string {
	var files []string
	fileInfo, err := os.ReadDir(directory)
	if err != nil {
		_, file, line, _ := runtime.Caller(0)
		log.Printf("%s:%d: Error reading directory:", file, line)
		fmt.Println("Error reading directory:", err)
		return files
	}

	for _, file := range fileInfo {
		if !file.IsDir() && (file.Name() == "input.encoded.txt" || file.Name() == "input.art.txt") {
			files = append(files, file.Name())
			break
		}
	}

	for _, file := range fileInfo {
		if file.IsDir() || file.Name() == "input.encoded.txt" || file.Name() == "input.art.txt" {
			continue
		}
		files = append(files, file.Name())
	}

	return files
}

// selectFileAndRead prompts the user to select a file and reads its content.
func selectFileAndRead(files []string, directory string) string {
	fmt.Println("Select a file (or enter 0 to step back):")
	for i, file := range files {
		fmt.Printf("%d. %s\n", i+1, file)
	}
	fmt.Println("0. Back to top")

	for {
		reader := bufio.NewReader(os.Stdin)
		fileNumberInput, _ := reader.ReadString('\n')
		fileNumberInput = strings.TrimSpace(fileNumberInput)

		if fileNumberInput == "0" {
			return ""
		}

		fileNumber, err := strconv.Atoi(fileNumberInput)
		if err != nil || fileNumber < 1 || fileNumber > len(files) {
			fmt.Println("Invalid file number. Please try again.")
			continue
		}

		cwd, err := os.Getwd()
		if err != nil {
			_, file, line, _ := runtime.Caller(0)
			log.Printf("%s:%d: Error getting current directory:", file, line)
			fmt.Println("Error getting current directory:", err)
			continue
		}

		filePath := filepath.Join(cwd, directory, files[fileNumber-1])

		fileContent, err := os.ReadFile(filePath)
		if err != nil {
			_, file, line, _ := runtime.Caller(0)
			log.Printf("%s:%d: Error reading file:", file, line)
			fmt.Println("Error reading file:", err)
			continue
		}

		return string(fileContent)
	}
}

// writeToFile writes data to a file.
func writeToFile(filename, data string) error {

	if data == "" {
		return errors.New("data is empty")
	}

	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString(data)
	if err != nil {
		return err
	}
	return nil
}
