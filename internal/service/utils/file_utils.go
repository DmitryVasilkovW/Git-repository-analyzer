package utils

import (
	"github.com/DmitryVasilkovW/Git-repository-analyzer.git/configs"
	"github.com/DmitryVasilkovW/Git-repository-analyzer.git/internal/model"
	"path"
	"path/filepath"
)

func IsLangExcluded(request model.RepoFlags, fileName string) bool {
	if len(request.Languages) == 0 {
		return false
	}

	fileExtension := filepath.Ext(fileName)
	for _, lang := range request.Languages {
		langExtensions := configs.GetExts(lang)

		for _, ext := range langExtensions {
			if ext == fileExtension {
				return false
			}
		}
	}

	return true
}

func IsExtensionExcluded(request model.RepoFlags, fileName string) bool {
	if len(request.Extensions) == 0 {
		return false
	}

	fileExtension := filepath.Ext(fileName)
	for _, ext := range request.Extensions {
		if ext == fileExtension {
			return false
		}
	}

	return true
}

func IsExcludedByPattern(request model.RepoFlags, fileName string) bool {
	if len(request.Exclude) == 0 {
		return false
	}

	for _, pattern := range request.Exclude {
		match, err := path.Match(pattern, fileName)
		if err != nil {
			return true
		}
		if match {
			return true
		}
	}

	return false
}

func IsRestricted(request model.RepoFlags, fileName string) bool {
	if len(request.RestrictTo) == 0 {
		return false
	}

	for _, pattern := range request.RestrictTo {
		match, err := path.Match(pattern, fileName)
		if err != nil {
			return true
		}
		if match {
			return false
		}
	}

	return true
}
