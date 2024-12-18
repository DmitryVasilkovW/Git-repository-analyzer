package sorterchain

import "github.com/DmitryVasilkovW/Git-repository-analyzer.git/internal/model"

type SortHandler interface {
	Handle(request model.RepoFlags, result []model.UserData) bool
	SetNext(handler SortHandler) SortHandler
}

type BaseHandler struct {
	next SortHandler
}

func NewSortChain() SortHandler {
	linesHandler := &SortByLinesHandler{}
	commitsHandler := &SortByCommitsHandler{}
	filesHandler := &SortByFilesHandler{}

	linesHandler.
		SetNext(commitsHandler).
		SetNext(filesHandler)

	return linesHandler
}

func (b *BaseHandler) SetNext(handler SortHandler) SortHandler {
	b.next = handler
	return handler
}

func (b *BaseHandler) Handle(request model.RepoFlags, result []model.UserData) bool {
	if b.next != nil {
		return b.next.Handle(request, result)
	}
	return false
}
