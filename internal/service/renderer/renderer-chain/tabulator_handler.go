package rendererchain

import (
	"fmt"
	"github.com/DmitryVasilkovW/Git-repository-analyzer.git/internal/model"
	"github.com/DmitryVasilkovW/Git-repository-analyzer.git/internal/service/renderer"
	"os"
)

type TabularHandler struct {
	BaseHandler
}

func (h *TabularHandler) Handle(request model.RepoFlags, result []model.UserData) bool {
	if request.Format == "tabular" {
		err := renderer.WriteTabular(result, request.ShowLanguages)
		if err != nil {
			fmt.Println("Error writing tabular result", err)
			os.Exit(1)
		}
		return true
	}
	return h.BaseHandler.Handle(request, result)
}
