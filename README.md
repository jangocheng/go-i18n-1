# Introduction
- GoLang international library, easy to use, welcome to share and pull request
- Fast、Easy to Use

# Use
```go
// Load language config
path, _ := os.Getwd()
filename := "/example/locale/zh.yaml"
zh, _ := LoadFileParseLanguage(fmt.Sprintf("%s%s",path, filename))
filename = "/example/locale/en.yaml"
en, _ := LoadFileParseLanguage(fmt.Sprintf("%s%s",path, filename))

// Inject to bundle
// First param is default language option
// Second param is need translate language map list
bundle := NewI18N("zh", map[string]Lang{
"zh": zh,
"en": en,
})

// Initialize translate program
// First param is bundle
// Second param is visitor lang, maybe source Accept-Language or custom language
// accord with golang language package standard
locale, _ := NewLocale(bundle, "fr")


message := locale.Translate("hello world")
testing.Log(message)

message = locale.WithTranslateDefault("hello world23", "默认值")
testing.Log(message)

message = locale.WithTranslateParams("hello world", map[string]interface{}{
"haha": "yangyj",
})
testing.Log(message)

message = locale.WithTranslateParamsDefault("hello world22223", map[string]interface{}{
"haha": "yangyj",
}, "默认值${haha}")
testing.Log(message)
```

```shell
=== RUN   TestI18n
    i18n_test.go:33: 你好，世界，变量替换:${haha}
    i18n_test.go:36: 默认值
    i18n_test.go:41: 你好，世界，变量替换:yangyj
    i18n_test.go:46: 默认值yangyj
--- PASS: TestI18n (0.00s)
PASS

Process finished with exit code 0
```

# Contcat
- Email: 624508914@qq.com
- WECHAT: iyangyj
