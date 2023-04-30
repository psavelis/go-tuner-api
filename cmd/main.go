package main

import (
	"github.com/gin-gonic/gin"
	"github.com/psavelis/go-tuner-api/pkg/domain/usecase"
	"github.com/psavelis/go-tuner-api/pkg/infra/handler/http"
	"github.com/psavelis/go-tuner-api/pkg/infra/repository/memory"
)

func main() {
	router := getRouter()

	router.Run(":3000")
}

func getRouter() *gin.Engine {
	noteRepository := memory.NewInMemoryNotesRepository()
	findNoteUseCase := usecase.NewFindNoteUseCase(noteRepository)
	httpHandler := http.NewHTTPHandler(findNoteUseCase)

	router := gin.New()
	router.GET("/tune/:frequency", httpHandler.Handle)
	return router
}
