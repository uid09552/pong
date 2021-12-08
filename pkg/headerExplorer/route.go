package headerExplorer

import (
	"github.com/gin-gonic/gin"
)

func Register(r *gin.Engine) {
	controller := HeaderController{}
	headers := r.Group("/request")
	{
		headers.Any("", controller.GetAllHeaders)
	}
}

type HeaderController struct{}

func (h *HeaderController) GetAllHeaders(c *gin.Context) {
	exp := NewGinExplorer(c)
	c.JSON(200, exp.Get())
}
