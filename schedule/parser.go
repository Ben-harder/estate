package schedule

import (
	"bufio"
	"os"
	"strconv"
	"strings"
	"time"
)

// Takes a path to an ICS file and returns a slice of jobs
func parseICS(path string) ([]*job, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	jobs := make([]*job, 0)
	scanner := bufio.NewScanner(f)
	var (
		date             time.Time
		responsibilities string
	)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "DTSTART;VALUE=DATE:") {
			date = formatDate(strings.Split(line, ":")[1])
		}
		if strings.Contains(line, "SUMMARY:") {
			unformattedResponsibilities := strings.Split(line, ":")[1]
			if strings.Contains(unformattedResponsibilities, "Garbage") {
				responsibilities = "Garbage/Recycling/Greenbin"
			} else {
				responsibilities = "Recycling/Greenbin"
			}
			jobs = append(jobs, &job{date: date, responsibilities: responsibilities})
		}
	}
	return jobs, nil
}

// Takes date as a string like YYYYMMDD and converts it to the respective time.Time
func formatDate(unformattedDate string) time.Time {
	trimmedDate := strings.TrimSpace(unformattedDate)
	year := trimmedDate[0:4]
	yearInt, _ := strconv.Atoi(year)
	month := trimmedDate[4:6]
	monthInt, _ := strconv.Atoi(month)
	day := trimmedDate[6:8]
	dayInt, _ := strconv.Atoi(day)
	return time.Date(yearInt, time.Month(monthInt), dayInt, 0, 0, 0, 0, time.UTC)
}
