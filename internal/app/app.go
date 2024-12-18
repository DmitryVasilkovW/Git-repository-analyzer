package app

import (
	"fmt"
	"github.com/DmitryVasilkovW/Git-repository-analyzer.git/internal/model"
	"github.com/DmitryVasilkovW/Git-repository-analyzer.git/internal/service"
	rendererchain "github.com/DmitryVasilkovW/Git-repository-analyzer.git/internal/service/renderer/renderer-chain"
	sorterchain "github.com/DmitryVasilkovW/Git-repository-analyzer.git/internal/service/sorter/sorter-chain"
	flag "github.com/spf13/pflag"
	"os"
)

const (
	errorMessage = "Can't get statistics: "

	repositoryFlagDescription    = "Path to the Git Repository; defaults to the current directory"
	revisionFlagDescription      = "Commit pointer; default HEAD"
	orderByFlagDescription       = "Result sorting key; one of Lines (default), Commits, Files"
	showLanguagesFlagDescription = "Boolean flag showing additional language statistics"
	useCommitterFlagDescription  = "Boolean flag replacing author (default) with committer in calculations"
	formatFlagDescription        = "Output Format; one of tabular (default), csv, json, json-Lines"
	extensionsFlagDescription    = "Extension list, narrowing down the list of Files in the calculation; multiple constraints are separated by commas, e.g. '.java, .md'."
	languagesFlagDescription     = "List of Languages (programming, markup, etc.), narrowing down the list" +
		" of Files in the calculation; multiple constraints are separated by commas, e.g. 'java, markdown'"
	excludeFlagDescription    = "A set of Glob patterns that Exclude Files from the calculation, e.g. 'foo/*,bar/*'"
	restrictToFlagDescription = "Glob pattern set, excluding all Files that do not satisfy any of the patterns in the set"

	repositoryFlagName    = "repository"
	revisionFlagName      = "revision"
	orderByFlagName       = "order-by"
	showLanguagesFlagName = "show-languages"
	useCommitterFlagName  = "use-committer"
	formatFlagName        = "format"
	extensionsFlagName    = "extensions"
	languagesFlagName     = "languages"
	excludeFlagName       = "exclude"
	restrictToFlagName    = "restrict-to"

	repositoryFlagDefaultValue    = "."
	revisionFlagDefaultValue      = "HEAD"
	orderByFlagDefaultValue       = "lines"
	showLanguagesFlagDefaultValue = false
	useCommitterFlagDefaultValue  = false
	formatFlagDefaultValue        = "tabular"
)

var (
	extensionsFlagDefaultValue []string
	languagesFlagDefaultValue  []string
	excludeFlagDefaultValue    []string
	restrictToFlagDefaultValue []string
)

func Run() {
	request := getRequest()

	statistics, err := service.GetRepositoryStatistics(request)
	if err != nil {
		fmt.Println(errorMessage, err)
		os.Exit(1)
	}

	if !sorterchain.NewSortChain().Handle(request, statistics) {
		os.Exit(1)
	}
	if !rendererchain.NewRenderChain().Handle(request, statistics) {
		os.Exit(1)
	}
}

func getRequest() model.RepoFlags {
	var request model.RepoFlags

	flag.StringVar(&request.Repository, repositoryFlagName, repositoryFlagDefaultValue, repositoryFlagDescription)
	flag.StringVar(&request.Revision, revisionFlagName, revisionFlagDefaultValue, revisionFlagDescription)
	flag.StringVar(&request.OrderBy, orderByFlagName, orderByFlagDefaultValue, orderByFlagDescription)
	flag.BoolVar(&request.ShowLanguages, showLanguagesFlagName, showLanguagesFlagDefaultValue, showLanguagesFlagDescription)
	flag.BoolVar(&request.UseCommitter, useCommitterFlagName, useCommitterFlagDefaultValue, useCommitterFlagDescription)
	flag.StringVar(&request.Format, formatFlagName, formatFlagDefaultValue, formatFlagDescription)
	flag.StringSliceVar(&request.Extensions, extensionsFlagName, extensionsFlagDefaultValue, extensionsFlagDescription)
	flag.StringSliceVar(&request.Languages, languagesFlagName, languagesFlagDefaultValue, languagesFlagDescription)
	flag.StringSliceVar(&request.Exclude, excludeFlagName, excludeFlagDefaultValue, excludeFlagDescription)
	flag.StringSliceVar(&request.RestrictTo, restrictToFlagName, restrictToFlagDefaultValue, restrictToFlagDescription)

	flag.Parse()
	return request
}
