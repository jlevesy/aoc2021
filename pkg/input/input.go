package input

import (
	"bufio"
	"os"
)

func ReadInput(path string, lineFn func(string) error) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		if err = lineFn(scanner.Text()); err != nil {
			return err
		}
	}

	return scanner.Err()
}
