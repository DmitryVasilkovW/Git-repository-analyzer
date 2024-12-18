package rendererchain

import (
	"fmt"
	"github.com/DmitryVasilkovW/Git-repository-analyzer.git/internal/model"
	"github.com/DmitryVasilkovW/Git-repository-analyzer.git/internal/service/renderer"
	"os"
)

type JSONLinesHandler struct {
	BaseHandler
}

func (h *JSONLinesHandler) Handle(request model.RepoFlags, result []model.UserData) bool {
	if request.Format == "json-lines" {
		err := renderer.WriteJSONLines(result, request.ShowLanguages)
		if err != nil {
			fmt.Println("Error writing json-lines result", err)
			os.Exit(1)
		}
		return true
	}
	return h.BaseHandler.Handle(request, result)
}
