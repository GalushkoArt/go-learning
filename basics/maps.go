package basics

import "fmt"

func Maps() {
	developers := map[string]string{
		"frontender": "JavaScrips",
		"backender":  "Java",
		"mobile":     "Kotlin",
	}
	frontendLanguage := developers["frontender"]
	testLanguage := developers["QA"]
	backendLanguage, ok := developers["backender"]
	fmt.Println(frontendLanguage, testLanguage, ok, backendLanguage)
	developers["QA"] = "Kotlin"
	_, ok = developers["BA"]
	if !ok {
		fmt.Println("BA was not FOUND!!!")
	}

	languageSet := make(map[string]bool)
	for dev, language := range developers {
		fmt.Printf("%s knows %s\n", dev, language)
		languageSet[language] = true
	}
	for language := range languageSet {
		fmt.Print(language, " ")
	}
	fmt.Print("\n")
	fmt.Println(len(developers), len(languageSet))
	delete(developers, "mobile")
	fmt.Printf("%+v", developers)
}
