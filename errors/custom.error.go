package errors

import (
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Error defines a standard application error.
type CustomError struct {
	Id      primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Code    int                `json:"code,omitempty" bson:"code,omitempty"`
	Message string             `json:"message,omitempty" bson:"message,omitempty"`
	Op      string             `json:"op,omitempty" bson:"op,omitempty"`
	Err     string             `json:"err,omitempty" bson:"err,omitempty"`
	Date    string             `json:"string,omitempty" bson:"string,omitempty"`
}

func (e *CustomError) Error() string {
	return fmt.Sprintf("%v", e.Err)
}

// NewError creates a new custom error
func NewError(code int, message, op string, err error) *CustomError {
	return &CustomError{
		Code:    code,
		Message: message,
		Op:      op,
		Err:     err.Error(),
		Date:    time.Now().Format(time.RFC3339),
	}
}
