# Program "ART Cypher Tool"

The "ART Cypher Tool" program is a tool for encoding and decoding text using the ART algorithm.

## Usage
To run the program, execute the following command in the command line:
```
go run . "encoded data"
```
Where "encrypted data" is the text you want to decrypt. This command runs the decoder, which takes the encrypted data as an argument and outputs the decoding result to the terminal.

### Command Line Options
The program supports the following command line options:
```
go run . -o
```
Opens the main menu for selecting the operation, input method, output method, and user message.
### Error Logging
The program automatically creates an "app.log" log file, which records errors and informational messages. In case of runtime errors, they will be logged in this file for later analysis.

## Main Menu
After starting the program, you will see the main menu, where you need to make a choice:

**1.Encode**: Choose this option if you want to encode a message.
**2.Decode**: Choose this option if you want to decode a message.
**0.Back to top**: Choose this option to return to the main menu, or exit the program.

**Selecting Input and Output Methods**

After choosing an operation (encode or decode), you will need to select the input and output methods:

### Input Method:

**1.Art_terminal**: Choose this method if you want to enter a message manually through the terminal.

**2.Art_IO**: Choose this method if you want to select a file for input.

**0.Back to top**: Choose this option to return to the main menu.

### Output Method:

**1.Terminal**: Choose this method if you want to see the result in the terminal.

**2.File**: Choose this method if you want to write the result to a file.

**0.Back to top**: Choose this option to return to the main menu.

### Entering Message or Selecting File
Depending on the selected input method, you will be prompted to enter a message manually or select a file from the specified directory. Follow the program's prompts to complete this step.

### Encode Operation
Before running the encoding operation (encode_Art), the program checks the entered data. If the data is already encoded, an error is displayed, and the program exits.

### Decode Operation
Before running the decoding operation (decode_Art), the program checks the entered data. If the data is already decoded, an error is displayed, and the program exits.

### Message Status Checking
The program provides functions to check the status of messages:

**- if_decod:** Checks if the message is decoded.

**- if_encod:** Checks if the message is encoded.

**- isBalanced:** Checks if the square brackets in the message are balanced.

## Input from File (Art_IO)
If you choose input method 2. Art_IO, depending on the selected operation (encode or decode), files are taken from different directories:

If you selected operation "1. Encode", files are taken from the decodeArtfile directory.

This directory contains files with ready-made images intended for encoding.
If you selected operation "2. Decode", files are taken from the encodeArtfile directory.

This directory contains files with ready-made encoded data intended for decoding.
Adding Custom Files
When adding a custom file to one of the directories, it will be automatically considered when the program is launched, and a unique number will be assigned to it for selection via the terminal. 

For example:

**1.input.art.txt**

**2.cats_art.txt**

**3.kood.art.txt**

**4.lion.art.txt**

**5.plane.art.txt**
**0.Back to top.**

**Files input.encoded.txt (in the encodeArtfile directory) or input.art.txt (in the decodeArtfile directory) will always be available under number "1". You can also insert your own variants of encoded data for decoding or decoded data for encoding into these files.**

## Obtaining Result

After entering the message or selecting the file, the program will perform encoding or decoding according to your choice. The result will be displayed in the terminal or written to the specified file depending on the selected output method.

## Exiting the Program
After obtaining the result, the program will exit. If you need to perform another operation, simply run the program again and follow the prompts.