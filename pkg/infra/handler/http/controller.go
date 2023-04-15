package http

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/psavelis/go-tuner-api/pkg/domain/ports"
)

type HTTPHandler struct {
	inPort ports.FindNoteByFrequency
}

func New(inPort ports.FindNoteByFrequency) *HTTPHandler {
	return &HTTPHandler{inPort: inPort}
}

func (handler *HTTPHandler) Handle(c *gin.Context) {
	param := c.Param("frequency")

	frequency, err := strconv.ParseFloat(param, 64)

	if err != nil {
		c.JSON(422, gin.H{"422": err.Error()})
		return
	}

	standardNote, err := handler.inPort.Execute(frequency)

	if err != nil {
		c.JSON(404, gin.H{"404": err.Error()})
		return
	}

	c.JSON(200, standardNote)
}
