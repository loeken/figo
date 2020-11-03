package validation

import (
	"github.com/loeken/figo/pkg/entity"
)

// CreateOrUpdateContent to verify the required fields
func CreateOrUpdateContent(content entity.Content) []string {
	var messages []string

	if !isValid(content.Title) {
		messages = append(messages, "Title is required")
	}

	if !isValid(content.Body) {
		messages = append(messages, "Body is required")
	}

	return messages
}
