# Introduction
- GoLang international library, easy to use, welcome to share and pull request
- Fast、Easy to Use

# Use
```go
package main

import (
	"fmt"
	"github.com/smithyj/go-i18n/i18n"
	"os"
)

func main()  {
	// Load language config
	path, _ := os.Getwd()
	filename := "/locale/zh.yaml"
	zh, _ := i18n.LoadFileParseLanguage(fmt.Sprintf("%s%s",path, filename))
	filename = "/locale/en.yaml"
	en, _ := i18n.LoadFileParseLanguage(fmt.Sprintf("%s%s",path, filename))

	// Inject to bundle
	// First param is default language option
	// Second param is need translate language map list
	bundle := i18n.NewI18N("zh", map[string]i18n.Lang{
		"zh": zh,
		"en": en,
	})

	// Initialize translate program
	// First param is bundle
	// Second param is visitor lang, maybe source Accept-Language or custom language
	// accord with golang language package standard
	locale, _ := i18n.NewLocale(bundle, "fr")


	message := locale.Translate("hello world")
	fmt.Println(message)

	message = locale.WithTranslateDefault("hello world23", "默认值")
	fmt.Println(message)

	message = locale.WithTranslateParams("hello world", map[string]interface{}{
		"haha": "yangyj",
	})
	fmt.Println(message)

	message = locale.WithTranslateParamsDefault("hello world22223", map[string]interface{}{
		"haha": "yangyj",
	}, "默认值${haha}")
	fmt.Println(message)
}
```

```text
你好，世界，变量替换:${haha}
默认值
你好，世界，变量替换:yangyj
默认值yangyj
```

# Contcat
- Email: 624508914@qq.com
- WECHAT: iyangyj
