package renderer

import (
	"fmt"
	"github.com/DmitryVasilkovW/Git-repository-analyzer.git/internal/model"
	"os"
	"strings"
	"text/tabwriter"
)

func WriteTabular(answer []model.UserData, showLanguages bool) error {
	if showLanguages {
		return writeTabularWithLanguages(answer)
	}
	return writeTabularWithOutLanguages(answer)
}

func writeTabularWithOutLanguages(answer []model.UserData) error {
	w := tabwriter.NewWriter(os.Stdout, 0, 1, 1, ' ', tabwriter.TabIndent)
	_, err := fmt.Fprintln(w, "Name\tLines\tCommits\tFiles")
	if err != nil {
		return err
	}
	for _, k := range answer {
		_, err = fmt.Fprintf(w, "%s\t%d\t%d\t%d", k.Name, k.Lines, len(k.Commits), k.Files)
		if err != nil {
			return err
		}
		_, err = fmt.Fprintln(w)
		if err != nil {
			return err
		}
	}
	err = w.Flush()
	if err != nil {
		return err
	}
	return nil
}

func writeTabularWithLanguages(data []model.UserData) error {
	fmt.Printf("%-20s %-10s %-10s %-10s %s\n", "User", "Files", "Lines", "Commits", "Languages")
	for _, user := range data {
		var langs []string
		for lang, Lines := range user.Languages {
			langs = append(langs, fmt.Sprintf("%s (%d)", lang, Lines))
		}
		fmt.Printf("%-20s %-10d %-10d %-10d %s\n", user.Name, user.Files, user.Lines, len(user.Commits), strings.Join(langs, ", "))
	}
	return nil
}
