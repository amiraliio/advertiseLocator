package configs

import (
	"errors"

	"github.com/go-playground/locales"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/universal-translator"
	validator "gopkg.in/go-playground/validator.v9"
	en_translations "gopkg.in/go-playground/validator.v9/translations/en"
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
	if err := en_translations.RegisterDefaultTranslations(validate, translator); err != nil {
		framework().Logger.Fatal(err.Error())
	}
	//required
	customeRequiredMessage(validate)
	return validate
}

func (v *validation) Validate(i interface{}) error {
	err := v.validator.Struct(i)
	//TODO change response error to []error
	for _, e := range err.(validator.ValidationErrors) {
		return errors.New(e.Translate(translator))
	}
	return nil
}

func customeRequiredMessage(validate *validator.Validate) {
	_ = validate.RegisterTranslation("required", translator, func(ut ut.Translator) error {
		return ut.Add("required", "{0} is a required field", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("required", fe.Field())
		return t
	})
}
