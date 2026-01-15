package parsinglogfiles

import (
	"fmt"
	"regexp"
)

func IsValidLine(text string) bool {
	var re = regexp.MustCompile(`^\[(TRC|DBG|INF|WRN|ERR|FTL)\]`)
	return re.MatchString(text)
}

func SplitLogLine(text string) []string {
	var re = regexp.MustCompile(`<([\W]*?)>`)
	return re.Split(text, -1)
}

func CountQuotedPasswords(lines []string) int {
	var count int
	var re = regexp.MustCompile(`(?i)(.+?)password"`)
	for _, line := range lines {
		if re.MatchString(line) {
			count++
		}
	}
	return count
}

func RemoveEndOfLineText(text string) string {
	var re = regexp.MustCompile(`end-of-line(\d+)`)
	return re.ReplaceAllString(text, "")
}

func TagWithUserName(lines []string) []string {
	var re = regexp.MustCompile(`User(\s+)([\w]+[\d]+)`)
	for i, line := range lines {
		var matches = re.FindStringSubmatch(line)
		if len(matches) > 0 {
			lines[i] = fmt.Sprintf("[USR] %s %s", matches[len(matches)-1], lines[i])
		}
	}
	return lines
}
