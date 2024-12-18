package renderer

import (
	"encoding/csv"
	"fmt"
	"github.com/DmitryVasilkovW/Git-repository-analyzer.git/internal/model"
	"os"
)

func WriteCSV(answer []model.UserData, showLanguages bool) error {
	if showLanguages {
		return writeCSVWithLanguages(answer)
	}
	return writeCSVWithoutLanguages(answer)
}

func writeCSVWithoutLanguages(answer []model.UserData) error {
	csvWriter := csv.NewWriter(os.Stdout)
	defer csvWriter.Flush()

	err := csvWriter.Write([]string{"Name", "Lines", "Commits", "Files"})
	if err != nil {
		return err
	}

	for _, user := range answer {
		err := csvWriter.Write([]string{user.Name, fmt.Sprint(user.Lines), fmt.Sprint(len(user.Commits)), fmt.Sprint(user.Files)})
		if err != nil {
			return err
		}
	}

	csvWriter.Flush()
	return nil
}

func writeCSVWithLanguages(answer []model.UserData) error {
	csvWriter := csv.NewWriter(os.Stdout)
	defer csvWriter.Flush()

	err := csvWriter.Write([]string{"Name", "Lines", "Commits", "Files", "Languages"})
	if err != nil {
		return err
	}

	for _, user := range answer {
		langs := formatLanguages(user.Languages)
		err := csvWriter.Write([]string{user.Name, fmt.Sprint(user.Lines), fmt.Sprint(len(user.Commits)), fmt.Sprint(user.Files), langs})
		if err != nil {
			return err
		}
	}

	csvWriter.Flush()
	return nil
}
