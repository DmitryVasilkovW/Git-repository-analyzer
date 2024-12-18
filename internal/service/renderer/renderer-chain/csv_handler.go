package rendererchain

import (
	"fmt"
	"github.com/DmitryVasilkovW/Git-repository-analyzer.git/internal/model"
	"github.com/DmitryVasilkovW/Git-repository-analyzer.git/internal/service/renderer"
	"os"
)

type CSVHandler struct {
	BaseHandler
}

func (h *CSVHandler) Handle(request model.RepoFlags, result []model.UserData) bool {
	if request.Format == "csv" {
		err := renderer.WriteCSV(result, request.ShowLanguages)
		if err != nil {
			fmt.Println("Error writing csv result", err)
			os.Exit(1)
		}
		return true
	}
	return h.BaseHandler.Handle(request, result)
}
