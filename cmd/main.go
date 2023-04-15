package main

import (
	"github.com/gin-gonic/gin"
	"github.com/psavelis/go-tuner-api/pkg/adapter/http"
	"github.com/psavelis/go-tuner-api/pkg/adapter/repository"
	"github.com/psavelis/go-tuner-api/pkg/domain/usecase"
)

func main() {
	router := getRouter()

	router.Run(":8080")
}

func getRouter() *gin.Engine {
	noteRepository := repository.NewInMemoryNotesRepository()
	findNoteUseCase := usecase.New(noteRepository)
	httpHandler := http.NewHTTPHandler(findNoteUseCase)

	router := gin.New()
	router.GET("/tune/:frequency", httpHandler.Handle)
	return router
}
