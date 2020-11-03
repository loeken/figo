package validation

import (
	"github.com/loeken/figo/pkg/entity"
)

// CreateOrUpdateRelease to verify the required fields
func CreateOrUpdateRelease(release entity.Release) []string {
	var messages []string

	if !isValid(release.Title) {
		messages = append(messages, "Title is required")
	}

	if !isValid(release.Artist) {
		messages = append(messages, "Artist is required")
	}

	return messages
}
