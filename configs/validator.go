package configs

import (
	"errors"

	"github.com/go-playground/locales"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/universal-translator"
	"gopkg.in/go-playground/validator.v9"
	enTranslations "gopkg.in/go-playground/validator.v9/translations/en"
)

var (
	localTranslator     locales.Translator
	universalTranslator *ut.UniversalTranslator
	translator          ut.Translator
)

type validation struct {
	validator *validator.Validate
}

func instantiateValidator() *validator.Validate {
	var found bool
	localTranslator = en.New()
	universalTranslator = ut.New(localTranslator, localTranslator)
	translator, found = universalTranslator.GetTranslator("en")
	if !found {
		framework().Logger.Fatal("translator not found")
	}
	validate := validator.New()
	if err := enTranslations.RegisterDefaultTranslations(validate, translator); err != nil {
		framework().Logger.Fatal(err.Error())
	}
	return customMessages(validate)
}

func (v *validation) Validate(i interface{}) error {
	if err := v.validator.Struct(i); err != nil {
		//TODO change response error to []error
		for _, e := range err.(validator.ValidationErrors) {
			return errors.New(e.Translate(translator))
		}
	}
	return nil
}

var j int

func customMessages(validate *validator.Validate) *validator.Validate {
	messages := map[string]string{
		"required":  "is a required field",
		"email":     "must be a valid email",
		"unique":    "must be array of unique values",
		"numeric":   "must be numeric",
		"min":       "must follow minimum length",
		"max":       "must follow maximum length",
		"uuid":      "must be a valid uuid",
		"latitude":  "must be a valid latitude",
		"longitude": "must be a valid longitude",
	}
	for i, v := range messages {
		if j == len(messages) {
			goto END
		}
		_ = validate.RegisterTranslation(i, translator, func(ut ut.Translator) error {
			return ut.Add(i, "{0} "+v, true)
		}, func(ut ut.Translator, fe validator.FieldError) string {
			t, _ := ut.T(i, fe.Field())
			return t
		})
		j++
		return customMessages(validate)
	}
END:
	return validate
}
