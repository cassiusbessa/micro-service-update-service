package entity

import (
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Service struct {
	Id          primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Name        string             `json:"name,omitempty" bson:"name,omitempty" validate:"gt=0"`
	MinDuration int                `json:"min_duration,omitempty" bson:"min_duration" validate:"gt=0"`
	MinPrice    float32            `json:"min_price,omitempty" bson:"min_price" validate:"gte=0"`
}

type ValidateError struct {
	Param string `json:"param"`
	Msg   string `json:"msg"`
}

func (c *Service) Validate() []ValidateError {
	valid := validator.New()
	err := valid.Struct(c)
	if err != nil {
		validationErrors := err.(validator.ValidationErrors)
		out := make([]ValidateError, len(validationErrors))
		for i, e := range validationErrors {
			out[i] = ValidateError{e.Field(), message(e)}
		}
		return out
	}
	return nil
}

func message(e validator.FieldError) string {
	switch e.Tag() {
	case "gt":
		return e.Field() + " must be greater than " + e.Param()
	case "gte":
		return e.Field() + " must be greater than or equal to " + e.Param()
	default:
		return ""
	}
}
