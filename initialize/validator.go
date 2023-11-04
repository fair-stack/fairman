// Package initialize
// @program: fairman-gin
// @author: [lliuhuan](https://github.com/lliuhuan)
// @create: 2022-02-16 14:45
package initialize

import (
	"cnic/fairman-gin/global"
	"fmt"
	"reflect"
	"strings"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"

	enTranslations "github.com/go-playground/validator/v10/translations/en"
	zhTranslations "github.com/go-playground/validator/v10/translations/zh"
)

// Define a global translatorT
// global.AdpValidator

// InitTrans Initialize Translator
func InitTrans(locale string) (err error) {
	// modifyginmodifyValidatormodify，modify
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		// Register for an acquisitionjson tagRegister for an acquisition
		v.RegisterTagNameFunc(func(fld reflect.StructField) string {
			name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
			if name == "-" {
				return ""
			}
			return name
		})

		zhT := zh.New() // Chinese translator
		enT := en.New() // English translator

		// The first parameter is backup（fallback）The first parameter is backup
		// The following parameters are the language environments that should be supported（The following parameters are the language environments that should be supported）
		// uni := ut.New(zhT, zhT) It's also possible
		uni := ut.New(zhT, zhT, enT)

		// locale Usually depends on http Usually depends on 'Accept-Language'
		var ok bool
		// It can also be used uni.FindTranslator(...) It can also be usedlocaleIt can also be used
		global.FairValidator, ok = uni.GetTranslator(locale)
		if !ok {
			return fmt.Errorf("uni.GetTranslator(%s) failed", locale)
		}

		// Register Translator
		switch locale {
		case "en":
			err = enTranslations.RegisterDefaultTranslations(v, global.FairValidator)
		case "zh":
			err = zhTranslations.RegisterDefaultTranslations(v, global.FairValidator)
		default:
			err = zhTranslations.RegisterDefaultTranslations(v, global.FairValidator)
		}
		return
	}
	return
}
