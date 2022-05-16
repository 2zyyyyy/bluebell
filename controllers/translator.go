package controllers

import (
	"fmt"
	"reflect"
	"strings"
	"webapp-scaffold/models"

	zhTranslations "github.com/go-playground/validator/v10/translations/zh"

	"go.uber.org/zap"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	enTranslations "github.com/go-playground/validator/v10/translations/en"
)

// 定义全局翻译器
var trans ut.Translator

// InitTrans 初始化翻译器
func InitTrans(lang string) (err error) {
	// 修改gin框架中的Validator引擎属性，实现自定义
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		// 注册一个获取json tag的自定义方法
		v.RegisterTagNameFunc(func(field reflect.StructField) string {
			name := strings.SplitN(field.Tag.Get("json"), ",", 2)[0]
			if name == "-" {
				return ""
			}
			return name
		})

		// 为ParamSignUp注册自定义校验方法
		v.RegisterStructValidation(ParamSignUpStructLevelValidation, models.ParamSignUp{})

		zhT := zh.New() // 中文翻译器
		enT := en.New() // 英文翻译器

		uni := ut.New(enT, zhT, enT) // 第一个备用的语言环境 后面参数是支持的语言环境

		// lang通常取决于http请求头的 Accept-Language
		var ok bool
		// 也可以使用uni.FindTranslator(...)传入多个lang来查找
		trans, ok = uni.GetTranslator(lang)
		if !ok {
			zap.L().Error("uni.GetTranslator(%s)failed")
			return fmt.Errorf("uni.GetTranslator(%s)failed", lang)
		}
		// 注册翻译器
		switch lang {
		case "en":
			err = enTranslations.RegisterDefaultTranslations(v, trans)
		case "zh":
			err = zhTranslations.RegisterDefaultTranslations(v, trans)
		default:
			err = enTranslations.RegisterDefaultTranslations(v, trans)
		}
		return
	}
	return
}

// 去除提示信息中的结构体名称
func removeTopStruct(fields map[string]string) map[string]string {
	res := map[string]string{}
	for field, err := range fields {
		res[field[strings.Index(field, ".")+1:]] = err
	}
	return res
}

// ParamSignUpStructLevelValidation 自定义ParamSignUp结构体校验函数
func ParamSignUpStructLevelValidation(sl validator.StructLevel) {
	su := sl.Current().Interface().(models.ParamSignUp)

	if su.Password != su.RePassword {
		// 输出错误提示信息，最后一个参数就是传递的param
		sl.ReportError(su.RePassword, "re_password", "RePassword", "eqfield", "password")
	}
}
