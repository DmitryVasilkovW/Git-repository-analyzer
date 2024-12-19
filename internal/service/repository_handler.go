package service

import (
	"fmt"
	"github.com/DmitryVasilkovW/Git-repository-analyzer.git/internal/model"
	"github.com/DmitryVasilkovW/Git-repository-analyzer.git/internal/service/utils"

	//"github.com/go-enry/go-enry/v2"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

func GetRepositoryStatistics(request model.RepoFlags) ([]model.UserData, error) {
	if isRemoteRepository(request.Repository) {
		tmpDir, err := cloneRepository(request.Repository)
		request.Repository = tmpDir
		if err != nil {
			return nil, err
		}
		defer removeDirectory(tmpDir)
	}

	fileTree, err := getFileTree(request)
	if err != nil {
		return nil, err
	}

	usersData, err := getInfoFromAllFiles(request, fileTree)
	if err != nil {
		return nil, err
	}

	return getInfoAboutAllUsers(usersData), nil
}

func isRemoteRepository(url string) bool {
	return strings.HasPrefix(url, "https://") ||
		strings.HasPrefix(url, "git@")
}

func removeDirectory(path string) {
	err := os.RemoveAll(path)
	if err != nil {
		panic(err)
	}
}

func cloneRepository(url string) (string, error) {
	tempDir, err := os.MkdirTemp("", "analyzer-*")
	if err != nil {
		return "", fmt.Errorf("failed to create temporary directory: %w", err)
	}

	cmd := exec.Command("git", "clone", "--quiet", "--no-checkout", url, tempDir)
	if err := cmd.Run(); err != nil {
		return "", fmt.Errorf("failed to clone Repository: %w", err)
	}

	return tempDir, nil
}

func getInfoAboutAllUsers(usersData model.UserDataSet) []model.UserData {
	var userInfo []model.UserData

	for name, stat := range usersData {
		var info model.UserData
		info.Name = name
		info.Files = stat.Files
		info.Lines = stat.Lines
		info.Commits = stat.Commits
		info.Languages = stat.Languages

		userInfo = append(userInfo, info)
	}

	return userInfo
}

func getInfoFromAllFiles(request model.RepoFlags, fileTree []string) (model.UserDataSet, error) {
	userStats := make(model.UserDataSet)
	showLanguages := request.ShowLanguages

	for _, fileName := range fileTree {
		if isInvalidFile(request, fileName) {
			continue
		}

		processedFile, err := handleFile(fileName, request)
		if err != nil {
			return nil, err
		}

		updateUserStats(processedFile, userStats, showLanguages)
	}

	return userStats, nil
}

func updateUserStats(
	processedFile map[string]model.UserData,
	userStats model.UserDataSet,
	showLanguages bool) {
	for name, singleFile := range processedFile {
		userStat, ok := userStats[name]
		if !ok {
			userStat.Commits = make(model.IntSet)
			userStat.Languages = make(map[string]int)
		}
		userStat.Files += 1
		userStat.Lines += singleFile.Lines
		if showLanguages {
			for lang, count := range singleFile.Languages {
				userStat.Languages[lang] += count
			}
		}

		for commitHash := range singleFile.Commits {
			userStat.Commits[commitHash] = 1
		}

		userStats[name] = userStat
	}
}

func isInvalidFile(request model.RepoFlags, fileName string) bool {
	return utils.IsExcludedByPattern(request, fileName) || utils.IsRestricted(request, fileName) ||
		utils.IsExtensionExcluded(request, fileName) || utils.IsLangExcluded(request, fileName)
}

func getFileTree(request model.RepoFlags) ([]string, error) {
	cmd := exec.Command("git", "ls-tree", "-r", "--name-only", request.Revision)
	cmd.Dir = request.Repository
	fileNames, err := cmd.Output()
	if err != nil {
		return nil, err
	}

	if len(fileNames) == 0 {
		return nil, nil
	} else {
		return strings.Split(strings.TrimSpace(string(fileNames)), "\n"), nil
	}
}

func handleFile(fileName string, info model.RepoFlags) (map[string]model.UserData, error) {
	blame, err := getBlame(info, fileName)
	if err != nil {
		return nil, err
	}

	if len(blame) == 0 {
		return emptyBlameHandler(info, fileName)
	}

	return getInfoAboutFile(info, fileName, blame), nil
}

func getInfoAboutFile(info model.RepoFlags, fileName, blame string) model.UserDataSet {
	commitsFromUser, commitAmount := blameHandler(info, blame)
	result := make(model.UserDataSet)

	for commitHash, user := range commitsFromUser {
		stat, ok := result[user]
		if !ok {
			stat.Commits = make(model.IntSet)
			stat.Languages = make(map[string]int)
		}
		stat.Commits[commitHash] = 1
		stat.Lines += commitAmount[commitHash]

		// не удается подключить "github.com/go-enry/go-enry/v2" во время сборки в гитлабе
		//language := enry.GetLanguage(fileName, nil)
		language := ""

		if language != "" {
			stat.Languages[language] += commitAmount[commitHash]
		}

		result[user] = stat
	}

	return result
}

func blameHandler(info model.RepoFlags, blame string) (model.StringSet, model.IntSet) {
	Lines := strings.Split(blame, "\n")
	userRole := getUserRole(info.UseCommitter)
	commitAmount := make(model.IntSet)
	commitsFromUser := make(model.StringSet)

	var currentHash string
	for _, line := range Lines {
		currentFields := strings.Fields(line)

		if len(currentFields) == 0 {
			continue
		}

		if len(currentFields) == 4 {
			currentHash = currentFields[0]
			count, _ := strconv.Atoi(currentFields[3])
			commitAmount[currentHash] += count
			continue
		}
		if currentFields[0] == userRole {
			_, ok := commitsFromUser[currentHash]
			if !ok {
				prefix := userRole + " "
				commitsFromUser[currentHash] = strings.TrimPrefix(line, prefix)
			}
		}
	}

	return commitsFromUser, commitAmount
}

func getUserRole(commiter bool) string {
	if commiter {
		return "committer"
	}
	return "author"
}

func emptyBlameHandler(info model.RepoFlags, fileName string) (model.UserDataSet, error) {
	cmd := exec.Command("git", "log", info.Revision, "-1", "--pretty=format:%H %an", "--", fileName)
	cmd.Dir = info.Repository
	cmdOutput, err := cmd.Output()
	if err != nil {
		return nil, err
	}

	log := string(cmdOutput)
	return handleAllLinesFromLog(log), nil
}

func handleAllLinesFromLog(log string) model.UserDataSet {
	lines := strings.Split(log, "\n")
	result := make(model.UserDataSet)

	for _, line := range lines {
		var singleFile model.UserData

		currentFields := strings.Fields(line)
		prefix := currentFields[0] + " "
		name := strings.TrimPrefix(line, prefix)

		singleFile.Commits = make(model.IntSet)
		singleFile.Commits[currentFields[0]] = 1

		result[name] = singleFile
	}

	return result
}

func getBlame(info model.RepoFlags, fileName string) (string, error) {
	cmd := exec.Command("git", "blame", fileName, "--porcelain", info.Revision)
	cmd.Dir = info.Repository
	cmdOutput, err := cmd.Output()
	if err != nil {
		return "", err
	}

	return string(cmdOutput), nil
}
