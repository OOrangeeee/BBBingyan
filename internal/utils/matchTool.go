package utils

import (
	"regexp"
)

type MatchTool struct {
}

func (m *MatchTool) MatchAtInString(str string) []string {
	var atUsers []string
	re := regexp.MustCompile(`@(\S+?)\b`)
	matches := re.FindAllStringSubmatch(str, -1)
	for _, match := range matches {
		atUsers = append(atUsers, match[1])
	}
	return atUsers
}
