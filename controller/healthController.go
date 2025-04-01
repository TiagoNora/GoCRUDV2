package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"net/http"
)

type HealthController interface {
	Health(c *gin.Context)
}

type healthController struct {
}

func NewHealthController() HealthController {
	return &healthController{}

}

// @Summary Greets from Product Controller
// @Description Responds with a hello message from the Product Controller
// @Tags Status
// @Accept json
// @Produce json
// @Success 200 {object} map[string]string
// @Router /status/health [get]
func (p *healthController) Health(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Hello from Product Controller"})
	log.Info().Msg("Called Hello from Product Controller")
}
