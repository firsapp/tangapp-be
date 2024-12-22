package errorx

import (
	"fmt"

	"github.com/google/uuid"
)

// Custom error for "User not found"
type UserNotFoundError struct {
	ID uuid.UUID
}

func (e *UserNotFoundError) Error() string {
	return fmt.Sprintf("User with ID %s not found", e.ID)
}

// Custom error for Database issues
type DatabaseError struct {
	Err error
}

func (e *DatabaseError) Error() string {
	return fmt.Sprintf("Database error: %v", e.Err)
}
