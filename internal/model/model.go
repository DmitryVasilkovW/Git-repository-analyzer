package model

type IntSet map[string]int

type StringSet map[string]string

type UserDataSet map[string]UserData

type UserData struct {
	Name      string
	Commits   IntSet
	Files     int
	Lines     int
	Languages map[string]int
}

type RepoFlags struct {
	UseCommitter  bool
	ShowLanguages bool
	Repository    string
	Revision      string
	OrderBy       string
	Format        string
	Extensions    []string
	Languages     []string
	Exclude       []string
	RestrictTo    []string
}
