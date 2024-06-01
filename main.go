package main

import (
	"fmt"
	"golang.org/x/text/language"
)

func main() {

	//fmt.Println(language.AmericanEnglish.Base())
	//fmt.Println(language.English.Base())
	//base , _:=language.MustParse("en-US").Base()

	tags, _, err := language.ParseAcceptLanguage("en,en;q=0.8,en-US;q=0.9")
	if err != nil {
		panic(err)
	}
	base, _ := tags[0].Base()
	lang := language.MustParse(base.String())
	fmt.Println(lang)
}
