package http

import (
	"github.com/0xirvan/tdl-svelte-go/server/internal/core/domain"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/id"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
	id_translations "github.com/go-playground/validator/v10/translations/id"
)

type RequestValidator struct {
	validate *validator.Validate
	trans    ut.Translator
}

func NewRequestValidator(language string) *RequestValidator {
	v := validator.New()

	var uni *ut.UniversalTranslator
	switch language {
	case "id":
		uni = ut.New(id.New(), id.New(), en.New())
	default:
		uni = ut.New(en.New(), en.New(), id.New())
	}

	trans, _ := uni.GetTranslator(language)

	switch language {
	case "id":
		_ = id_translations.RegisterDefaultTranslations(v, trans)
	default:
		_ = en_translations.RegisterDefaultTranslations(v, trans)
	}

	return &RequestValidator{validate: v, trans: trans}
}

func (rv *RequestValidator) Validate(i any) error {
	if err := rv.validate.Struct(i); err != nil {
		if ve, ok := err.(validator.ValidationErrors); ok {
			translated := ve.Translate(rv.trans)
			return domain.ValidationErrors(translated)
		}
		return err
	}
	return nil
}
