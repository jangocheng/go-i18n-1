package i18n

import (
	"fmt"
	"golang.org/x/text/language"
	"strings"
)

type Locale struct {
	bundle *I18N
	tag []language.Tag
}

func (locale *Locale) Translate(id string) string {
	return locale.WithTranslateDefault(id, "")
}

func (locale *Locale) WithTranslateDefault(id string, defaultMessage string) string {
	params := make(map[string]interface{})
	return locale.WithTranslateParamsDefault(id, params, defaultMessage)
}


func (locale *Locale) WithTranslateParams(id string, params map[string]interface{}) string {
	return locale.WithTranslateParamsDefault(id, params, "")
}

func (locale *Locale) WithTranslateParamsDefault(id string, params map[string]interface{}, defaultMessage string) string {
	maps := locale.bundle.lang

	defaultTag := locale.bundle.defaultTag
	defaultLang := maps[defaultTag]

	result := ""
	var lang Lang
	var dict *Dict
	for _, v := range locale.tag {
		tag := fmt.Sprintf("%s", v)
		v, ok := maps[tag]
		if !ok {
			continue
		}
		lang = v
		d, ok := lang[id]
		if !ok {
			continue
		}
		dict = &d
		break
	}
	// not dict use default
	if dict == nil {
		d := defaultLang[id]
		dict = &d
	}
	if dict != nil {
		result = dict.Message
	}
	// id is not includes
	if result == "" {
		result = id
		if defaultMessage != "" {
			result = defaultMessage
		}
	}

	if result != "" && len(params) > 0 {
		for i, v := range params {
			result = strings.ReplaceAll(result, fmt.Sprintf("${%s}", i), fmt.Sprintf("%s", v))
		}
	}

	return result
}

func NewLocale(bundle *I18N, lang string) (locale *Locale, err error) {
	tag, _, err := language.ParseAcceptLanguage(lang)
	if err != nil {
		return
	}
	locale = &Locale{
		bundle: bundle,
		tag: tag,
	}
	return
}