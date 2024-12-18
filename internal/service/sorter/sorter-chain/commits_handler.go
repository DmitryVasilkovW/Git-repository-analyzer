package sorterchain

import (
	"github.com/DmitryVasilkovW/Git-repository-analyzer.git/internal/model"
	"github.com/DmitryVasilkovW/Git-repository-analyzer.git/internal/service/sorter"
)

type SortByCommitsHandler struct {
	BaseHandler
}

func (h *SortByCommitsHandler) Handle(request model.RepoFlags, result []model.UserData) bool {
	if request.OrderBy == "commits" {
		sorter.SortByCommits(result)
		return true
	}
	return h.BaseHandler.Handle(request, result)
}
