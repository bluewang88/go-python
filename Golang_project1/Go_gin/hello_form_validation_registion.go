package Go_gin

import (
	"fmt"
	"net/http"
	"reflect"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator" // Import universal translator as ut
	"github.com/go-playground/validator/v10"           //引入validator库
	en_translation "github.com/go-playground/validator/v10/translations/en"
	zh_translation "github.com/go-playground/validator/v10/translations/zh"
)

// 测试
// curl --location 'localhost:8080/signup' \
// --header 'Content-Type: application/json' \
//
//	--data-raw '{
//	    "username": "wangwei",
//	    "age": 12,
//	    "password": "123456",
//	    "re_password": "123456",
//	    "email": "123@123.com"
//	}'

// 错误测试
// curl --location 'localhost:8080/signup' \
// --header 'Content-Type: application/json' \
//
//	--data-raw '{
//	    "username": "wangwei",
//	    "age": 12,
//	    "password": "123456",
//	    "re_password": "12345",
//	    "email": "123@123.com"
//	}'
//
//	{
//	    "error": {
//	        "RegisterForm.RePassword": "RePassword必须等于Password"
//	    }
//	}
type RegisterForm struct {
	Username   string `json:"username" binding:"required,min=5,max=10"`
	Age        int    `json:"age" binding:"required,gte=1,lte=130"` // gte: greater than or equal, lte: less than or equal
	Password   string `json:"password" binding:"required"`
	RePassword string `json:"re_password" binding:"required,eqfield=Password"` // eqfield: equal to the specified field
	Email      string `json:"email" binding:"required,email"`                  // email: email format
}

// remoteTopStruct函数的作用是将验证器返回的错误信息中的字段名中的结构体前缀去掉，只保留字段名
// remoteTopStruct 函数用于处理传入的字段映射，移除字段名中的顶层结构前缀。
// 例如，如果字段名为 "top.field"，则移除 "top." 前缀，保留 "field"。
//
// 参数:
//   - fields: 包含字段名和对应值的映射，字段名可能包含顶层结构前缀。
//
// 返回值:
//   - map[string]string: 处理后的字段映射，移除了字段名中的顶层结构前缀。
func remoteTopStruct(fields map[string]string) map[string]string {
	rsp := map[string]string{} // 创建一个新的字段映射

	// 遍历输入字段映射，处理每个字段名
	for key, value := range fields {
		// 查找字段名中的 "." 分隔符，判断是否存在顶层结构前缀
		index := strings.Index(key, ".")
		if index != -1 {
			// 如果存在顶层结构前缀，则移除前缀并保留剩余部分作为新的字段名
			rsp[key[index+1:]] = value
		} else {
			// 如果不存在顶层结构前缀，则直接使用原字段名
			rsp[key] = value
		}
	}

	return rsp
}

func InitTrans(locale string) (trans ut.Translator, err error) {

	//修改因框架中的validator的引擎属性，实现定制
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		//注册一个获取json的tag的自定义方式
		v.RegisterTagNameFunc(func(fld reflect.StructField) string { //自定义获取json tag的方法
			name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
			if name == "-" {
				return ""
			}
			return name
		})

		//中文翻译器
		zhTrans := zh.New()
		enTrans := en.New() //英文翻译器
		//第一个参数是备用（英文）语言环境，后面的参数是应该支持的语言环境
		uni := ut.New(zhTrans, enTrans, zhTrans)

		var ok bool
		trans, ok = uni.GetTranslator(locale)

		if !ok {
			return nil, fmt.Errorf("uni.GetTranslator(%s) failed", locale)
		}

		switch locale {
		case "zh":
			zh_translation.RegisterDefaultTranslations(v, trans)
		case "en":
			en_translation.RegisterDefaultTranslations(v, trans)

		default:
			return nil, fmt.Errorf("unsupported locale: %s", locale)
		}
		return trans, nil
	}
	return nil, fmt.Errorf("binding.Validator.Engine() is not a validator.Validate")
}

func HelloFormValidationRegistion() {
	trans, err := InitTrans("zh")
	if err != nil {
		fmt.Println("InitTrans failed, err:", err)
		// panic(err)
		return
	}
	r := gin.Default()
	r.POST("/signup", func(c *gin.Context) {
		var registerform RegisterForm
		if err := c.ShouldBindJSON(&registerform); err != nil {
			errs, ok := err.(validator.ValidationErrors)
			if !ok {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}
			// 使用修复后的remoteTopStruct函数
			c.JSON(http.StatusBadRequest, gin.H{"error": remoteTopStruct(errs.Translate(trans))})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "Registration successful!", "data": registerform})
	})
	r.Run()
}
