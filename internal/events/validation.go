package events

import "strings"

func IsValidEventType(eventType string) bool {
	return strings.TrimSpace(eventType) != ""
}
