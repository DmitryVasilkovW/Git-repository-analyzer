package sorterchain

import (
	"github.com/DmitryVasilkovW/Git-repository-analyzer.git/internal/model"
	"github.com/DmitryVasilkovW/Git-repository-analyzer.git/internal/service/sorter"
)

type SortByFilesHandler struct {
	BaseHandler
}

func (h *SortByFilesHandler) Handle(request model.RepoFlags, result []model.UserData) bool {
	if request.OrderBy == "files" {
		sorter.SortByFiles(result)
		return true
	}
	return h.BaseHandler.Handle(request, result)
}
