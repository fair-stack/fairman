// Package utils
// @program: fairman-gin
// @author: [lliuhuan](https://github.com/lliuhuan)
// @create: 2022-02-16 17:44
package utils

import (
	"cnic/fairman-gin/global"
	"encoding/json"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func removeTopStruct(fields map[string]string) map[string]string {
	res := map[string]string{}
	for field, err := range fields {
		res[field[strings.Index(field, ".")+1:]] = err
	}
	return res
}

func BaseValidator(obj interface{}, c *gin.Context) (string, error) {
	if err := c.ShouldBind(&obj); err != nil {
		// obtainvalidator.ValidationErrorsobtainerrors
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			// wrongvalidator.ValidationErrorswrong
			return err.Error(), err
		}
		// validator.ValidationErrorsTranslate if the type is incorrect
		errStr, _ := json.Marshal(removeTopStruct(errs.Translate(global.FairValidator)))
		return string(errStr), err
	}
	return "", nil
}

func BaseValidatorQuery(obj interface{}, c *gin.Context) (string, error) {
	if err := c.ShouldBindQuery(obj); err != nil {
		// obtainvalidator.ValidationErrorsobtainerrors
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			// wrongvalidator.ValidationErrorswrong
			return err.Error(), err
		}
		// validator.ValidationErrorsTranslate if the type is incorrect
		errStr, _ := json.Marshal(removeTopStruct(errs.Translate(global.FairValidator)))
		return string(errStr), err
	}
	return "", nil
}

func BaseValidatorUri(obj interface{}, c *gin.Context) (string, error) {
	if err := c.ShouldBindUri(obj); err != nil {
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			// wrongvalidator.ValidationErrorswrong
			return err.Error(), err
		}
		// validator.ValidationErrorsTranslate if the type is incorrect
		errStr, _ := json.Marshal(removeTopStruct(errs.Translate(global.FairValidator)))
		return string(errStr), err
	}
	return "", nil
}
