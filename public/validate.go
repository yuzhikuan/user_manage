package public

import (
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	validator "gopkg.in/go-playground/validator.v9"
)

var (
	Uni      *ut.UniversalTranslator
	Validate *validator.Validate
)

func InitValidate() {
	en := en.New()
	zh := zh.New()
	Uni = ut.New(en, zh)
	Validate = validator.New()
}
