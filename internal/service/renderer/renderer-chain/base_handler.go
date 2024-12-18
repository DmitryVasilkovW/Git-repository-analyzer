package rendererchain

import (
	"fmt"
	"github.com/DmitryVasilkovW/Git-repository-analyzer.git/internal/model"
)

const errorMessage = "Unsupported format: "

type RenderHandler interface {
	Handle(request model.RepoFlags, result []model.UserData) bool
	SetNext(handler RenderHandler) RenderHandler
}

type BaseHandler struct {
	next RenderHandler
}

func NewRenderChain() RenderHandler {
	tabularHandler := &TabularHandler{}
	csvHandler := &CSVHandler{}
	jsonHandler := &JSONHandler{}
	jsonLinesHandler := &JSONLinesHandler{}

	tabularHandler.
		SetNext(csvHandler).
		SetNext(jsonHandler).
		SetNext(jsonLinesHandler)

	return tabularHandler
}

func (h *BaseHandler) SetNext(handler RenderHandler) RenderHandler {
	h.next = handler
	return handler
}

func (h *BaseHandler) Handle(request model.RepoFlags, result []model.UserData) bool {
	if h.next != nil {
		return h.next.Handle(request, result)
	}
	fmt.Println(errorMessage, request.Format)
	return false
}
