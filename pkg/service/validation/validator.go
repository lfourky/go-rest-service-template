package validation

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
)

// Validator represents custom validator that validates structs based on tags.
type Validator struct {
	validateStruct      *validator.Validate
	universalTranslator *ut.UniversalTranslator
	translator          ut.Translator
	translations        map[string]string
}

// Error represents validation error.
type Error struct {
	Errors []string `json:"validation_errors"`
}

func (e *Error) Error() string {
	return fmt.Sprintf("validation error: (%s)", strings.Join(e.Errors, ", "))
}

// NewValidator returns instance of validator.
func NewValidator() (*Validator, error) {
	en := en.New()
	validator := &Validator{
		validateStruct:      validator.New(),
		universalTranslator: ut.New(en, en),
	}
	validator.setTranslations()

	if err := validator.registerTranslations("en"); err != nil {
		return nil, err
	}

	return validator, nil
}

// Validate validates the struct and returns translated errors.
func (v *Validator) Validate(i interface{}) error {
	if err := v.validateStruct.Struct(i); err != nil {
		if err, ok := err.(validator.ValidationErrors); ok {
			errs := v.translateErrors(err)

			return &Error{Errors: errs}
		}
	}

	return nil
}

func (v *Validator) setTranslations() {
	v.translations = map[string]string{
		"required": "{0} is a required field",
	}
}

// setupTranslations sets up the translations for certain validation rules.
func (v *Validator) registerTranslations(language string) error {
	v.translator, _ = v.universalTranslator.GetTranslator(language)

	for condition, translation := range v.translations {
		condition := condition
		translation := translation

		if err := v.validateStruct.RegisterTranslation(
			condition,
			v.translator,
			func(ut ut.Translator) error {
				return ut.Add(condition, translation, true)
			},
			func(ut ut.Translator, fe validator.FieldError) string {
				pathAndName := strings.SplitN(fe.Namespace(), ".", 2)[1]
				t, _ := ut.T(fe.Tag(), pathAndName, fe.Param())

				return t
			},
		); err != nil {
			return err
		}
	}

	v.validateStruct.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
		if name == "-" {
			return ""
		}

		return name
	})

	return nil
}

func (v *Validator) translateErrors(errs validator.ValidationErrors) []string {
	translatedErrsStr := []string{}
	for _, fieldErr := range errs {
		translatedErrsStr = append(translatedErrsStr, fieldErr.Translate(v.translator))
	}

	return translatedErrsStr
}
