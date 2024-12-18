package rendererchain

import (
	"fmt"
	"github.com/DmitryVasilkovW/Git-repository-analyzer.git/internal/model"
	"github.com/DmitryVasilkovW/Git-repository-analyzer.git/internal/service/renderer"
	"os"
)

type JSONHandler struct {
	BaseHandler
}

func (h *JSONHandler) Handle(request model.RepoFlags, result []model.UserData) bool {
	if request.Format == "json" {
		err := renderer.WriteJSON(result, request.ShowLanguages)
		if err != nil {
			fmt.Println("Error writing JSON result", err)
			os.Exit(1)
		}
		return true
	}
	return h.BaseHandler.Handle(request, result)
}
