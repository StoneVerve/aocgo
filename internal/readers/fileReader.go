/*
 * Package readers provides  functionality for reading input text files during the advento of code 2023
 */
package readers

import (
	"bufio"
	"os"
)

/*
 * Returns the contentes of a text file as an slice of string. The length of slice corresponds the amount of
 * rows in the text file.
 * Each row in the text file ends with '\n'
 * If there is an error while tryinf to read the text file, the function will return the value nil and the corresponding
 * I/O error
 */
func ReadFile(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	var text []string
	reader := bufio.NewScanner(file)
	for reader.Scan() {
		text = append(text, reader.Text())
	}
	return text, nil
}
