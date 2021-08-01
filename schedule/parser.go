package schedule

import (
	"bufio"
	"os"
	"strings"
)

func Parse(path string) error {
	f, err := os.Open(path)
	if err != nil {
		return err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	// optionally, resize scanner's capacity for lines over 64K, see next example
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "DTSTART;VALUE=DATE:") {
			date := strings.Split(line, ":")[1]
		}
		if strings.Contains(line, "SUMMARY:") {
			unformattedResponsibilities := strings.Split(line, ":")[1]
		}
	}
	return nil
}
