package i18n

type Dict struct {
	Message string `json:"message"`
}

type Lang map[string]Dict

type I18N struct {
	defaultTag string
	lang map[string]Lang
}

func NewI18N(defaultTag string, lang map[string]Lang) *I18N {
	return &I18N{
		defaultTag: defaultTag,
		lang: lang,
	}
}
