package sorter

import (
	"github.com/DmitryVasilkovW/Git-repository-analyzer.git/internal/model"
	"sort"
	"strings"
)

func SortByLine(answer []model.UserData) []model.UserData {
	sort.SliceStable(answer, func(i, j int) bool {
		return compareUserData(answer[i], answer[j], "Lines", "Commits", "Files")
	})
	return answer
}

func SortByCommits(answer []model.UserData) []model.UserData {
	sort.SliceStable(answer, func(i, j int) bool {
		return compareUserData(answer[i], answer[j], "Commits", "Lines", "Files")
	})
	return answer
}

func SortByFiles(answer []model.UserData) []model.UserData {
	sort.SliceStable(answer, func(i, j int) bool {
		return compareUserData(answer[i], answer[j], "Files", "Lines", "Commits")
	})
	return answer
}

func compareUserData(a, b model.UserData, primary, secondary, tertiary string) bool {
	switch primary {
	case "Lines":
		if a.Lines != b.Lines {
			return a.Lines > b.Lines
		}
	case "Commits":
		if len(a.Commits) != len(b.Commits) {
			return len(a.Commits) > len(b.Commits)
		}
	case "Files":
		if a.Files != b.Files {
			return a.Files > b.Files
		}
	}

	switch secondary {
	case "Lines":
		if a.Lines != b.Lines {
			return a.Lines > b.Lines
		}
	case "Commits":
		if len(a.Commits) != len(b.Commits) {
			return len(a.Commits) > len(b.Commits)
		}
	case "Files":
		if a.Files != b.Files {
			return a.Files > b.Files
		}
	}

	switch tertiary {
	case "Lines":
		if a.Lines != b.Lines {
			return a.Lines > b.Lines
		}
	case "Commits":
		if len(a.Commits) != len(b.Commits) {
			return len(a.Commits) > len(b.Commits)
		}
	case "Files":
		if a.Files != b.Files {
			return a.Files > b.Files
		}
	}

	return strings.Compare(a.Name, b.Name) < 0
}
