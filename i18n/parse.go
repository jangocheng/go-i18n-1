package i18n

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

func LoadFileParseLanguage(filename string) (lang Lang, err error) {
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		return
	}
	err = yaml.Unmarshal(bytes, &lang)
	if err != nil {
		return
	}
	return
}
