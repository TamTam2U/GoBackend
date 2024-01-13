package utils

import (
	"encoding/json"
	"os"
	"reflect"
	"strings"

	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/id"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
)

var (
	Validate *validator.Validate
	Utrans   *ut.UniversalTranslator
	dicts    map[string]string
	forms    map[string]string
	messages map[string]string
)

func loadTranslation(lang string) map[string]map[string]string {
	langFile := "../languages/" + lang + ".json"
	data, err := os.ReadFile(langFile)
	if err != nil {
		return nil
	}

	var translations map[string]map[string]string
	if err := json.Unmarshal(data, &translations); err != nil {
		return nil
	}

	return translations
}

func translationFunc(t ut.Translator, fe validator.FieldError) string {
	field, err := t.T(fe.Field())
	if err != nil {
		field = fe.Field()
	}
	msg, err := t.T(fe.Tag(), field, fe.Param())
	if err != nil {
		return fe.Error()
	}
	return msg
}

func GetTranslationFromHeader(lang string) ut.Translator {
	t, found := Utrans.GetTranslator(lang)
	if !found {
		t, _ = Utrans.GetTranslator("en")
	}
	return t
}

func SetupTranslation(locale string) {
	Validate = validator.New()
	enLocale := en.New()
	Utrans = ut.New(enLocale, enLocale, id.New())

	// Load and use JSON translations for en and id
	language := loadTranslation(locale)

	dicts = language["dict"]
	forms = language["form"]
	messages = language["message"]

	engine, _ := Utrans.GetTranslator(locale) // Changed FindTranslator to GetTranslator
	for key, trans := range dicts {
		_ = engine.Add(key, trans, true) // Use overwrite parameter
	}

	for tag, trans := range forms {
		_ = Validate.RegisterTranslation(tag, engine, func(ut ut.Translator) error {
			return ut.Add(tag, trans, true) // Use overwrite parameter
		}, translationFunc)
	}
}

func GetError(data interface{}, errs validator.ValidationErrorsTranslations) map[string]string {
	ts := map[string]string{}
	rt := reflect.TypeOf(data)
	i := 0

	for index, e := range errs {
		fieldName, err := rt.FieldByName(strings.Split(index, ".")[1])
		if err != true {
			ts[index] = e
		} else {
			ts[string(fieldName.Tag.Get("json"))] = e
		}
		i++
	}
	return ts
}

func GetMessage(key string) string {
	value, ok := messages[key]
	if ok {
		return value
	}
	return key
}
