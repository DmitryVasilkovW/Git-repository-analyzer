//go:build !solution

package renderer

import (
	"fmt"
	"github.com/DmitryVasilkovW/Git-repository-analyzer.git/internal/model"
)

func WriteJSONLines(answer []model.UserData, showLanguages bool) error {
	if showLanguages {
		return writeJSONLinesWithLanguages(answer)
	}
	return writeJSONLinesWithoutLanguages(answer)
}

func writeJSONLinesWithoutLanguages(answer []model.UserData) error {
	for _, k := range answer {
		fmt.Printf("{\"name\":\"%s\",\"lines\":%d,\"commits\":%d,\"files\":%d}\n", k.Name, k.Lines, len(k.Commits), k.Files)
	}
	return nil
}

func writeJSONLinesWithLanguages(answer []model.UserData) error {
	for _, k := range answer {
		langs := formatLanguages(k.Languages)
		fmt.Printf("{\"name\":\"%s\",\"lines\":%d,\"commits\":%d,\"files\":%d,\"languages\":\"%s\"}\n", k.Name, k.Lines, len(k.Commits), k.Files, langs)
	}
	return nil
}
