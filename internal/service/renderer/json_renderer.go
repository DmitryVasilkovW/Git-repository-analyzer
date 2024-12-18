package renderer

import (
	"fmt"
	"github.com/DmitryVasilkovW/Git-repository-analyzer.git/internal/model"
	"strings"
)

func WriteJSON(answer []model.UserData, showLanguages bool) error {
	if showLanguages {
		return writeJSONWithLanguages(answer)
	}
	return writeJSONWithoutLanguages(answer)
}

func writeJSONWithoutLanguages(answer []model.UserData) error {
	fmt.Printf("%s", "[")
	for i, k := range answer {
		fmt.Printf("{\"name\":\"%s\",\"lines\":%d,\"commits\":%d,\"files\":%d}", k.Name, k.Lines, len(k.Commits), k.Files)
		if i != len(answer)-1 {
			fmt.Printf(",")
		}
	}
	fmt.Println("]")
	return nil
}

func writeJSONWithLanguages(answer []model.UserData) error {
	fmt.Printf("%s", "[")
	for i, k := range answer {
		langs := formatLanguages(k.Languages)
		fmt.Printf("{\"name\":\"%s\",\"lines\":%d,\"commits\":%d,\"files\":%d,\"languages\":\"%s\"}", k.Name, k.Lines, len(k.Commits), k.Files, langs)
		if i != len(answer)-1 {
			fmt.Printf(",")
		}
	}
	fmt.Println("]")
	return nil
}

func formatLanguages(languages map[string]int) string {
	var langList []string
	for lang, lines := range languages {
		langList = append(langList, fmt.Sprintf("%s (%d)", lang, lines))
	}
	return strings.Join(langList, ", ")
}
