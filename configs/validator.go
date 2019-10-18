package configs

import validator "gopkg.in/go-playground/validator.v9"

type validation struct {
	validator *validator.Validate
}

func (v *validation) Validate(i interface{}) error {
	return v.validator.Struct(i)
}
